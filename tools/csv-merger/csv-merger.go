package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/coreos/go-semver/semver"
	yaml "github.com/ghodss/yaml"
	csvv1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type csvClusterPermissions struct {
	ServiceAccountName string              `json:"serviceAccountName"`
	Rules              []rbacv1.PolicyRule `json:"rules"`
}

type csvPermissions struct {
	ServiceAccountName string              `json:"serviceAccountName"`
	Rules              []rbacv1.PolicyRule `json:"rules"`
}

type csvDeployments struct {
	Name string                `json:"name"`
	Spec appsv1.DeploymentSpec `json:"spec,omitempty"`
}

type csvStrategySpec struct {
	ClusterPermissions []csvClusterPermissions `json:"clusterPermissions"`
	Permissions        []csvPermissions        `json:"permissions"`
	Deployments        []csvDeployments        `json:"deployments"`
}

var (
	csvVersion           = flag.String("csv-version", "", "the unified CSV version")
	replacesCsvVersion   = flag.String("replaces-csv-version", "", "the unified CSV version this new CSV will replace")
	rookCSVStr           = flag.String("rook-csv-filepath", "", "path to rook csv yaml file")
	noobaaCSVStr         = flag.String("noobaa-csv-filepath", "", "path to noobaa csv yaml file")
	ocsCSVStr            = flag.String("ocs-csv-filepath", "", "path to ocs csv yaml file")
	rookContainerImage   = flag.String("rook-container-image", "", "rook operator container image")
	noobaaContainerImage = flag.String("noobaa-container-image", "", "noobaa operator container image")
	ocsContainerImage    = flag.String("ocs-container-image", "", "ocs operator container image")

	inputCrdsDir = flag.String("crds-directory", "", "The directory containing all the crds to be included in the registry bundle")

	outputDir = flag.String("olm-bundle-directory", "", "The directory to output the unified CSV and CRDs to")
)

type templateData struct {
	RookOperatorImage        string
	RookOperatorCsvVersion   string
	NoobaaOperatorImage      string
	NoobaaOperatorCsvVersion string
	OcsOperatorCsvVersion    string
	OcsOperatorImage         string
}

func finalizedCsvFilename() string {
	return "ocs-operator.v" + *csvVersion + ".clusterserviceversion.yaml"
}

func copyFile(src string, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	outFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, srcFile)
	if err != nil {
		panic(err)
	}
}

func unmarshalCSV(filePath string) *csvv1.ClusterServiceVersion {
	data := templateData{
		RookOperatorImage:        *rookContainerImage,
		NoobaaOperatorImage:      *noobaaContainerImage,
		NoobaaOperatorCsvVersion: *csvVersion,
		RookOperatorCsvVersion:   *csvVersion,
		OcsOperatorCsvVersion:    *csvVersion,
		OcsOperatorImage:         *ocsContainerImage,
	}

	writer := strings.Builder{}

	fmt.Printf("reading in csv at %s\n", filePath)
	tmpl := template.Must(template.ParseFiles(filePath))
	err := tmpl.Execute(&writer, data)
	if err != nil {
		panic(err)
	}

	bytes := []byte(writer.String())

	csvStruct := &csvv1.ClusterServiceVersion{}
	err = yaml.Unmarshal(bytes, csvStruct)
	if err != nil {
		panic(err)
	}

	return csvStruct
}

func unmarshalStrategySpec(csv *csvv1.ClusterServiceVersion) *csvStrategySpec {

	templateStrategySpec := &csvStrategySpec{}
	err := json.Unmarshal(csv.Spec.InstallStrategy.StrategySpecRaw, templateStrategySpec)
	if err != nil {
		panic(err)
	}

	if strings.Contains(csv.Name, "noobaa") {
		// TODO remove this once issue https://github.com/noobaa/noobaa-operator/issues/35
		// is resolved.
		// Until then, noobaa's CSV isn't actually a template, which means we have
		// to explicitly inject the noobaa container image into the deployment spec.
		templateStrategySpec.Deployments[0].Spec.Template.Spec.Containers[0].Image = *noobaaContainerImage
	}

	return templateStrategySpec
}

func marshallObject(obj interface{}, writer io.Writer) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	var r unstructured.Unstructured
	if err := json.Unmarshal(jsonBytes, &r.Object); err != nil {
		return err
	}

	// remove status and metadata.creationTimestamp
	unstructured.RemoveNestedField(r.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "template", "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "spec", "template", "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "status")

	deployments, exists, err := unstructured.NestedSlice(r.Object, "spec", "install", "spec", "deployments")
	if exists {
		for _, obj := range deployments {
			deployment := obj.(map[string]interface{})
			unstructured.RemoveNestedField(deployment, "metadata", "creationTimestamp")
			unstructured.RemoveNestedField(deployment, "spec", "template", "metadata", "creationTimestamp")
			unstructured.RemoveNestedField(deployment, "status")
		}
		unstructured.SetNestedSlice(r.Object, deployments, "spec", "install", "spec", "deployments")
	}

	jsonBytes, err = json.Marshal(r.Object)
	if err != nil {
		return err
	}

	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		return err
	}

	// fix double quoted strings by removing unneeded single quotes...
	s := string(yamlBytes)
	s = strings.Replace(s, " '\"", " \"", -1)
	s = strings.Replace(s, "\"'\n", "\"\n", -1)

	yamlBytes = []byte(s)

	_, err = writer.Write([]byte("---\n"))
	if err != nil {
		return err
	}

	_, err = writer.Write(yamlBytes)
	if err != nil {
		return err
	}

	return nil
}

func generateUnifiedCSV() *csvv1.ClusterServiceVersion {

	csvs := []string{
		*rookCSVStr,
		*noobaaCSVStr,
		*ocsCSVStr,
	}

	ocsCSV := unmarshalCSV(*ocsCSVStr)
	templateStrategySpec := unmarshalStrategySpec(ocsCSV)

	for _, csvStr := range csvs {
		if csvStr != "" {
			csvStruct := unmarshalCSV(csvStr)
			strategySpec := unmarshalStrategySpec(csvStruct)

			deployments := strategySpec.Deployments
			clusterPermissions := strategySpec.ClusterPermissions
			permissions := strategySpec.Permissions

			templateStrategySpec.Deployments = append(templateStrategySpec.Deployments, deployments...)
			templateStrategySpec.ClusterPermissions = append(templateStrategySpec.ClusterPermissions, clusterPermissions...)
			templateStrategySpec.Permissions = append(templateStrategySpec.Permissions, permissions...)

			for _, owned := range csvStruct.Spec.CustomResourceDefinitions.Owned {
				ocsCSV.Spec.CustomResourceDefinitions.Required = append(
					ocsCSV.Spec.CustomResourceDefinitions.Required,
					csvv1.CRDDescription{
						Name:    owned.Name,
						Version: owned.Version,
						Kind:    owned.Kind,
					})
			}
		}
	}

	// Re-serialize deployments and permissions into csv strategy.
	updatedStrat, err := json.Marshal(templateStrategySpec)
	if err != nil {
		panic(err)
	}
	ocsCSV.Spec.InstallStrategy.StrategySpecRaw = updatedStrat

	// Set correct csv versions and name
	semverVersion := semver.New(*csvVersion)
	ocsCSV.Spec.Version = *semverVersion
	ocsCSV.Name = "ocs-operator.v" + *csvVersion
	if *replacesCsvVersion != "" {
		ocsCSV.Spec.Replaces = "ocs-operator.v" + *replacesCsvVersion
	}

	// Set api maturity
	ocsCSV.Spec.Maturity = "alpha"

	// Set maintainers
	ocsCSV.Spec.Maintainers = []csvv1.Maintainer{
		{
			Name:  "Jose Rivera",
			Email: "jarrpa@redhat.com",
		},
		{
			Name:  "Kaushal M",
			Email: "kaushal@redhat.com",
		},
	}

	// Set links
	ocsCSV.Spec.Links = []csvv1.AppLink{
		{
			Name: "Source Code",
			URL:  "https://github.com/openshift/ocs-operator",
		},
	}

	// Set Keywords
	ocsCSV.Spec.Keywords = []string{
		"storage",
		"rook",
		"ceph",
		"noobaa",
		"block storage",
		"shared filesystem",
		"object storage",
	}

	// Set Provider
	ocsCSV.Spec.Provider = csvv1.AppLink{
		Name: "Red Hat",
	}

	return ocsCSV
}

func main() {
	flag.Parse()

	if *csvVersion == "" {
		log.Fatal("--csv-version is required")
	} else if *rookCSVStr == "" {
		log.Fatal("--rook-csv-filepath is required")
	} else if *noobaaCSVStr == "" {
		log.Fatal("--noobaa-csv-filepath is required")
	} else if *ocsCSVStr == "" {
		log.Fatal("--ocs-csv-filepath is required")
	} else if *rookContainerImage == "" {
		log.Fatal("--rook-container-image is required")
	} else if *noobaaContainerImage == "" {
		log.Fatal("--noobaa-container-image is required")
	} else if *ocsContainerImage == "" {
		log.Fatal("--ocs-container-image is required")
	} else if *inputCrdsDir == "" {
		log.Fatal("--crds-directory is required")
	} else if *outputDir == "" {
		log.Fatal("--olm-bundle-directory is required")
	}

	// start with a fresh output directory if it already exists
	os.RemoveAll(*outputDir)

	// create output directory
	os.MkdirAll(*outputDir, os.FileMode(0755))
	os.MkdirAll(filepath.Join(*outputDir, "crds"), os.FileMode(0755))

	ocsCSV := generateUnifiedCSV()

	// write unified CSV to out dir
	writer := strings.Builder{}
	marshallObject(ocsCSV, &writer)
	err := ioutil.WriteFile(filepath.Join(*outputDir, finalizedCsvFilename()), []byte(writer.String()), 0644)

	fmt.Printf("CSV written to %s\n", filepath.Join(*outputDir, finalizedCsvFilename()))

	crds, err := ioutil.ReadDir(*inputCrdsDir)
	if err != nil {
		panic(err)
	}

	for _, crd := range crds {
		// only copy crd manifests, this will ignore cr manifests
		if strings.Contains(crd.Name(), "crd.yaml") {
			fmt.Printf("CRD %s written to %s\n", crd.Name(), filepath.Join(*outputDir, crd.Name()))
			copyFile(filepath.Join(*inputCrdsDir, crd.Name()), filepath.Join(*outputDir, crd.Name()))
		}
	}
}

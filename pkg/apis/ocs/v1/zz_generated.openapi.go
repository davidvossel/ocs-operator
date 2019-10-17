// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitialization":                  schema_pkg_apis_ocs_v1_OCSInitialization(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationSpec":              schema_pkg_apis_ocs_v1_OCSInitializationSpec(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationStatus":            schema_pkg_apis_ocs_v1_OCSInitializationStatus(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageCluster":                     schema_pkg_apis_ocs_v1_StorageCluster(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitialization":       schema_pkg_apis_ocs_v1_StorageClusterInitialization(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationSpec":   schema_pkg_apis_ocs_v1_StorageClusterInitializationSpec(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationStatus": schema_pkg_apis_ocs_v1_StorageClusterInitializationStatus(ref),
		"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterStatus":               schema_pkg_apis_ocs_v1_StorageClusterStatus(ref),
	}
}

func schema_pkg_apis_ocs_v1_OCSInitialization(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OCSInitialization is the Schema for the ocsinitialization API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationSpec", "github.com/openshift/ocs-operator/pkg/apis/ocs/v1.OCSInitializationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ocs_v1_OCSInitializationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OCSInitializationSpec defines the desired state of OCSInitialization",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_ocs_v1_OCSInitializationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OCSInitializationStatus defines the observed state of OCSInitialization",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase describes the Phase of OCSInitialization This is used by OLM UI to provide status information to the user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions describes the state of the OCSInitialization resource.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects created and maintained by this operator. Object references will be added to this list after they have been created AND found in the cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"errorMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"sCCsCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_ocs_v1_StorageCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageCluster is the Schema for the storageclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterSpec", "github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ocs_v1_StorageClusterInitialization(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageClusterInitialization is the Schema for the storageclusterinitializations API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationSpec", "github.com/openshift/ocs-operator/pkg/apis/ocs/v1.StorageClusterInitializationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ocs_v1_StorageClusterInitializationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageClusterInitializationSpec defines the desired state of StorageClusterInitialization",
				Properties: map[string]spec.Schema{
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources is set by the StorageCluster controller when it creates the StorageClusterInitialization resource, and is set to the StorageCluster.Spec.Resources",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
									},
								},
							},
						},
					},
					"failureDomain": {
						SchemaProps: spec.SchemaProps{
							Description: "FailureDomain is the base CRUSH element Ceph will use to distribute its data replicas for the default CephBlockPool",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_ocs_v1_StorageClusterInitializationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageClusterInitializationStatus defines the observed state of StorageClusterInitialization",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase describes the Phase of StorageClusterInitialization This is used by OLM UI to provide status information to the user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions describes the state of the StorageClusterInitialization resource.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects created and maintained by this operator. Object references will be added to this list after they have been created AND found in the cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"storageClassesCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"cephObjectStoresCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"cephBlockPoolsCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"cephObjectStoreUsersCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"cephFilesystemsCreated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"errorMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_ocs_v1_StorageClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageClusterStatus defines the observed state of StorageCluster",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase describes the Phase of StorageCluster This is used by OLM UI to provide status information to the user",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions describes the state of the StorageCluster resource.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects created and maintained by this operator. Object references will be added to this list after they have been created AND found in the cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"nodeTopologies": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeTopologies is a list of topology labels on all nodes matching the StorageCluster's placement selector.",
							Ref:         ref("github.com/openshift/ocs-operator/pkg/apis/ocs/v1.NodeTopologyMap"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition", "github.com/openshift/ocs-operator/pkg/apis/ocs/v1.NodeTopologyMap", "k8s.io/api/core/v1.ObjectReference"},
	}
}

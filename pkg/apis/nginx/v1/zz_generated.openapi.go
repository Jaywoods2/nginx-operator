// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"nginx-operator/pkg/apis/nginx/v1.NginxService":       schema_pkg_apis_nginx_v1_NginxService(ref),
		"nginx-operator/pkg/apis/nginx/v1.NginxServiceSpec":   schema_pkg_apis_nginx_v1_NginxServiceSpec(ref),
		"nginx-operator/pkg/apis/nginx/v1.NginxServiceStatus": schema_pkg_apis_nginx_v1_NginxServiceStatus(ref),
	}
}

func schema_pkg_apis_nginx_v1_NginxService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NginxService is the Schema for the nginxservices API",
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
							Ref: ref("nginx-operator/pkg/apis/nginx/v1.NginxServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("nginx-operator/pkg/apis/nginx/v1.NginxServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "nginx-operator/pkg/apis/nginx/v1.NginxServiceSpec", "nginx-operator/pkg/apis/nginx/v1.NginxServiceStatus"},
	}
}

func schema_pkg_apis_nginx_v1_NginxServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NginxServiceSpec defines the desired state of NginxService",
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"envs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ServicePort"),
									},
								},
							},
						},
					},
				},
				Required: []string{"size", "image"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.ServicePort"},
	}
}

func schema_pkg_apis_nginx_v1_NginxServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NginxServiceStatus defines the observed state of NginxService",
				Properties: map[string]spec.Schema{
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "The generation observed by the deployment controller.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of non-terminated pods targeted by this deployment (their labels match the selector).",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"updatedReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of non-terminated pods targeted by this deployment that have the desired template spec.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"readyReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of ready pods targeted by this deployment.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"availableReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"unavailableReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Represents the latest available observations of a deployment's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/apps/v1.DeploymentCondition"),
									},
								},
							},
						},
					},
					"collisionCount": {
						SchemaProps: spec.SchemaProps{
							Description: "Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/apps/v1.DeploymentCondition"},
	}
}
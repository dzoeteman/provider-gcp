/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegistryObservation struct {

	// The URI of the created resource.
	BucketSelfLink *string `json:"bucketSelfLink,omitempty" tf:"bucket_self_link,omitempty"`

	// The name of the bucket that supports the Container Registry. In the form of artifacts.{project}.appspot.com or {location}.artifacts.{project}.appspot.com if location is specified.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RegistryParameters struct {

	// The location of the registry. One of ASIA, EU, US or not specified. See the official documentation for more information on registry locations.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Registry is the Schema for the Registrys API. Ensures the GCS bucket backing Google Container Registry exists.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistrySpec   `json:"spec"`
	Status            RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}

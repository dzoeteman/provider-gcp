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

type DefaultSupportedIdPConfigObservation struct {

	// an identifier for the resource with format projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the DefaultSupportedIdpConfig resource
	// The name of the DefaultSupportedIdpConfig resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DefaultSupportedIdPConfigParameters struct {

	// OAuth client ID
	// OAuth client ID
	// +kubebuilder:validation:Required
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// OAuth client secret
	// OAuth client secret
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// If this IDP allows the user to sign in
	// If this IDP allows the user to sign in
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// ID of the IDP. Possible values include:
	// ID of the IDP. Possible values include:
	//
	// * 'apple.com'
	//
	// * 'facebook.com'
	//
	// * 'gc.apple.com'
	//
	// * 'github.com'
	//
	// * 'google.com'
	//
	// * 'linkedin.com'
	//
	// * 'microsoft.com'
	//
	// * 'playgames.google.com'
	//
	// * 'twitter.com'
	//
	// * 'yahoo.com'
	// +kubebuilder:validation:Required
	IdPID *string `json:"idpId" tf:"idp_id,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DefaultSupportedIdPConfigSpec defines the desired state of DefaultSupportedIdPConfig
type DefaultSupportedIdPConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSupportedIdPConfigParameters `json:"forProvider"`
}

// DefaultSupportedIdPConfigStatus defines the observed state of DefaultSupportedIdPConfig.
type DefaultSupportedIdPConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSupportedIdPConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSupportedIdPConfig is the Schema for the DefaultSupportedIdPConfigs API. Configurations options for authenticating with a the standard set of Identity Toolkit-trusted IDPs.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DefaultSupportedIdPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSupportedIdPConfigSpec   `json:"spec"`
	Status            DefaultSupportedIdPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSupportedIdPConfigList contains a list of DefaultSupportedIdPConfigs
type DefaultSupportedIdPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSupportedIdPConfig `json:"items"`
}

// Repository type metadata.
var (
	DefaultSupportedIdPConfig_Kind             = "DefaultSupportedIdPConfig"
	DefaultSupportedIdPConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultSupportedIdPConfig_Kind}.String()
	DefaultSupportedIdPConfig_KindAPIVersion   = DefaultSupportedIdPConfig_Kind + "." + CRDGroupVersion.String()
	DefaultSupportedIdPConfig_GroupVersionKind = CRDGroupVersion.WithKind(DefaultSupportedIdPConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultSupportedIdPConfig{}, &DefaultSupportedIdPConfigList{})
}
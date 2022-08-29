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

type WebTypeAppEngineIAMMemberConditionObservation struct {
}

type WebTypeAppEngineIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type WebTypeAppEngineIAMMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WebTypeAppEngineIAMMemberParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/appengine/v1beta1.Application
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	AppIDRef *v1.Reference `json:"appIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AppIDSelector *v1.Selector `json:"appIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Condition []WebTypeAppEngineIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// WebTypeAppEngineIAMMemberSpec defines the desired state of WebTypeAppEngineIAMMember
type WebTypeAppEngineIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebTypeAppEngineIAMMemberParameters `json:"forProvider"`
}

// WebTypeAppEngineIAMMemberStatus defines the observed state of WebTypeAppEngineIAMMember.
type WebTypeAppEngineIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebTypeAppEngineIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WebTypeAppEngineIAMMember is the Schema for the WebTypeAppEngineIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type WebTypeAppEngineIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebTypeAppEngineIAMMemberSpec   `json:"spec"`
	Status            WebTypeAppEngineIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebTypeAppEngineIAMMemberList contains a list of WebTypeAppEngineIAMMembers
type WebTypeAppEngineIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebTypeAppEngineIAMMember `json:"items"`
}

// Repository type metadata.
var (
	WebTypeAppEngineIAMMember_Kind             = "WebTypeAppEngineIAMMember"
	WebTypeAppEngineIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebTypeAppEngineIAMMember_Kind}.String()
	WebTypeAppEngineIAMMember_KindAPIVersion   = WebTypeAppEngineIAMMember_Kind + "." + CRDGroupVersion.String()
	WebTypeAppEngineIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(WebTypeAppEngineIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&WebTypeAppEngineIAMMember{}, &WebTypeAppEngineIAMMemberList{})
}

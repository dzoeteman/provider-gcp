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

type AppConnectorInitParameters struct {

	// An arbitrary user-provided name for the AppConnector.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo []PrincipalInfoInitParameters `json:"principalInfo,omitempty" tf:"principal_info,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AppConnectorObservation struct {

	// An arbitrary user-provided name for the AppConnector.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/appConnectors/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user provided metadata.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	PrincipalInfo []PrincipalInfoObservation `json:"principalInfo,omitempty" tf:"principal_info,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppConnector.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Represents the different states of a AppConnector.
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type AppConnectorParameters struct {

	// An arbitrary user-provided name for the AppConnector.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Principal information about the Identity of the AppConnector.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PrincipalInfo []PrincipalInfoParameters `json:"principalInfo,omitempty" tf:"principal_info,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppConnector.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type PrincipalInfoInitParameters struct {

	// ServiceAccount represents a GCP service account.
	// Structure is documented below.
	ServiceAccount []ServiceAccountInitParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type PrincipalInfoObservation struct {

	// ServiceAccount represents a GCP service account.
	// Structure is documented below.
	ServiceAccount []ServiceAccountObservation `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type PrincipalInfoParameters struct {

	// ServiceAccount represents a GCP service account.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServiceAccount []ServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type ServiceAccountInitParameters struct {
}

type ServiceAccountObservation struct {

	// Email address of the service account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`
}

type ServiceAccountParameters struct {

	// Email address of the service account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("email",true)
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// Reference to a ServiceAccount in cloudplatform to populate email.
	// +kubebuilder:validation:Optional
	EmailRef *v1.Reference `json:"emailRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in cloudplatform to populate email.
	// +kubebuilder:validation:Optional
	EmailSelector *v1.Selector `json:"emailSelector,omitempty" tf:"-"`
}

// AppConnectorSpec defines the desired state of AppConnector
type AppConnectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppConnectorParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AppConnectorInitParameters `json:"initProvider,omitempty"`
}

// AppConnectorStatus defines the observed state of AppConnector.
type AppConnectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppConnectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppConnector is the Schema for the AppConnectors API. A BeyondCorp AppConnector resource represents an application facing component deployed proximal to and with direct access to the application instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AppConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principalInfo) || has(self.initProvider.principalInfo)",message="principalInfo is a required parameter"
	Spec   AppConnectorSpec   `json:"spec"`
	Status AppConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppConnectorList contains a list of AppConnectors
type AppConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppConnector `json:"items"`
}

// Repository type metadata.
var (
	AppConnector_Kind             = "AppConnector"
	AppConnector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppConnector_Kind}.String()
	AppConnector_KindAPIVersion   = AppConnector_Kind + "." + CRDGroupVersion.String()
	AppConnector_GroupVersionKind = CRDGroupVersion.WithKind(AppConnector_Kind)
)

func init() {
	SchemeBuilder.Register(&AppConnector{}, &AppConnectorList{})
}

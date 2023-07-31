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

type AnalyticsHubDataExchangeInitParameters struct {

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Description of the data exchange.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the data exchange.
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// Base64 encoded image representing the data exchange.
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AnalyticsHubDataExchangeObservation struct {

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Description of the data exchange.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the data exchange.
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Base64 encoded image representing the data exchange.
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// Number of listings contained in the data exchange.
	ListingCount *float64 `json:"listingCount,omitempty" tf:"listing_count,omitempty"`

	// The name of the location this data exchange.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name of the data exchange, for example:
	// "projects/myproject/locations/US/dataExchanges/123"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Email or URL of the primary point of contact of the data exchange.
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type AnalyticsHubDataExchangeParameters struct {

	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	// +kubebuilder:validation:Optional
	DataExchangeID *string `json:"dataExchangeId,omitempty" tf:"data_exchange_id,omitempty"`

	// Description of the data exchange.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and must not start or end with spaces.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Documentation describing the data exchange.
	// +kubebuilder:validation:Optional
	Documentation *string `json:"documentation,omitempty" tf:"documentation,omitempty"`

	// Base64 encoded image representing the data exchange.
	// +kubebuilder:validation:Optional
	Icon *string `json:"icon,omitempty" tf:"icon,omitempty"`

	// The name of the location this data exchange.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Email or URL of the primary point of contact of the data exchange.
	// +kubebuilder:validation:Optional
	PrimaryContact *string `json:"primaryContact,omitempty" tf:"primary_contact,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// AnalyticsHubDataExchangeSpec defines the desired state of AnalyticsHubDataExchange
type AnalyticsHubDataExchangeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsHubDataExchangeParameters `json:"forProvider"`
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
	InitProvider AnalyticsHubDataExchangeInitParameters `json:"initProvider,omitempty"`
}

// AnalyticsHubDataExchangeStatus defines the observed state of AnalyticsHubDataExchange.
type AnalyticsHubDataExchangeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsHubDataExchangeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsHubDataExchange is the Schema for the AnalyticsHubDataExchanges API. A Bigquery Analytics Hub data exchange
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AnalyticsHubDataExchange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataExchangeId) || has(self.initProvider.dataExchangeId)",message="dataExchangeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   AnalyticsHubDataExchangeSpec   `json:"spec"`
	Status AnalyticsHubDataExchangeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsHubDataExchangeList contains a list of AnalyticsHubDataExchanges
type AnalyticsHubDataExchangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsHubDataExchange `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsHubDataExchange_Kind             = "AnalyticsHubDataExchange"
	AnalyticsHubDataExchange_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsHubDataExchange_Kind}.String()
	AnalyticsHubDataExchange_KindAPIVersion   = AnalyticsHubDataExchange_Kind + "." + CRDGroupVersion.String()
	AnalyticsHubDataExchange_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsHubDataExchange_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsHubDataExchange{}, &AnalyticsHubDataExchangeList{})
}

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

type DatasetIAMPolicyInitParameters struct {

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DatasetIAMPolicyObservation struct {

	// The dataset ID.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// (Computed) The etag of the dataset's IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DatasetIAMPolicyParameters struct {

	// The dataset ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The policy data generated by
	// a google_iam_policy data source.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// DatasetIAMPolicySpec defines the desired state of DatasetIAMPolicy
type DatasetIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasetIAMPolicyParameters `json:"forProvider"`
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
	InitProvider DatasetIAMPolicyInitParameters `json:"initProvider,omitempty"`
}

// DatasetIAMPolicyStatus defines the observed state of DatasetIAMPolicy.
type DatasetIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasetIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMPolicy is the Schema for the DatasetIAMPolicys API. Collection of resources to manage IAM policy for a BigQuery dataset.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DatasetIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData) || has(self.initProvider.policyData)",message="policyData is a required parameter"
	Spec   DatasetIAMPolicySpec   `json:"spec"`
	Status DatasetIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasetIAMPolicyList contains a list of DatasetIAMPolicys
type DatasetIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasetIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	DatasetIAMPolicy_Kind             = "DatasetIAMPolicy"
	DatasetIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatasetIAMPolicy_Kind}.String()
	DatasetIAMPolicy_KindAPIVersion   = DatasetIAMPolicy_Kind + "." + CRDGroupVersion.String()
	DatasetIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DatasetIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DatasetIAMPolicy{}, &DatasetIAMPolicyList{})
}

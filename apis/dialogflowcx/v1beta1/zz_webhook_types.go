// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type GenericWebServiceInitParameters struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GenericWebServiceObservation struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GenericWebServiceParameters struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	// +kubebuilder:validation:Optional
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	// +kubebuilder:validation:Optional
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ServiceDirectoryGenericWebServiceInitParameters struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ServiceDirectoryGenericWebServiceObservation struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ServiceDirectoryGenericWebServiceParameters struct {

	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	// +kubebuilder:validation:Optional
	AllowedCACerts []*string `json:"allowedCaCerts,omitempty" tf:"allowed_ca_certs,omitempty"`

	// The HTTP request headers to send together with webhook requests.
	// +kubebuilder:validation:Optional
	RequestHeaders map[string]*string `json:"requestHeaders,omitempty" tf:"request_headers,omitempty"`

	// Whether to use speech adaptation for speech recognition.
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type ServiceDirectoryInitParameters struct {

	// The name of Service Directory service.
	// Structure is documented below.
	GenericWebService []ServiceDirectoryGenericWebServiceInitParameters `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// The name of Service Directory service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceDirectoryObservation struct {

	// The name of Service Directory service.
	// Structure is documented below.
	GenericWebService []ServiceDirectoryGenericWebServiceObservation `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// The name of Service Directory service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ServiceDirectoryParameters struct {

	// The name of Service Directory service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GenericWebService []ServiceDirectoryGenericWebServiceParameters `json:"genericWebService" tf:"generic_web_service,omitempty"`

	// The name of Service Directory service.
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`
}

type WebhookInitParameters struct {

	// Indicates whether the webhook is disabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The human-readable name of the webhook, unique within the agent.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService []GenericWebServiceInitParameters `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory []ServiceDirectoryInitParameters `json:"serviceDirectory,omitempty" tf:"service_directory,omitempty"`

	// Webhook execution timeout.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type WebhookObservation struct {

	// Indicates whether the webhook is disabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The human-readable name of the webhook, unique within the agent.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService []GenericWebServiceObservation `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// an identifier for the resource with format {{parent}}/webhooks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The unique identifier of the webhook.
	// Format: projects//locations//agents//webhooks/.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The agent to create a webhook for.
	// Format: projects//locations//agents/.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory []ServiceDirectoryObservation `json:"serviceDirectory,omitempty" tf:"service_directory,omitempty"`

	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects//locations//agents//flows/.
	StartFlow *string `json:"startFlow,omitempty" tf:"start_flow,omitempty"`

	// Webhook execution timeout.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type WebhookParameters struct {

	// Indicates whether the webhook is disabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The human-readable name of the webhook, unique within the agent.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +kubebuilder:validation:Optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	// +kubebuilder:validation:Optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Configuration for a generic web service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GenericWebService []GenericWebServiceParameters `json:"genericWebService,omitempty" tf:"generic_web_service,omitempty"`

	// The agent to create a webhook for.
	// Format: projects//locations//agents/.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dialogflowcx/v1beta1.Agent
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Reference to a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentRef *v1.Reference `json:"parentRef,omitempty" tf:"-"`

	// Selector for a Agent in dialogflowcx to populate parent.
	// +kubebuilder:validation:Optional
	ParentSelector *v1.Selector `json:"parentSelector,omitempty" tf:"-"`

	// Name of the SecuritySettings reference for the agent. Format: projects//locations//securitySettings/.
	// +kubebuilder:validation:Optional
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Configuration for a Service Directory service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ServiceDirectory []ServiceDirectoryParameters `json:"serviceDirectory,omitempty" tf:"service_directory,omitempty"`

	// Webhook execution timeout.
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

// WebhookSpec defines the desired state of Webhook
type WebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider WebhookInitParameters `json:"initProvider,omitempty"`
}

// WebhookStatus defines the observed state of Webhook.
type WebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Webhook is the Schema for the Webhooks API. Webhooks host the developer's business logic.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhooks
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}

// Repository type metadata.
var (
	Webhook_Kind             = "Webhook"
	Webhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Webhook_Kind}.String()
	Webhook_KindAPIVersion   = Webhook_Kind + "." + CRDGroupVersion.String()
	Webhook_GroupVersionKind = CRDGroupVersion.WithKind(Webhook_Kind)
)

func init() {
	SchemeBuilder.Register(&Webhook{}, &WebhookList{})
}

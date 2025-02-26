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

type IdPCertificatesInitParameters struct {
}

type IdPCertificatesObservation struct {
}

type IdPCertificatesParameters struct {

	// (Output)
	// The x509 certificate
	// +kubebuilder:validation:Optional
	X509CertificateSecretRef *v1.SecretKeySelector `json:"x509CertificateSecretRef,omitempty" tf:"-"`
}

type IdPConfigInitParameters struct {

	// The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// Structure is documented below.
	IdPCertificates []IdPCertificatesInitParameters `json:"idpCertificates,omitempty" tf:"idp_certificates,omitempty"`

	// Unique identifier for all SAML entities
	IdPEntityID *string `json:"idpEntityId,omitempty" tf:"idp_entity_id,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	SignRequest *bool `json:"signRequest,omitempty" tf:"sign_request,omitempty"`

	// URL to send Authentication request to.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`
}

type IdPConfigObservation struct {

	// The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// Structure is documented below.
	IdPCertificates []IdPCertificatesParameters `json:"idpCertificates,omitempty" tf:"idp_certificates,omitempty"`

	// Unique identifier for all SAML entities
	IdPEntityID *string `json:"idpEntityId,omitempty" tf:"idp_entity_id,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	SignRequest *bool `json:"signRequest,omitempty" tf:"sign_request,omitempty"`

	// URL to send Authentication request to.
	SsoURL *string `json:"ssoUrl,omitempty" tf:"sso_url,omitempty"`
}

type IdPConfigParameters struct {

	// The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	IdPCertificates []IdPCertificatesParameters `json:"idpCertificates" tf:"idp_certificates,omitempty"`

	// Unique identifier for all SAML entities
	// +kubebuilder:validation:Optional
	IdPEntityID *string `json:"idpEntityId" tf:"idp_entity_id,omitempty"`

	// Indicates if outbounding SAMLRequest should be signed.
	// +kubebuilder:validation:Optional
	SignRequest *bool `json:"signRequest,omitempty" tf:"sign_request,omitempty"`

	// URL to send Authentication request to.
	// +kubebuilder:validation:Optional
	SsoURL *string `json:"ssoUrl" tf:"sso_url,omitempty"`
}

type InboundSAMLConfigInitParameters struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdPConfig []IdPConfigInitParameters `json:"idpConfig,omitempty" tf:"idp_config,omitempty"`

	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig []SpConfigInitParameters `json:"spConfig,omitempty" tf:"sp_config,omitempty"`
}

type InboundSAMLConfigObservation struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// an identifier for the resource with format projects/{{project}}/inboundSamlConfigs/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	IdPConfig []IdPConfigObservation `json:"idpConfig,omitempty" tf:"idp_config,omitempty"`

	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	SpConfig []SpConfigObservation `json:"spConfig,omitempty" tf:"sp_config,omitempty"`
}

type InboundSAMLConfigParameters struct {

	// Human friendly display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// SAML IdP configuration when the project acts as the relying party
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	IdPConfig []IdPConfigParameters `json:"idpConfig,omitempty" tf:"idp_config,omitempty"`

	// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
	// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
	// alphanumeric character, and have at least 2 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	// and accept an authentication assertion issued by a SAML identity provider.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SpConfig []SpConfigParameters `json:"spConfig,omitempty" tf:"sp_config,omitempty"`
}

type SpCertificatesInitParameters struct {
}

type SpCertificatesObservation struct {

	// (Output)
	// The x509 certificate
	X509Certificate *string `json:"x509Certificate,omitempty" tf:"x509_certificate,omitempty"`
}

type SpCertificatesParameters struct {
}

type SpConfigInitParameters struct {

	// Callback URI where responses from IDP are handled. Must start with https://.
	CallbackURI *string `json:"callbackUri,omitempty" tf:"callback_uri,omitempty"`

	// Unique identifier for all SAML entities.
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

type SpConfigObservation struct {

	// Callback URI where responses from IDP are handled. Must start with https://.
	CallbackURI *string `json:"callbackUri,omitempty" tf:"callback_uri,omitempty"`

	// (Output)
	// The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
	// Structure is documented below.
	SpCertificates []SpCertificatesObservation `json:"spCertificates,omitempty" tf:"sp_certificates,omitempty"`

	// Unique identifier for all SAML entities.
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

type SpConfigParameters struct {

	// Callback URI where responses from IDP are handled. Must start with https://.
	// +kubebuilder:validation:Optional
	CallbackURI *string `json:"callbackUri,omitempty" tf:"callback_uri,omitempty"`

	// Unique identifier for all SAML entities.
	// +kubebuilder:validation:Optional
	SpEntityID *string `json:"spEntityId,omitempty" tf:"sp_entity_id,omitempty"`
}

// InboundSAMLConfigSpec defines the desired state of InboundSAMLConfig
type InboundSAMLConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InboundSAMLConfigParameters `json:"forProvider"`
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
	InitProvider InboundSAMLConfigInitParameters `json:"initProvider,omitempty"`
}

// InboundSAMLConfigStatus defines the observed state of InboundSAMLConfig.
type InboundSAMLConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InboundSAMLConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InboundSAMLConfig is the Schema for the InboundSAMLConfigs API. Inbound SAML configuration for a Identity Toolkit project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InboundSAMLConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.idpConfig) || (has(self.initProvider) && has(self.initProvider.idpConfig))",message="spec.forProvider.idpConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.spConfig) || (has(self.initProvider) && has(self.initProvider.spConfig))",message="spec.forProvider.spConfig is a required parameter"
	Spec   InboundSAMLConfigSpec   `json:"spec"`
	Status InboundSAMLConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InboundSAMLConfigList contains a list of InboundSAMLConfigs
type InboundSAMLConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InboundSAMLConfig `json:"items"`
}

// Repository type metadata.
var (
	InboundSAMLConfig_Kind             = "InboundSAMLConfig"
	InboundSAMLConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InboundSAMLConfig_Kind}.String()
	InboundSAMLConfig_KindAPIVersion   = InboundSAMLConfig_Kind + "." + CRDGroupVersion.String()
	InboundSAMLConfig_GroupVersionKind = CRDGroupVersion.WithKind(InboundSAMLConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&InboundSAMLConfig{}, &InboundSAMLConfigList{})
}

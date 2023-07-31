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

type AccessRoleInitParameters struct {

	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`
}

type AccessRoleObservation struct {

	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`

	// (Output)
	// A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's AWS IAM Role.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`
}

type AccessRoleParameters struct {

	// The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.
	// +kubebuilder:validation:Optional
	IAMRoleID *string `json:"iamRoleId,omitempty" tf:"iam_role_id,omitempty"`
}

type AwsInitParameters struct {

	// Authentication using Google owned service account to assume into customer's AWS IAM Role.
	// Structure is documented below.
	AccessRole []AccessRoleInitParameters `json:"accessRole,omitempty" tf:"access_role,omitempty"`
}

type AwsObservation struct {

	// Authentication using Google owned service account to assume into customer's AWS IAM Role.
	// Structure is documented below.
	AccessRole []AccessRoleObservation `json:"accessRole,omitempty" tf:"access_role,omitempty"`
}

type AwsParameters struct {

	// Authentication using Google owned service account to assume into customer's AWS IAM Role.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AccessRole []AccessRoleParameters `json:"accessRole,omitempty" tf:"access_role,omitempty"`
}

type AzureInitParameters struct {

	// The id of customer's directory that host the data.
	CustomerTenantID *string `json:"customerTenantId,omitempty" tf:"customer_tenant_id,omitempty"`

	// The Azure Application (client) ID where the federated credentials will be hosted.
	FederatedApplicationClientID *string `json:"federatedApplicationClientId,omitempty" tf:"federated_application_client_id,omitempty"`
}

type AzureObservation struct {

	// (Output)
	// The name of the Azure Active Directory Application.
	Application *string `json:"application,omitempty" tf:"application,omitempty"`

	// (Output)
	// The client id of the Azure Active Directory Application.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The id of customer's directory that host the data.
	CustomerTenantID *string `json:"customerTenantId,omitempty" tf:"customer_tenant_id,omitempty"`

	// The Azure Application (client) ID where the federated credentials will be hosted.
	FederatedApplicationClientID *string `json:"federatedApplicationClientId,omitempty" tf:"federated_application_client_id,omitempty"`

	// (Output)
	// A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's Azure Active Directory Application.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// (Output)
	// The object id of the Azure Active Directory Application.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// (Output)
	// The URL user will be redirected to after granting consent during connection setup.
	RedirectURI *string `json:"redirectUri,omitempty" tf:"redirect_uri,omitempty"`
}

type AzureParameters struct {

	// The id of customer's directory that host the data.
	// +kubebuilder:validation:Optional
	CustomerTenantID *string `json:"customerTenantId,omitempty" tf:"customer_tenant_id,omitempty"`

	// The Azure Application (client) ID where the federated credentials will be hosted.
	// +kubebuilder:validation:Optional
	FederatedApplicationClientID *string `json:"federatedApplicationClientId,omitempty" tf:"federated_application_client_id,omitempty"`
}

type CloudResourceInitParameters struct {
}

type CloudResourceObservation struct {

	// (Output)
	// The account ID of the service created for the purpose of this connection.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type CloudResourceParameters struct {
}

type CloudSQLInitParameters struct {

	// Cloud SQL properties.
	// Structure is documented below.
	Credential []CredentialInitParameters `json:"credential,omitempty" tf:"credential,omitempty"`

	// Type of the Cloud SQL database.
	// Possible values are: DATABASE_TYPE_UNSPECIFIED, POSTGRES, MYSQL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CloudSQLObservation struct {

	// Cloud SQL properties.
	// Structure is documented below.
	Credential []CredentialObservation `json:"credential,omitempty" tf:"credential,omitempty"`

	// Database name.
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Cloud SQL instance ID in the form project:location:instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// (Output)
	// When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Type of the Cloud SQL database.
	// Possible values are: DATABASE_TYPE_UNSPECIFIED, POSTGRES, MYSQL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CloudSQLParameters struct {

	// Cloud SQL properties.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Credential []CredentialParameters `json:"credential,omitempty" tf:"credential,omitempty"`

	// Database name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta1.Database
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// Reference to a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseRef *v1.Reference `json:"databaseRef,omitempty" tf:"-"`

	// Selector for a Database in sql to populate database.
	// +kubebuilder:validation:Optional
	DatabaseSelector *v1.Selector `json:"databaseSelector,omitempty" tf:"-"`

	// Cloud SQL instance ID in the form project:location:instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta1.DatabaseInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("connection_name",true)
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a DatabaseInstance in sql to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a DatabaseInstance in sql to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Type of the Cloud SQL database.
	// Possible values are: DATABASE_TYPE_UNSPECIFIED, POSTGRES, MYSQL.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CloudSpannerInitParameters struct {

	// Cloud Spanner database in the form `project/instance/database'
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner
	UseParallelism *bool `json:"useParallelism,omitempty" tf:"use_parallelism,omitempty"`

	// If the serverless analytics service should be used to read data from Cloud Spanner. useParallelism must be set when using serverless analytics
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty" tf:"use_serverless_analytics,omitempty"`
}

type CloudSpannerObservation struct {

	// Cloud Spanner database in the form `project/instance/database'
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner
	UseParallelism *bool `json:"useParallelism,omitempty" tf:"use_parallelism,omitempty"`

	// If the serverless analytics service should be used to read data from Cloud Spanner. useParallelism must be set when using serverless analytics
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty" tf:"use_serverless_analytics,omitempty"`
}

type CloudSpannerParameters struct {

	// Cloud Spanner database in the form `project/instance/database'
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner
	// +kubebuilder:validation:Optional
	UseParallelism *bool `json:"useParallelism,omitempty" tf:"use_parallelism,omitempty"`

	// If the serverless analytics service should be used to read data from Cloud Spanner. useParallelism must be set when using serverless analytics
	// +kubebuilder:validation:Optional
	UseServerlessAnalytics *bool `json:"useServerlessAnalytics,omitempty" tf:"use_serverless_analytics,omitempty"`
}

type ConnectionInitParameters struct {

	// Connection properties specific to Amazon Web Services.
	// Structure is documented below.
	Aws []AwsInitParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// Container for connection properties specific to Azure.
	// Structure is documented below.
	Azure []AzureInitParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// Container for connection properties for delegation of access to GCP resources.
	// Structure is documented below.
	CloudResource []CloudResourceInitParameters `json:"cloudResource,omitempty" tf:"cloud_resource,omitempty"`

	// Connection properties specific to the Cloud SQL.
	// Structure is documented below.
	CloudSQL []CloudSQLInitParameters `json:"cloudSql,omitempty" tf:"cloud_sql,omitempty"`

	// Connection properties specific to Cloud Spanner
	// Structure is documented below.
	CloudSpanner []CloudSpannerInitParameters `json:"cloudSpanner,omitempty" tf:"cloud_spanner,omitempty"`

	// Optional connection id that should be assigned to the created connection.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// A descriptive description for the connection
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the connection
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	// Spanner Connections same as spanner region
	// AWS allowed regions are aws-us-east-1
	// Azure allowed regions are azure-eastus2
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ConnectionObservation struct {

	// Connection properties specific to Amazon Web Services.
	// Structure is documented below.
	Aws []AwsObservation `json:"aws,omitempty" tf:"aws,omitempty"`

	// Container for connection properties specific to Azure.
	// Structure is documented below.
	Azure []AzureObservation `json:"azure,omitempty" tf:"azure,omitempty"`

	// Container for connection properties for delegation of access to GCP resources.
	// Structure is documented below.
	CloudResource []CloudResourceObservation `json:"cloudResource,omitempty" tf:"cloud_resource,omitempty"`

	// Connection properties specific to the Cloud SQL.
	// Structure is documented below.
	CloudSQL []CloudSQLObservation `json:"cloudSql,omitempty" tf:"cloud_sql,omitempty"`

	// Connection properties specific to Cloud Spanner
	// Structure is documented below.
	CloudSpanner []CloudSpannerObservation `json:"cloudSpanner,omitempty" tf:"cloud_spanner,omitempty"`

	// Optional connection id that should be assigned to the created connection.
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// A descriptive description for the connection
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the connection
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// True if the connection has credential assigned.
	HasCredential *bool `json:"hasCredential,omitempty" tf:"has_credential,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	// Spanner Connections same as spanner region
	// AWS allowed regions are aws-us-east-1
	// Azure allowed regions are azure-eastus2
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name of the connection in the form of:
	// "projects/{project_id}/locations/{location_id}/connections/{connectionId}"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ConnectionParameters struct {

	// Connection properties specific to Amazon Web Services.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Aws []AwsParameters `json:"aws,omitempty" tf:"aws,omitempty"`

	// Container for connection properties specific to Azure.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Azure []AzureParameters `json:"azure,omitempty" tf:"azure,omitempty"`

	// Container for connection properties for delegation of access to GCP resources.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudResource []CloudResourceParameters `json:"cloudResource,omitempty" tf:"cloud_resource,omitempty"`

	// Connection properties specific to the Cloud SQL.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudSQL []CloudSQLParameters `json:"cloudSql,omitempty" tf:"cloud_sql,omitempty"`

	// Connection properties specific to Cloud Spanner
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudSpanner []CloudSpannerParameters `json:"cloudSpanner,omitempty" tf:"cloud_spanner,omitempty"`

	// Optional connection id that should be assigned to the created connection.
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// A descriptive description for the connection
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A descriptive name for the connection
	// +kubebuilder:validation:Optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`

	// The geographic location where the connection should reside.
	// Cloud SQL instance must be in the same location as the connection
	// with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	// Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	// Spanner Connections same as spanner region
	// AWS allowed regions are aws-us-east-1
	// Azure allowed regions are azure-eastus2
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type CredentialInitParameters struct {
}

type CredentialObservation struct {

	// Username for database.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialParameters struct {

	// Password for database.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Username for database.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/sql/v1beta1.User
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a User in sql to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. A connection allows BigQuery connections to external data sources.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}

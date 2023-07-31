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

type BucketObjectInitParameters struct {

	// Cache-Control
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Content-Disposition of the object data.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content-Encoding of the object data.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Content-Language of the object data.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). [Google documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption []CustomerEncryptionInitParameters `json:"customerEncryption,omitempty" tf:"customer_encryption,omitempty"`

	// MD5 hash of the data, encoded using base64. This field is not present for composite objects. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
	DetectMd5Hash *string `json:"detectMd5Hash,omitempty" tf:"detect_md5hash,omitempty"`

	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold *bool `json:"eventBasedHold,omitempty" tf:"event_based_hold,omitempty"`

	// The resource name of the Cloud KMS key that will be used to encrypt the object.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User-provided metadata, in key/value pairs.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A path to the data you want to upload. Must be defined
	// if content is not.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The StorageClass of the new bucket object.
	// Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default
	// storage class or to a standard class.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold *bool `json:"temporaryHold,omitempty" tf:"temporary_hold,omitempty"`
}

type BucketObjectObservation struct {

	// The name of the containing bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Cache-Control
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Content-Disposition of the object data.
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content-Encoding of the object data.
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Content-Language of the object data.
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32C *string `json:"crc32c,omitempty" tf:"crc32c,omitempty"`

	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). [Google documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption []CustomerEncryptionObservation `json:"customerEncryption,omitempty" tf:"customer_encryption,omitempty"`

	// MD5 hash of the data, encoded using base64. This field is not present for composite objects. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
	DetectMd5Hash *string `json:"detectMd5Hash,omitempty" tf:"detect_md5hash,omitempty"`

	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold *bool `json:"eventBasedHold,omitempty" tf:"event_based_hold,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the Cloud KMS key that will be used to encrypt the object.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5Hash *string `json:"md5hash,omitempty" tf:"md5hash,omitempty"`

	// (Computed) A url reference to download this object.
	MediaLink *string `json:"mediaLink,omitempty" tf:"media_link,omitempty"`

	// User-provided metadata, in key/value pairs.
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed) The name of the object. Use this field in interpolations with google_storage_object_acl to recreate
	// google_storage_object_acl resources when your google_storage_bucket_object is recreated.
	OutputName *string `json:"outputName,omitempty" tf:"output_name,omitempty"`

	// (Computed) A url reference to this object.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A path to the data you want to upload. Must be defined
	// if content is not.
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The StorageClass of the new bucket object.
	// Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default
	// storage class or to a standard class.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold *bool `json:"temporaryHold,omitempty" tf:"temporary_hold,omitempty"`
}

type BucketObjectParameters struct {

	// The name of the containing bucket.
	// +crossplane:generate:reference:type=Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Cache-Control
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view the raw contents of the object, please define an output.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Content-Disposition of the object data.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Content-Encoding of the object data.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// Content-Language of the object data.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). [Google documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CustomerEncryption []CustomerEncryptionParameters `json:"customerEncryption,omitempty" tf:"customer_encryption,omitempty"`

	// MD5 hash of the data, encoded using base64. This field is not present for composite objects. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
	// +kubebuilder:validation:Optional
	DetectMd5Hash *string `json:"detectMd5Hash,omitempty" tf:"detect_md5hash,omitempty"`

	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	// +kubebuilder:validation:Optional
	EventBasedHold *bool `json:"eventBasedHold,omitempty" tf:"event_based_hold,omitempty"`

	// The resource name of the Cloud KMS key that will be used to encrypt the object.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// User-provided metadata, in key/value pairs.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A path to the data you want to upload. Must be defined
	// if content is not.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// The StorageClass of the new bucket object.
	// Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. If not provided, this defaults to the bucket's default
	// storage class or to a standard class.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites.
	// +kubebuilder:validation:Optional
	TemporaryHold *bool `json:"temporaryHold,omitempty" tf:"temporary_hold,omitempty"`
}

type CustomerEncryptionInitParameters struct {

	// Encryption algorithm. Default: AES256
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`
}

type CustomerEncryptionObservation struct {

	// Encryption algorithm. Default: AES256
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`
}

type CustomerEncryptionParameters struct {

	// Encryption algorithm. Default: AES256
	// +kubebuilder:validation:Optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm,omitempty"`

	// Base64 encoded Customer-Supplied Encryption Key.
	// +kubebuilder:validation:Required
	EncryptionKeySecretRef v1.SecretKeySelector `json:"encryptionKeySecretRef" tf:"-"`
}

// BucketObjectSpec defines the desired state of BucketObject
type BucketObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketObjectParameters `json:"forProvider"`
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
	InitProvider BucketObjectInitParameters `json:"initProvider,omitempty"`
}

// BucketObjectStatus defines the observed state of BucketObject.
type BucketObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObject is the Schema for the BucketObjects API. Creates a new object inside a specified bucket
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   BucketObjectSpec   `json:"spec"`
	Status BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketObjectList contains a list of BucketObjects
type BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketObject `json:"items"`
}

// Repository type metadata.
var (
	BucketObject_Kind             = "BucketObject"
	BucketObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketObject_Kind}.String()
	BucketObject_KindAPIVersion   = BucketObject_Kind + "." + CRDGroupVersion.String()
	BucketObject_GroupVersionKind = CRDGroupVersion.WithKind(BucketObject_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketObject{}, &BucketObjectList{})
}

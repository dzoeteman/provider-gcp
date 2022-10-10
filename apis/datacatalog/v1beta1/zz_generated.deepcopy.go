//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigqueryDateShardedSpecObservation) DeepCopyInto(out *BigqueryDateShardedSpecObservation) {
	*out = *in
	if in.Dataset != nil {
		in, out := &in.Dataset, &out.Dataset
		*out = new(string)
		**out = **in
	}
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(float64)
		**out = **in
	}
	if in.TablePrefix != nil {
		in, out := &in.TablePrefix, &out.TablePrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigqueryDateShardedSpecObservation.
func (in *BigqueryDateShardedSpecObservation) DeepCopy() *BigqueryDateShardedSpecObservation {
	if in == nil {
		return nil
	}
	out := new(BigqueryDateShardedSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigqueryDateShardedSpecParameters) DeepCopyInto(out *BigqueryDateShardedSpecParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigqueryDateShardedSpecParameters.
func (in *BigqueryDateShardedSpecParameters) DeepCopy() *BigqueryDateShardedSpecParameters {
	if in == nil {
		return nil
	}
	out := new(BigqueryDateShardedSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigqueryTableSpecObservation) DeepCopyInto(out *BigqueryTableSpecObservation) {
	*out = *in
	if in.TableSourceType != nil {
		in, out := &in.TableSourceType, &out.TableSourceType
		*out = new(string)
		**out = **in
	}
	if in.TableSpec != nil {
		in, out := &in.TableSpec, &out.TableSpec
		*out = make([]TableSpecObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ViewSpec != nil {
		in, out := &in.ViewSpec, &out.ViewSpec
		*out = make([]ViewSpecObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigqueryTableSpecObservation.
func (in *BigqueryTableSpecObservation) DeepCopy() *BigqueryTableSpecObservation {
	if in == nil {
		return nil
	}
	out := new(BigqueryTableSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigqueryTableSpecParameters) DeepCopyInto(out *BigqueryTableSpecParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigqueryTableSpecParameters.
func (in *BigqueryTableSpecParameters) DeepCopy() *BigqueryTableSpecParameters {
	if in == nil {
		return nil
	}
	out := new(BigqueryTableSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Entry) DeepCopyInto(out *Entry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entry.
func (in *Entry) DeepCopy() *Entry {
	if in == nil {
		return nil
	}
	out := new(Entry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Entry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroup) DeepCopyInto(out *EntryGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroup.
func (in *EntryGroup) DeepCopy() *EntryGroup {
	if in == nil {
		return nil
	}
	out := new(EntryGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntryGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroupList) DeepCopyInto(out *EntryGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EntryGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroupList.
func (in *EntryGroupList) DeepCopy() *EntryGroupList {
	if in == nil {
		return nil
	}
	out := new(EntryGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntryGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroupObservation) DeepCopyInto(out *EntryGroupObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroupObservation.
func (in *EntryGroupObservation) DeepCopy() *EntryGroupObservation {
	if in == nil {
		return nil
	}
	out := new(EntryGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroupParameters) DeepCopyInto(out *EntryGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EntryGroupID != nil {
		in, out := &in.EntryGroupID, &out.EntryGroupID
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroupParameters.
func (in *EntryGroupParameters) DeepCopy() *EntryGroupParameters {
	if in == nil {
		return nil
	}
	out := new(EntryGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroupSpec) DeepCopyInto(out *EntryGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroupSpec.
func (in *EntryGroupSpec) DeepCopy() *EntryGroupSpec {
	if in == nil {
		return nil
	}
	out := new(EntryGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryGroupStatus) DeepCopyInto(out *EntryGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryGroupStatus.
func (in *EntryGroupStatus) DeepCopy() *EntryGroupStatus {
	if in == nil {
		return nil
	}
	out := new(EntryGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryList) DeepCopyInto(out *EntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Entry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryList.
func (in *EntryList) DeepCopy() *EntryList {
	if in == nil {
		return nil
	}
	out := new(EntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryObservation) DeepCopyInto(out *EntryObservation) {
	*out = *in
	if in.BigqueryDateShardedSpec != nil {
		in, out := &in.BigqueryDateShardedSpec, &out.BigqueryDateShardedSpec
		*out = make([]BigqueryDateShardedSpecObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BigqueryTableSpec != nil {
		in, out := &in.BigqueryTableSpec, &out.BigqueryTableSpec
		*out = make([]BigqueryTableSpecObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GcsFilesetSpec != nil {
		in, out := &in.GcsFilesetSpec, &out.GcsFilesetSpec
		*out = make([]GcsFilesetSpecObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IntegratedSystem != nil {
		in, out := &in.IntegratedSystem, &out.IntegratedSystem
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryObservation.
func (in *EntryObservation) DeepCopy() *EntryObservation {
	if in == nil {
		return nil
	}
	out := new(EntryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryParameters) DeepCopyInto(out *EntryParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EntryGroup != nil {
		in, out := &in.EntryGroup, &out.EntryGroup
		*out = new(string)
		**out = **in
	}
	if in.EntryGroupRef != nil {
		in, out := &in.EntryGroupRef, &out.EntryGroupRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EntryGroupSelector != nil {
		in, out := &in.EntryGroupSelector, &out.EntryGroupSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.EntryID != nil {
		in, out := &in.EntryID, &out.EntryID
		*out = new(string)
		**out = **in
	}
	if in.GcsFilesetSpec != nil {
		in, out := &in.GcsFilesetSpec, &out.GcsFilesetSpec
		*out = make([]GcsFilesetSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LinkedResource != nil {
		in, out := &in.LinkedResource, &out.LinkedResource
		*out = new(string)
		**out = **in
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserSpecifiedSystem != nil {
		in, out := &in.UserSpecifiedSystem, &out.UserSpecifiedSystem
		*out = new(string)
		**out = **in
	}
	if in.UserSpecifiedType != nil {
		in, out := &in.UserSpecifiedType, &out.UserSpecifiedType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryParameters.
func (in *EntryParameters) DeepCopy() *EntryParameters {
	if in == nil {
		return nil
	}
	out := new(EntryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntrySpec) DeepCopyInto(out *EntrySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntrySpec.
func (in *EntrySpec) DeepCopy() *EntrySpec {
	if in == nil {
		return nil
	}
	out := new(EntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryStatus) DeepCopyInto(out *EntryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryStatus.
func (in *EntryStatus) DeepCopy() *EntryStatus {
	if in == nil {
		return nil
	}
	out := new(EntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcsFilesetSpecObservation) DeepCopyInto(out *GcsFilesetSpecObservation) {
	*out = *in
	if in.SampleGcsFileSpecs != nil {
		in, out := &in.SampleGcsFileSpecs, &out.SampleGcsFileSpecs
		*out = make([]SampleGcsFileSpecsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcsFilesetSpecObservation.
func (in *GcsFilesetSpecObservation) DeepCopy() *GcsFilesetSpecObservation {
	if in == nil {
		return nil
	}
	out := new(GcsFilesetSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcsFilesetSpecParameters) DeepCopyInto(out *GcsFilesetSpecParameters) {
	*out = *in
	if in.FilePatterns != nil {
		in, out := &in.FilePatterns, &out.FilePatterns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcsFilesetSpecParameters.
func (in *GcsFilesetSpecParameters) DeepCopy() *GcsFilesetSpecParameters {
	if in == nil {
		return nil
	}
	out := new(GcsFilesetSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleGcsFileSpecsObservation) DeepCopyInto(out *SampleGcsFileSpecsObservation) {
	*out = *in
	if in.FilePath != nil {
		in, out := &in.FilePath, &out.FilePath
		*out = new(string)
		**out = **in
	}
	if in.SizeBytes != nil {
		in, out := &in.SizeBytes, &out.SizeBytes
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleGcsFileSpecsObservation.
func (in *SampleGcsFileSpecsObservation) DeepCopy() *SampleGcsFileSpecsObservation {
	if in == nil {
		return nil
	}
	out := new(SampleGcsFileSpecsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SampleGcsFileSpecsParameters) DeepCopyInto(out *SampleGcsFileSpecsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SampleGcsFileSpecsParameters.
func (in *SampleGcsFileSpecsParameters) DeepCopy() *SampleGcsFileSpecsParameters {
	if in == nil {
		return nil
	}
	out := new(SampleGcsFileSpecsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecObservation) DeepCopyInto(out *TableSpecObservation) {
	*out = *in
	if in.GroupedEntry != nil {
		in, out := &in.GroupedEntry, &out.GroupedEntry
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecObservation.
func (in *TableSpecObservation) DeepCopy() *TableSpecObservation {
	if in == nil {
		return nil
	}
	out := new(TableSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableSpecParameters) DeepCopyInto(out *TableSpecParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableSpecParameters.
func (in *TableSpecParameters) DeepCopy() *TableSpecParameters {
	if in == nil {
		return nil
	}
	out := new(TableSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewSpecObservation) DeepCopyInto(out *ViewSpecObservation) {
	*out = *in
	if in.ViewQuery != nil {
		in, out := &in.ViewQuery, &out.ViewQuery
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewSpecObservation.
func (in *ViewSpecObservation) DeepCopy() *ViewSpecObservation {
	if in == nil {
		return nil
	}
	out := new(ViewSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ViewSpecParameters) DeepCopyInto(out *ViewSpecParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ViewSpecParameters.
func (in *ViewSpecParameters) DeepCopy() *ViewSpecParameters {
	if in == nil {
		return nil
	}
	out := new(ViewSpecParameters)
	in.DeepCopyInto(out)
	return out
}
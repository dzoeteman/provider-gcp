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
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this WorkloadIdentityPool
func (mg *WorkloadIdentityPool) GetTerraformResourceType() string {
	return "google_iam_workload_identity_pool"
}

// GetConnectionDetailsMapping for this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this WorkloadIdentityPool
func (tr *WorkloadIdentityPool) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this WorkloadIdentityPool using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *WorkloadIdentityPool) LateInitialize(attrs []byte) (bool, error) {
	params := &WorkloadIdentityPoolParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *WorkloadIdentityPool) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this WorkloadIdentityPoolProvider
func (mg *WorkloadIdentityPoolProvider) GetTerraformResourceType() string {
	return "google_iam_workload_identity_pool_provider"
}

// GetConnectionDetailsMapping for this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this WorkloadIdentityPoolProvider
func (tr *WorkloadIdentityPoolProvider) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this WorkloadIdentityPoolProvider using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *WorkloadIdentityPoolProvider) LateInitialize(attrs []byte) (bool, error) {
	params := &WorkloadIdentityPoolProviderParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *WorkloadIdentityPoolProvider) GetTerraformSchemaVersion() int {
	return 0
}

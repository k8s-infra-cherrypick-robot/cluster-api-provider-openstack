/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/api/v1beta1"
)

// ResolvedServerSpecApplyConfiguration represents an declarative configuration of the ResolvedServerSpec type for use
// with apply.
type ResolvedServerSpecApplyConfiguration struct {
	ServerGroupID *string                                      `json:"serverGroupID,omitempty"`
	ImageID       *string                                      `json:"imageID,omitempty"`
	FlavorID      *string                                      `json:"flavorID,omitempty"`
	Ports         []v1beta1.ResolvedPortSpecApplyConfiguration `json:"ports,omitempty"`
}

// ResolvedServerSpecApplyConfiguration constructs an declarative configuration of the ResolvedServerSpec type for use with
// apply.
func ResolvedServerSpec() *ResolvedServerSpecApplyConfiguration {
	return &ResolvedServerSpecApplyConfiguration{}
}

// WithServerGroupID sets the ServerGroupID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerGroupID field is set to the value of the last call.
func (b *ResolvedServerSpecApplyConfiguration) WithServerGroupID(value string) *ResolvedServerSpecApplyConfiguration {
	b.ServerGroupID = &value
	return b
}

// WithImageID sets the ImageID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageID field is set to the value of the last call.
func (b *ResolvedServerSpecApplyConfiguration) WithImageID(value string) *ResolvedServerSpecApplyConfiguration {
	b.ImageID = &value
	return b
}

// WithFlavorID sets the FlavorID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FlavorID field is set to the value of the last call.
func (b *ResolvedServerSpecApplyConfiguration) WithFlavorID(value string) *ResolvedServerSpecApplyConfiguration {
	b.FlavorID = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *ResolvedServerSpecApplyConfiguration) WithPorts(values ...*v1beta1.ResolvedPortSpecApplyConfiguration) *ResolvedServerSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

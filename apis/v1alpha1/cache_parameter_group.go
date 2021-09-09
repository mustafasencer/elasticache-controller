// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CacheParameterGroupSpec defines the desired state of CacheParameterGroup.
//
// Represents the output of a CreateCacheParameterGroup operation.
type CacheParameterGroupSpec struct {
	// The name of the cache parameter group family that the cache parameter group
	// can be used with.
	//
	// Valid values are: memcached1.4 | memcached1.5 | memcached1.6 | redis2.6 |
	// redis2.8 | redis3.2 | redis4.0 | redis5.0 | redis6.x |
	// +kubebuilder:validation:Required
	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily"`
	// A user-specified name for the cache parameter group.
	// +kubebuilder:validation:Required
	CacheParameterGroupName *string `json:"cacheParameterGroupName"`
	// A user-specified description for the cache parameter group.
	// +kubebuilder:validation:Required
	Description *string `json:"description"`
	// An array of parameter names and values for the parameter update. You must
	// supply at least one parameter name and value; subsequent arguments are optional.
	// A maximum of 20 parameters may be modified per request.
	ParameterNameValues []*ParameterNameValue `json:"parameterNameValues,omitempty"`
	// A list of tags to be added to this resource. A tag is a key-value pair. A
	// tag key must be accompanied by a tag value, although null is accepted.
	Tags []*Tag `json:"tags,omitempty"`
}

// CacheParameterGroupStatus defines the observed state of CacheParameterGroup
type CacheParameterGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// A list of events. Each element in the list contains detailed information
	// about one event.
	// +kubebuilder:validation:Optional
	Events []*Event `json:"events,omitempty"`
	// Indicates whether the parameter group is associated with a Global datastore
	// +kubebuilder:validation:Optional
	IsGlobal *bool `json:"isGlobal,omitempty"`
	// A list of Parameter instances.
	// +kubebuilder:validation:Optional
	Parameters []*Parameter `json:"parameters,omitempty"`
}

// CacheParameterGroup is the Schema for the CacheParameterGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type CacheParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CacheParameterGroupSpec   `json:"spec,omitempty"`
	Status            CacheParameterGroupStatus `json:"status,omitempty"`
}

// CacheParameterGroupList contains a list of CacheParameterGroup
// +kubebuilder:object:root=true
type CacheParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CacheParameterGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CacheParameterGroup{}, &CacheParameterGroupList{})
}

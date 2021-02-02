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

package v1alpha1

// AWSRegion represents an AWS regional identifier
type AWSRegion string

// AWSAccountID represents an AWS account identifier
type AWSAccountID string

// AWSResourceName represents an AWS Resource Name (ARN)
type AWSResourceName string

// AWSResourceReference represents a reference to an AWS Resource
// Either the combination of namepsace and name is used to identify a resource in the cluster
// or external field if resource is not managed in the cluster
type AWSResourceReference struct {
	// Namespace of the resource being referred to
	Namespace string `json:"namespace,omitempty"`
	// Name of the resource being referred to
	Name string `json:"name,omitempty"`

	// external is an alternative field to specify resources which are not
	// managed by ACK controller
	External string `json:"external,omitempty"`
}

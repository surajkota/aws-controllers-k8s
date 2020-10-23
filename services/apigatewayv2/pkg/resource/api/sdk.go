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

package api

import (
	"context"
	corev1 "k8s.io/api/core/v1"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.API{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetApiWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetApi", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiEndpoint != nil {
		ko.Status.APIEndpoint = resp.ApiEndpoint
	}
	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.ApiId != nil {
		ko.Status.APIID = resp.ApiId
	}
	if resp.ApiKeySelectionExpression != nil {
		ko.Spec.APIKeySelectionExpression = resp.ApiKeySelectionExpression
	}
	if resp.CorsConfiguration != nil {
		f4 := &svcapitypes.Cors{}
		if resp.CorsConfiguration.AllowCredentials != nil {
			f4.AllowCredentials = resp.CorsConfiguration.AllowCredentials
		}
		if resp.CorsConfiguration.AllowHeaders != nil {
			f4f1 := []*string{}
			for _, f4f1iter := range resp.CorsConfiguration.AllowHeaders {
				var f4f1elem string
				f4f1elem = *f4f1iter
				f4f1 = append(f4f1, &f4f1elem)
			}
			f4.AllowHeaders = f4f1
		}
		if resp.CorsConfiguration.AllowMethods != nil {
			f4f2 := []*string{}
			for _, f4f2iter := range resp.CorsConfiguration.AllowMethods {
				var f4f2elem string
				f4f2elem = *f4f2iter
				f4f2 = append(f4f2, &f4f2elem)
			}
			f4.AllowMethods = f4f2
		}
		if resp.CorsConfiguration.AllowOrigins != nil {
			f4f3 := []*string{}
			for _, f4f3iter := range resp.CorsConfiguration.AllowOrigins {
				var f4f3elem string
				f4f3elem = *f4f3iter
				f4f3 = append(f4f3, &f4f3elem)
			}
			f4.AllowOrigins = f4f3
		}
		if resp.CorsConfiguration.ExposeHeaders != nil {
			f4f4 := []*string{}
			for _, f4f4iter := range resp.CorsConfiguration.ExposeHeaders {
				var f4f4elem string
				f4f4elem = *f4f4iter
				f4f4 = append(f4f4, &f4f4elem)
			}
			f4.ExposeHeaders = f4f4
		}
		if resp.CorsConfiguration.MaxAge != nil {
			f4.MaxAge = resp.CorsConfiguration.MaxAge
		}
		ko.Spec.CorsConfiguration = f4
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	}
	if resp.DisableExecuteApiEndpoint != nil {
		ko.Spec.DisableExecuteAPIEndpoint = resp.DisableExecuteApiEndpoint
	}
	if resp.DisableSchemaValidation != nil {
		ko.Spec.DisableSchemaValidation = resp.DisableSchemaValidation
	}
	if resp.ImportInfo != nil {
		f9 := []*string{}
		for _, f9iter := range resp.ImportInfo {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		ko.Status.ImportInfo = f9
	}
	if resp.Name != nil {
		ko.Spec.Name = resp.Name
	}
	if resp.ProtocolType != nil {
		ko.Spec.ProtocolType = resp.ProtocolType
	}
	if resp.RouteSelectionExpression != nil {
		ko.Spec.RouteSelectionExpression = resp.RouteSelectionExpression
	}
	if resp.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range resp.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		ko.Spec.Tags = f13
	}
	if resp.Version != nil {
		ko.Spec.Version = resp.Version
	}
	if resp.Warnings != nil {
		f15 := []*string{}
		for _, f15iter := range resp.Warnings {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		ko.Status.Warnings = f15
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.APIID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetApiInput, error) {
	res := &svcsdk.GetApiInput{}

	if r.ko.Status.APIID != nil {
		res.SetApiId(*r.ko.Status.APIID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateApiWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateApi", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiEndpoint != nil {
		ko.Status.APIEndpoint = resp.ApiEndpoint
	}
	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.ApiId != nil {
		ko.Status.APIID = resp.ApiId
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.ImportInfo != nil {
		f9 := []*string{}
		for _, f9iter := range resp.ImportInfo {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		ko.Status.ImportInfo = f9
	}
	if resp.Warnings != nil {
		f15 := []*string{}
		for _, f15iter := range resp.Warnings {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		ko.Status.Warnings = f15
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateApiInput, error) {
	res := &svcsdk.CreateApiInput{}

	if r.ko.Spec.APIKeySelectionExpression != nil {
		res.SetApiKeySelectionExpression(*r.ko.Spec.APIKeySelectionExpression)
	}
	if r.ko.Spec.CorsConfiguration != nil {
		f1 := &svcsdk.Cors{}
		if r.ko.Spec.CorsConfiguration.AllowCredentials != nil {
			f1.SetAllowCredentials(*r.ko.Spec.CorsConfiguration.AllowCredentials)
		}
		if r.ko.Spec.CorsConfiguration.AllowHeaders != nil {
			f1f1 := []*string{}
			for _, f1f1iter := range r.ko.Spec.CorsConfiguration.AllowHeaders {
				var f1f1elem string
				f1f1elem = *f1f1iter
				f1f1 = append(f1f1, &f1f1elem)
			}
			f1.SetAllowHeaders(f1f1)
		}
		if r.ko.Spec.CorsConfiguration.AllowMethods != nil {
			f1f2 := []*string{}
			for _, f1f2iter := range r.ko.Spec.CorsConfiguration.AllowMethods {
				var f1f2elem string
				f1f2elem = *f1f2iter
				f1f2 = append(f1f2, &f1f2elem)
			}
			f1.SetAllowMethods(f1f2)
		}
		if r.ko.Spec.CorsConfiguration.AllowOrigins != nil {
			f1f3 := []*string{}
			for _, f1f3iter := range r.ko.Spec.CorsConfiguration.AllowOrigins {
				var f1f3elem string
				f1f3elem = *f1f3iter
				f1f3 = append(f1f3, &f1f3elem)
			}
			f1.SetAllowOrigins(f1f3)
		}
		if r.ko.Spec.CorsConfiguration.ExposeHeaders != nil {
			f1f4 := []*string{}
			for _, f1f4iter := range r.ko.Spec.CorsConfiguration.ExposeHeaders {
				var f1f4elem string
				f1f4elem = *f1f4iter
				f1f4 = append(f1f4, &f1f4elem)
			}
			f1.SetExposeHeaders(f1f4)
		}
		if r.ko.Spec.CorsConfiguration.MaxAge != nil {
			f1.SetMaxAge(*r.ko.Spec.CorsConfiguration.MaxAge)
		}
		res.SetCorsConfiguration(f1)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.DisableExecuteAPIEndpoint != nil {
		res.SetDisableExecuteApiEndpoint(*r.ko.Spec.DisableExecuteAPIEndpoint)
	}
	if r.ko.Spec.DisableSchemaValidation != nil {
		res.SetDisableSchemaValidation(*r.ko.Spec.DisableSchemaValidation)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.ProtocolType != nil {
		res.SetProtocolType(*r.ko.Spec.ProtocolType)
	}
	if r.ko.Spec.RouteKey != nil {
		res.SetRouteKey(*r.ko.Spec.RouteKey)
	}
	if r.ko.Spec.RouteSelectionExpression != nil {
		res.SetRouteSelectionExpression(*r.ko.Spec.RouteSelectionExpression)
	}
	if r.ko.Spec.Tags != nil {
		f10 := map[string]*string{}
		for f10key, f10valiter := range r.ko.Spec.Tags {
			var f10val string
			f10val = *f10valiter
			f10[f10key] = &f10val
		}
		res.SetTags(f10)
	}
	if r.ko.Spec.Target != nil {
		res.SetTarget(*r.ko.Spec.Target)
	}
	if r.ko.Spec.Version != nil {
		res.SetVersion(*r.ko.Spec.Version)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateApiWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateApi", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiEndpoint != nil {
		ko.Status.APIEndpoint = resp.ApiEndpoint
	}
	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.ApiId != nil {
		ko.Status.APIID = resp.ApiId
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.ImportInfo != nil {
		f9 := []*string{}
		for _, f9iter := range resp.ImportInfo {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		ko.Status.ImportInfo = f9
	}
	if resp.Warnings != nil {
		f15 := []*string{}
		for _, f15iter := range resp.Warnings {
			var f15elem string
			f15elem = *f15iter
			f15 = append(f15, &f15elem)
		}
		ko.Status.Warnings = f15
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateApiInput, error) {
	res := &svcsdk.UpdateApiInput{}

	if r.ko.Status.APIID != nil {
		res.SetApiId(*r.ko.Status.APIID)
	}
	if r.ko.Spec.APIKeySelectionExpression != nil {
		res.SetApiKeySelectionExpression(*r.ko.Spec.APIKeySelectionExpression)
	}
	if r.ko.Spec.CorsConfiguration != nil {
		f2 := &svcsdk.Cors{}
		if r.ko.Spec.CorsConfiguration.AllowCredentials != nil {
			f2.SetAllowCredentials(*r.ko.Spec.CorsConfiguration.AllowCredentials)
		}
		if r.ko.Spec.CorsConfiguration.AllowHeaders != nil {
			f2f1 := []*string{}
			for _, f2f1iter := range r.ko.Spec.CorsConfiguration.AllowHeaders {
				var f2f1elem string
				f2f1elem = *f2f1iter
				f2f1 = append(f2f1, &f2f1elem)
			}
			f2.SetAllowHeaders(f2f1)
		}
		if r.ko.Spec.CorsConfiguration.AllowMethods != nil {
			f2f2 := []*string{}
			for _, f2f2iter := range r.ko.Spec.CorsConfiguration.AllowMethods {
				var f2f2elem string
				f2f2elem = *f2f2iter
				f2f2 = append(f2f2, &f2f2elem)
			}
			f2.SetAllowMethods(f2f2)
		}
		if r.ko.Spec.CorsConfiguration.AllowOrigins != nil {
			f2f3 := []*string{}
			for _, f2f3iter := range r.ko.Spec.CorsConfiguration.AllowOrigins {
				var f2f3elem string
				f2f3elem = *f2f3iter
				f2f3 = append(f2f3, &f2f3elem)
			}
			f2.SetAllowOrigins(f2f3)
		}
		if r.ko.Spec.CorsConfiguration.ExposeHeaders != nil {
			f2f4 := []*string{}
			for _, f2f4iter := range r.ko.Spec.CorsConfiguration.ExposeHeaders {
				var f2f4elem string
				f2f4elem = *f2f4iter
				f2f4 = append(f2f4, &f2f4elem)
			}
			f2.SetExposeHeaders(f2f4)
		}
		if r.ko.Spec.CorsConfiguration.MaxAge != nil {
			f2.SetMaxAge(*r.ko.Spec.CorsConfiguration.MaxAge)
		}
		res.SetCorsConfiguration(f2)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.DisableExecuteAPIEndpoint != nil {
		res.SetDisableExecuteApiEndpoint(*r.ko.Spec.DisableExecuteAPIEndpoint)
	}
	if r.ko.Spec.DisableSchemaValidation != nil {
		res.SetDisableSchemaValidation(*r.ko.Spec.DisableSchemaValidation)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.RouteKey != nil {
		res.SetRouteKey(*r.ko.Spec.RouteKey)
	}
	if r.ko.Spec.RouteSelectionExpression != nil {
		res.SetRouteSelectionExpression(*r.ko.Spec.RouteSelectionExpression)
	}
	if r.ko.Spec.Target != nil {
		res.SetTarget(*r.ko.Spec.Target)
	}
	if r.ko.Spec.Version != nil {
		res.SetVersion(*r.ko.Spec.Version)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteApiWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteApi", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteApiInput, error) {
	res := &svcsdk.DeleteApiInput{}

	if r.ko.Status.APIID != nil {
		res.SetApiId(*r.ko.Status.APIID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.API,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}

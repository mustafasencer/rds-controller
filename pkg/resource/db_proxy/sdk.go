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

package db_proxy

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/rds-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.RDS{}
	_ = &svcapitypes.DBProxy{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeDBProxiesOutput
	resp, err = rm.sdkapi.DescribeDBProxiesWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeDBProxies", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "DBProxyNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.DBProxies {
		if elem.Auth != nil {
			f0 := []*svcapitypes.UserAuthConfig{}
			for _, f0iter := range elem.Auth {
				f0elem := &svcapitypes.UserAuthConfig{}
				if f0iter.AuthScheme != nil {
					f0elem.AuthScheme = f0iter.AuthScheme
				}
				if f0iter.Description != nil {
					f0elem.Description = f0iter.Description
				}
				if f0iter.IAMAuth != nil {
					f0elem.IAMAuth = f0iter.IAMAuth
				}
				if f0iter.SecretArn != nil {
					f0elem.SecretARN = f0iter.SecretArn
				}
				if f0iter.UserName != nil {
					f0elem.UserName = f0iter.UserName
				}
				f0 = append(f0, f0elem)
			}
			ko.Spec.Auth = f0
		} else {
			ko.Spec.Auth = nil
		}
		if elem.CreatedDate != nil {
			ko.Status.CreatedDate = &metav1.Time{*elem.CreatedDate}
		} else {
			ko.Status.CreatedDate = nil
		}
		if elem.DBProxyArn != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.DBProxyArn)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.DBProxyName != nil {
			ko.Spec.Name = elem.DBProxyName
		} else {
			ko.Spec.Name = nil
		}
		if elem.DebugLogging != nil {
			ko.Spec.DebugLogging = elem.DebugLogging
		} else {
			ko.Spec.DebugLogging = nil
		}
		if elem.Endpoint != nil {
			ko.Status.Endpoint = elem.Endpoint
		} else {
			ko.Status.Endpoint = nil
		}
		if elem.EngineFamily != nil {
			ko.Spec.EngineFamily = elem.EngineFamily
		} else {
			ko.Spec.EngineFamily = nil
		}
		if elem.IdleClientTimeout != nil {
			ko.Spec.IdleClientTimeout = elem.IdleClientTimeout
		} else {
			ko.Spec.IdleClientTimeout = nil
		}
		if elem.RequireTLS != nil {
			ko.Spec.RequireTLS = elem.RequireTLS
		} else {
			ko.Spec.RequireTLS = nil
		}
		if elem.RoleArn != nil {
			ko.Spec.RoleARN = elem.RoleArn
		} else {
			ko.Spec.RoleARN = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		if elem.UpdatedDate != nil {
			ko.Status.UpdatedDate = &metav1.Time{*elem.UpdatedDate}
		} else {
			ko.Status.UpdatedDate = nil
		}
		if elem.VpcId != nil {
			ko.Status.VPCID = elem.VpcId
		} else {
			ko.Status.VPCID = nil
		}
		if elem.VpcSecurityGroupIds != nil {
			f13 := []*string{}
			for _, f13iter := range elem.VpcSecurityGroupIds {
				var f13elem string
				f13elem = *f13iter
				f13 = append(f13, &f13elem)
			}
			ko.Spec.VPCSecurityGroupIDs = f13
		} else {
			ko.Spec.VPCSecurityGroupIDs = nil
		}
		if elem.VpcSubnetIds != nil {
			f14 := []*string{}
			for _, f14iter := range elem.VpcSubnetIds {
				var f14elem string
				f14elem = *f14iter
				f14 = append(f14, &f14elem)
			}
			ko.Spec.VPCSubnetIDs = f14
		} else {
			ko.Spec.VPCSubnetIDs = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Spec.Name == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeDBProxiesInput, error) {
	res := &svcsdk.DescribeDBProxiesInput{}

	if r.ko.Spec.Name != nil {
		res.SetDBProxyName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateDBProxyOutput
	_ = resp
	resp, err = rm.sdkapi.CreateDBProxyWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateDBProxy", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.DBProxy.Auth != nil {
		f0 := []*svcapitypes.UserAuthConfig{}
		for _, f0iter := range resp.DBProxy.Auth {
			f0elem := &svcapitypes.UserAuthConfig{}
			if f0iter.AuthScheme != nil {
				f0elem.AuthScheme = f0iter.AuthScheme
			}
			if f0iter.Description != nil {
				f0elem.Description = f0iter.Description
			}
			if f0iter.IAMAuth != nil {
				f0elem.IAMAuth = f0iter.IAMAuth
			}
			if f0iter.SecretArn != nil {
				f0elem.SecretARN = f0iter.SecretArn
			}
			if f0iter.UserName != nil {
				f0elem.UserName = f0iter.UserName
			}
			f0 = append(f0, f0elem)
		}
		ko.Spec.Auth = f0
	} else {
		ko.Spec.Auth = nil
	}
	if resp.DBProxy.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.DBProxy.CreatedDate}
	} else {
		ko.Status.CreatedDate = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DBProxy.DBProxyArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DBProxy.DBProxyArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DBProxy.DBProxyName != nil {
		ko.Spec.Name = resp.DBProxy.DBProxyName
	} else {
		ko.Spec.Name = nil
	}
	if resp.DBProxy.DebugLogging != nil {
		ko.Spec.DebugLogging = resp.DBProxy.DebugLogging
	} else {
		ko.Spec.DebugLogging = nil
	}
	if resp.DBProxy.Endpoint != nil {
		ko.Status.Endpoint = resp.DBProxy.Endpoint
	} else {
		ko.Status.Endpoint = nil
	}
	if resp.DBProxy.EngineFamily != nil {
		ko.Spec.EngineFamily = resp.DBProxy.EngineFamily
	} else {
		ko.Spec.EngineFamily = nil
	}
	if resp.DBProxy.IdleClientTimeout != nil {
		ko.Spec.IdleClientTimeout = resp.DBProxy.IdleClientTimeout
	} else {
		ko.Spec.IdleClientTimeout = nil
	}
	if resp.DBProxy.RequireTLS != nil {
		ko.Spec.RequireTLS = resp.DBProxy.RequireTLS
	} else {
		ko.Spec.RequireTLS = nil
	}
	if resp.DBProxy.RoleArn != nil {
		ko.Spec.RoleARN = resp.DBProxy.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}
	if resp.DBProxy.Status != nil {
		ko.Status.Status = resp.DBProxy.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.DBProxy.UpdatedDate != nil {
		ko.Status.UpdatedDate = &metav1.Time{*resp.DBProxy.UpdatedDate}
	} else {
		ko.Status.UpdatedDate = nil
	}
	if resp.DBProxy.VpcId != nil {
		ko.Status.VPCID = resp.DBProxy.VpcId
	} else {
		ko.Status.VPCID = nil
	}
	if resp.DBProxy.VpcSecurityGroupIds != nil {
		f13 := []*string{}
		for _, f13iter := range resp.DBProxy.VpcSecurityGroupIds {
			var f13elem string
			f13elem = *f13iter
			f13 = append(f13, &f13elem)
		}
		ko.Spec.VPCSecurityGroupIDs = f13
	} else {
		ko.Spec.VPCSecurityGroupIDs = nil
	}
	if resp.DBProxy.VpcSubnetIds != nil {
		f14 := []*string{}
		for _, f14iter := range resp.DBProxy.VpcSubnetIds {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		ko.Spec.VPCSubnetIDs = f14
	} else {
		ko.Spec.VPCSubnetIDs = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateDBProxyInput, error) {
	res := &svcsdk.CreateDBProxyInput{}

	if r.ko.Spec.Auth != nil {
		f0 := []*svcsdk.UserAuthConfig{}
		for _, f0iter := range r.ko.Spec.Auth {
			f0elem := &svcsdk.UserAuthConfig{}
			if f0iter.AuthScheme != nil {
				f0elem.SetAuthScheme(*f0iter.AuthScheme)
			}
			if f0iter.Description != nil {
				f0elem.SetDescription(*f0iter.Description)
			}
			if f0iter.IAMAuth != nil {
				f0elem.SetIAMAuth(*f0iter.IAMAuth)
			}
			if f0iter.SecretARN != nil {
				f0elem.SetSecretArn(*f0iter.SecretARN)
			}
			if f0iter.UserName != nil {
				f0elem.SetUserName(*f0iter.UserName)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAuth(f0)
	}
	if r.ko.Spec.Name != nil {
		res.SetDBProxyName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.DebugLogging != nil {
		res.SetDebugLogging(*r.ko.Spec.DebugLogging)
	}
	if r.ko.Spec.EngineFamily != nil {
		res.SetEngineFamily(*r.ko.Spec.EngineFamily)
	}
	if r.ko.Spec.IdleClientTimeout != nil {
		res.SetIdleClientTimeout(*r.ko.Spec.IdleClientTimeout)
	}
	if r.ko.Spec.RequireTLS != nil {
		res.SetRequireTLS(*r.ko.Spec.RequireTLS)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.Tags != nil {
		f7 := []*svcsdk.Tag{}
		for _, f7iter := range r.ko.Spec.Tags {
			f7elem := &svcsdk.Tag{}
			if f7iter.Key != nil {
				f7elem.SetKey(*f7iter.Key)
			}
			if f7iter.Value != nil {
				f7elem.SetValue(*f7iter.Value)
			}
			f7 = append(f7, f7elem)
		}
		res.SetTags(f7)
	}
	if r.ko.Spec.VPCSecurityGroupIDs != nil {
		f8 := []*string{}
		for _, f8iter := range r.ko.Spec.VPCSecurityGroupIDs {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		res.SetVpcSecurityGroupIds(f8)
	}
	if r.ko.Spec.VPCSubnetIDs != nil {
		f9 := []*string{}
		for _, f9iter := range r.ko.Spec.VPCSubnetIDs {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		res.SetVpcSubnetIds(f9)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.ModifyDBProxyOutput
	_ = resp
	resp, err = rm.sdkapi.ModifyDBProxyWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyDBProxy", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.DBProxy.Auth != nil {
		f0 := []*svcapitypes.UserAuthConfig{}
		for _, f0iter := range resp.DBProxy.Auth {
			f0elem := &svcapitypes.UserAuthConfig{}
			if f0iter.AuthScheme != nil {
				f0elem.AuthScheme = f0iter.AuthScheme
			}
			if f0iter.Description != nil {
				f0elem.Description = f0iter.Description
			}
			if f0iter.IAMAuth != nil {
				f0elem.IAMAuth = f0iter.IAMAuth
			}
			if f0iter.SecretArn != nil {
				f0elem.SecretARN = f0iter.SecretArn
			}
			if f0iter.UserName != nil {
				f0elem.UserName = f0iter.UserName
			}
			f0 = append(f0, f0elem)
		}
		ko.Spec.Auth = f0
	} else {
		ko.Spec.Auth = nil
	}
	if resp.DBProxy.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.DBProxy.CreatedDate}
	} else {
		ko.Status.CreatedDate = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DBProxy.DBProxyArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DBProxy.DBProxyArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DBProxy.DBProxyName != nil {
		ko.Spec.Name = resp.DBProxy.DBProxyName
	} else {
		ko.Spec.Name = nil
	}
	if resp.DBProxy.DebugLogging != nil {
		ko.Spec.DebugLogging = resp.DBProxy.DebugLogging
	} else {
		ko.Spec.DebugLogging = nil
	}
	if resp.DBProxy.Endpoint != nil {
		ko.Status.Endpoint = resp.DBProxy.Endpoint
	} else {
		ko.Status.Endpoint = nil
	}
	if resp.DBProxy.EngineFamily != nil {
		ko.Spec.EngineFamily = resp.DBProxy.EngineFamily
	} else {
		ko.Spec.EngineFamily = nil
	}
	if resp.DBProxy.IdleClientTimeout != nil {
		ko.Spec.IdleClientTimeout = resp.DBProxy.IdleClientTimeout
	} else {
		ko.Spec.IdleClientTimeout = nil
	}
	if resp.DBProxy.RequireTLS != nil {
		ko.Spec.RequireTLS = resp.DBProxy.RequireTLS
	} else {
		ko.Spec.RequireTLS = nil
	}
	if resp.DBProxy.RoleArn != nil {
		ko.Spec.RoleARN = resp.DBProxy.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}
	if resp.DBProxy.Status != nil {
		ko.Status.Status = resp.DBProxy.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.DBProxy.UpdatedDate != nil {
		ko.Status.UpdatedDate = &metav1.Time{*resp.DBProxy.UpdatedDate}
	} else {
		ko.Status.UpdatedDate = nil
	}
	if resp.DBProxy.VpcId != nil {
		ko.Status.VPCID = resp.DBProxy.VpcId
	} else {
		ko.Status.VPCID = nil
	}
	if resp.DBProxy.VpcSecurityGroupIds != nil {
		f13 := []*string{}
		for _, f13iter := range resp.DBProxy.VpcSecurityGroupIds {
			var f13elem string
			f13elem = *f13iter
			f13 = append(f13, &f13elem)
		}
		ko.Spec.VPCSecurityGroupIDs = f13
	} else {
		ko.Spec.VPCSecurityGroupIDs = nil
	}
	if resp.DBProxy.VpcSubnetIds != nil {
		f14 := []*string{}
		for _, f14iter := range resp.DBProxy.VpcSubnetIds {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		ko.Spec.VPCSubnetIDs = f14
	} else {
		ko.Spec.VPCSubnetIDs = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.ModifyDBProxyInput, error) {
	res := &svcsdk.ModifyDBProxyInput{}

	if r.ko.Spec.Auth != nil {
		f0 := []*svcsdk.UserAuthConfig{}
		for _, f0iter := range r.ko.Spec.Auth {
			f0elem := &svcsdk.UserAuthConfig{}
			if f0iter.AuthScheme != nil {
				f0elem.SetAuthScheme(*f0iter.AuthScheme)
			}
			if f0iter.Description != nil {
				f0elem.SetDescription(*f0iter.Description)
			}
			if f0iter.IAMAuth != nil {
				f0elem.SetIAMAuth(*f0iter.IAMAuth)
			}
			if f0iter.SecretARN != nil {
				f0elem.SetSecretArn(*f0iter.SecretARN)
			}
			if f0iter.UserName != nil {
				f0elem.SetUserName(*f0iter.UserName)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAuth(f0)
	}
	if r.ko.Spec.Name != nil {
		res.SetDBProxyName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.DebugLogging != nil {
		res.SetDebugLogging(*r.ko.Spec.DebugLogging)
	}
	if r.ko.Spec.IdleClientTimeout != nil {
		res.SetIdleClientTimeout(*r.ko.Spec.IdleClientTimeout)
	}
	if r.ko.Spec.RequireTLS != nil {
		res.SetRequireTLS(*r.ko.Spec.RequireTLS)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteDBProxyOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteDBProxyWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteDBProxy", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDBProxyInput, error) {
	res := &svcsdk.DeleteDBProxyInput{}

	if r.ko.Spec.Name != nil {
		res.SetDBProxyName(*r.ko.Spec.Name)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.DBProxy,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
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
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidSubnet":
		return true
	default:
		return false
	}
}

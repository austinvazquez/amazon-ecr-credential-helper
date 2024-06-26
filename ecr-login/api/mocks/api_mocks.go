// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/awslabs/amazon-ecr-credential-helper/ecr-login/api (interfaces: ECRAPI)

package mock_api

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
)

type MockECRAPI struct {
	GetAuthorizationTokenFn func(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)
}

func (m MockECRAPI) GetAuthorizationToken(_ context.Context, input *ecr.GetAuthorizationTokenInput, _ ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	return m.GetAuthorizationTokenFn(input)
}

type MockECRPublicAPI struct {
	GetAuthorizationTokenFn func(input *ecrpublic.GetAuthorizationTokenInput) (*ecrpublic.GetAuthorizationTokenOutput, error)
}

func (m MockECRPublicAPI) GetAuthorizationToken(_ context.Context, input *ecrpublic.GetAuthorizationTokenInput, _ ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	return m.GetAuthorizationTokenFn(input)
}

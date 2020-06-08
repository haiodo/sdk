// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clients

import (
	"context"

	"google.golang.org/grpc"

	"github.com/networkservicemesh/sdk/pkg/registry/core/next"

	"github.com/networkservicemesh/api/pkg/api/registry"
)

type discoveryClientToServer struct {
	client registry.NetworkServiceDiscoveryClient
}

// NewNextDiscoveryClient - returns a registry.NetworkServiceDiscoveryClient wrapped around the supplied client, and allow to call next on passed context
func NewNextDiscoveryClient(client registry.NetworkServiceDiscoveryClient) registry.NetworkServiceDiscoveryClient {
	return &discoveryClientToServer{client: client}
}

func (c *discoveryClientToServer) FindNetworkService(ctx context.Context, request *registry.FindNetworkServiceRequest, opts ...grpc.CallOption) (*registry.FindNetworkServiceResponse, error) {
	result, err := c.client.FindNetworkService(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	var nextResult *registry.FindNetworkServiceResponse
	nextResult, err = next.DiscoveryClient(ctx).FindNetworkService(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	result.NetworkServiceEndpoints = append(result.NetworkServiceEndpoints, nextResult.NetworkServiceEndpoints...)
	for k, v := range nextResult.NetworkServiceManagers {
		result.NetworkServiceManagers[k] = v
	}
	return result, nil
}

// Implementation check
var _ registry.NetworkServiceDiscoveryClient = &discoveryClientToServer{}

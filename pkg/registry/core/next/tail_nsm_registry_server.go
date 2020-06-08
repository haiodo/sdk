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

package next

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/api/pkg/api/registry"
	"google.golang.org/grpc"
)

type tailRegistryNSMClient struct {
}

func (t *tailRegistryNSMClient) RegisterNSM(_ context.Context, in *registry.NetworkServiceManager, _ ...grpc.CallOption) (*registry.NetworkServiceManager, error) {
	return in, nil
}

func (t *tailRegistryNSMClient) GetEndpoints(_ context.Context, _ *empty.Empty, _ ...grpc.CallOption) (*registry.NetworkServiceEndpointList, error) {
	//  return empty
	return &registry.NetworkServiceEndpointList{}, nil
}

var _ registry.NsmRegistryClient = &tailRegistryNSMClient{}

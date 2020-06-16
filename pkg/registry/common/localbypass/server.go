// Copyright (c) 2020 Cisco Systems, Inc.
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

// Package localbypass provides NetworkServiceRegistryServer that registers local Endpoints
// and adds them to localbypass.SocketMap
package localbypass

import (
	"context"
	"errors"

	"github.com/networkservicemesh/sdk/pkg/registry/common/seturl"
	"github.com/networkservicemesh/sdk/pkg/registry/core/next"
	"github.com/networkservicemesh/sdk/pkg/tools/localbypass"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/api/pkg/api/registry"
)

type localBypassRegistry struct {
	sockets localbypass.SocketMap
}

func (l localBypassRegistry) Register(ctx context.Context, request *registry.NetworkServiceEndpoint) (*registry.NetworkServiceEndpoint, error) {
	endpointURL := seturl.EndpointURL(ctx)
	if endpointURL == nil {
		return nil, errors.New("invalid endpoint URL passed with context")
	}
	l.sockets.LoadOrStore(request.Name, endpointURL)
	return next.NetworkServiceEndpointRegistryServer(ctx).Register(ctx, request)
}

func (l localBypassRegistry) Find(query *registry.NetworkServiceEndpointQuery, s registry.NetworkServiceEndpointRegistry_FindServer) error {
	return next.NetworkServiceEndpointRegistryServer(s.Context()).Find(query, s)
}

func (l localBypassRegistry) Unregister(ctx context.Context, request *registry.NetworkServiceEndpoint) (*empty.Empty, error) {
	l.sockets.Delete(request.Name)
	return next.NetworkServiceEndpointRegistryServer(ctx).Unregister(ctx, request)
}

// NewNetworkServiceRegistryServer - creates a NetworkServiceRegistryServer that registers local Endpoints
//				and adds them to localbypass.SocketMap
//             - sockets - map of networkServiceEndpoint names to their unix socket addresses
func NewNetworkServiceRegistryServer(sockets localbypass.SocketMap) registry.NetworkServiceEndpointRegistryServer {
	return &localBypassRegistry{sockets: sockets}
}

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

package setid_test

import (
	"context"
	"testing"

	"github.com/networkservicemesh/api/pkg/api/registry"
	"github.com/stretchr/testify/require"

	"github.com/networkservicemesh/sdk/pkg/registry/common/setid"
	"github.com/networkservicemesh/sdk/pkg/registry/core/next"
	"github.com/networkservicemesh/sdk/pkg/registry/memory"
)

func TestSetIDNetworkServiceRegistryServer_RegisterNSE(t *testing.T) {
	storage := &memory.Storage{}
	storage.NetworkServiceManagers.Store("nsm-1", &registry.NetworkServiceManager{Name: "nsm-1"})
	chain := next.NewNetworkServiceRegistryServer(setid.NewNetworkServiceRegistryServer(), memory.NewNetworkServiceRegistryServer(storage))
	nse := &registry.NetworkServiceEndpoint{
		NetworkServiceName: "ns-1",
		Payload:            "IP",
	}
	registration := &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    "ns-1",
			Payload: "IP",
		},
		NetworkServiceEndpoint: nse,
	}
	resp, err := chain.RegisterNSE(context.Background(), registration)
	require.Nil(t, err)
	require.NotEmpty(t, resp.NetworkServiceEndpoint.Name)
}

func TestSetIDNetworkServiceRegistryServer_BulkRegisterNSE(t *testing.T) {
	s := next.NewNetworkServiceRegistryServer(setid.NewNetworkServiceRegistryServer(), &assertServer{T: t})
	require.Nil(t, s.BulkRegisterNSE(&emptyBulkRegisterNSEServer{}))
}

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

package updatepath

import (
	"github.com/google/uuid"
	"github.com/networkservicemesh/api/pkg/api/networkservice"
	"github.com/pkg/errors"
)

/*
	Logic for Update path:

	1. If current path segment.Name is equal to segmentName passed, it will just update current connection.Id and exit.
	2. If current path segment.Name is not equal to segmentName:
		2.1 if path has next segment available and next name is segmentName, take Id from next path segment
		2.2 if no next path segment available, it will add one more path segment and generate new Id, update connection.Id
		2.3 if path has next segment available, but next name is not equal to segmentName, will return error.
	3. if Index == 0, and there is no current segment will add one.
*/
func updatePath(conn *networkservice.Connection, segmentName string) (*networkservice.Connection, error) {
	if conn == nil {
		conn = &networkservice.Connection{}
	}

	// If we don't have a Path, add one
	if conn.GetPath() == nil || len(conn.GetPath().GetPathSegments()) == 0 {
		if conn.Path == nil {
			conn.Path = &networkservice.Path{}
		}
		conn.Path.Index = 0
		// Replace path for first segment.
		conn.Path.PathSegments = []*networkservice.PathSegment{
			{
				Name: segmentName,
				Id:   conn.Id,
			},
		}
		return conn, nil
	}

	path := conn.GetPath()

	if int(path.Index) < len(path.PathSegments) && path.PathSegments[path.Index].Name == segmentName {
		// We already in current segment, just update connection Id, no need to increment
		conn.Id = path.PathSegments[path.Index].Id
		return conn, nil
	}

	// We need to move to next item
	nextIndex := int(path.Index) + 1

	if nextIndex > len(path.GetPathSegments()) {
		// We have index > segments count
		return nil, errors.Errorf("NetworkServiceRequest.Connection.Path.Index+1==%d should be less or equal len(NetworkServiceRequest.Connection.Path.PathSegement)==%d",
			nextIndex, len(path.GetPathSegments()))
	}
	if nextIndex < len(path.GetPathSegments()) && path.GetPathSegments()[nextIndex].Name != segmentName {
		// We have next, but name is different
		return nil, errors.Errorf("NetworkServiceRequest.Connection.Path.PathSegement[%d].Name should be %s, but it is %s",
			nextIndex, segmentName, path.GetPathSegments()[nextIndex].Name)
	}

	if nextIndex >= len(path.GetPathSegments()) {
		// Generate new Id for Path segment and connection Id.
		conn.Id = uuid.New().String()
		path.PathSegments = append(path.PathSegments, &networkservice.PathSegment{
			Name: segmentName,
			Id:   conn.Id,
		})
	} else {
		// Just update ID from next segment.
		conn.Id = path.GetPathSegments()[nextIndex].Id
	}
	// Increment index
	conn.Path.Index++

	return conn, nil
}

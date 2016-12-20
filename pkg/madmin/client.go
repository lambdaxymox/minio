/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package madmin

const (
	minioAdminOpHeader = "X-Minio-Operation"
)

// AdminClient - interface to Minio Management API
type AdminClient struct {
	client *Client
}

// NewAdminClient - create new Management client
func NewAdminClient(addr string, access string, secret string, secure bool) (*AdminClient, error) {
	client, err := New(addr, access, secret, secure)
	if err != nil {
		return nil, err
	}

	return &AdminClient{client: client}, nil
}

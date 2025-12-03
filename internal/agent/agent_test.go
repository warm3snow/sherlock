// Copyright 2024 Sherlock Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agent

import (
	"testing"
)

func TestParseConnectionDirect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantHost string
		wantPort int
		wantUser string
		wantNil  bool
	}{
		{
			name:     "user@host:port pattern",
			input:    "ssh admin@example.com:2222",
			wantHost: "example.com",
			wantPort: 2222,
			wantUser: "admin",
		},
		{
			name:     "user@host pattern",
			input:    "root@192.168.1.100",
			wantHost: "192.168.1.100",
			wantPort: 22,
			wantUser: "root",
		},
		{
			name:     "simple IP address pattern",
			input:    "connect 192.168.40.22",
			wantHost: "192.168.40.22",
			wantPort: 22,
			wantUser: "root",
		},
		{
			name:     "Chinese connection with IP",
			input:    "连接192.168.40.22",
			wantHost: "192.168.40.22",
			wantPort: 22,
			wantUser: "root",
		},
		{
			name:     "IP address in sentence",
			input:    "please connect to 10.0.0.1 server",
			wantHost: "10.0.0.1",
			wantPort: 22,
			wantUser: "root",
		},
		{
			name:    "no connection info",
			input:   "show me disk usage",
			wantNil: true,
		},
		{
			name:    "invalid IP 999.999.999.999",
			input:   "connect 999.999.999.999",
			wantNil: true,
		},
		{
			name:    "invalid IP 256.1.1.1",
			input:   "connect 256.1.1.1",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseConnectionDirect(tt.input)
			if tt.wantNil {
				if result != nil {
					t.Errorf("parseConnectionDirect() = %v, want nil", result)
				}
				return
			}
			if result == nil {
				t.Errorf("parseConnectionDirect() = nil, want non-nil")
				return
			}
			if result.Host != tt.wantHost {
				t.Errorf("Host = %q, want %q", result.Host, tt.wantHost)
			}
			if result.Port != tt.wantPort {
				t.Errorf("Port = %d, want %d", result.Port, tt.wantPort)
			}
			if result.User != tt.wantUser {
				t.Errorf("User = %q, want %q", result.User, tt.wantUser)
			}
		})
	}
}

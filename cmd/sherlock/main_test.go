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

package main

import (
	"testing"
)

func TestIsConnectionRequest(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// English keywords
		{name: "connect keyword", input: "connect to 192.168.1.100", want: true},
		{name: "ssh keyword", input: "ssh root@host.example.com", want: true},
		{name: "login keyword", input: "login to server", want: true},
		{name: "log in keyword", input: "log in to server", want: true},
		// Chinese keywords
		{name: "连接 keyword", input: "连接192.168.40.22", want: true},
		{name: "登录 keyword", input: "登录服务器", want: true},
		{name: "登陆 keyword", input: "登陆到服务器", want: true},
		// User@host pattern
		{name: "user@host pattern", input: "root@192.168.1.100", want: true},
		// IP address pattern
		{name: "simple IP address", input: "192.168.40.22", want: true},
		{name: "IP in sentence", input: "please help me with 10.0.0.1", want: true},
		// Not connection requests
		{name: "disk usage", input: "show me disk usage", want: false},
		{name: "list files", input: "list files in current directory", want: false},
		{name: "help command", input: "help", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isConnectionRequest(tt.input); got != tt.want {
				t.Errorf("isConnectionRequest(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

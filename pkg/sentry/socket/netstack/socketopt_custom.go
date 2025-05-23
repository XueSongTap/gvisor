// Copyright 2024 The gVisor Authors.
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

//go:build !false
// +build !false

package netstack

import (
	"gvisor.dev/gvisor/pkg/marshal"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	"gvisor.dev/gvisor/pkg/sentry/socket"
	"gvisor.dev/gvisor/pkg/syserr"
)

// setSockOptSocketCustom handles SetSockOpt options not handled by setSockOptSocket.
// It returns a bool indicating whether the option was handled in addition to
// return values from setSockOptSocket.
func setSockOptSocketCustom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, optVal []byte) (*syserr.Error, bool) {
	return nil, false
}

// getSockOptSocketCustom handles GetSockOpt options not handled by getSockOptSocket.
// It returns a bool indicating whether the option was handled in addition to
// return values from getSockOptSocket.
func getSockOptSocketCustom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, outLen int) (marshal.Marshallable, *syserr.Error, bool) {
	return nil, nil, false
}

func setSockOptIPCustom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, optVal []byte) (*syserr.Error, bool) {
	return nil, false
}

func setSockOptIPv6Custom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, optVal []byte) (*syserr.Error, bool) {
	return nil, false
}

func setSockOptTCPCustom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, optVal []byte) (*syserr.Error, bool) {
	return nil, false
}

func setSockOptICMPv6Custom(t *kernel.Task, s socket.Socket, ep commonEndpoint, name int, optVal []byte) (*syserr.Error, bool) {
	return nil, false
}

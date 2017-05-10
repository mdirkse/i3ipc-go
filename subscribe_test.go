// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"testing"
)

// TODO: fix this test
func TestInit(t *testing.T) {
	// Because StartEventListener sends events to the socket we have to start
	// the test socket before calling it. In other tests we can first call
	// GetIPCSocket(), get the socket, and then start the server end of the socket
	// with startTestIPCSocket to be sure that the test socket is created
	// (newConn being called) before something is written to it. Here we don't have
	// that assurance, we're relying on timing, which is kind of terrible, but seems
	// to work. Should refactor this to be provably correct though.
	go startTestIPCSocket(testMessages["subscribe"])

	StartEventListener()

	for _, s := range eventSockets {
		if !s.open {
			t.Error("Init failed: closed event socket found.")
		}
	}
	if len(eventSockets) != int(eventmax) {
		t.Errorf("Too much or not enough event sockets. Got %d, expected %d.\n",
			len(eventSockets), int(eventmax))
	}

	_, err := Subscribe(I3WorkspaceEvent)
	if err != nil {
		t.Errorf("Failed to subscribe: %f\n", err)
	}
	// TODO: A test to ensure that subscriptions work as intended.
}

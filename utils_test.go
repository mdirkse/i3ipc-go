package i3ipc

import (
	"encoding/binary"
	"net"
)

type i3Message struct {
	messageType MessageType
	payload     string
}

type testRequestResponse struct {
	req, res i3Message
}

var server, client net.Conn

// Init the test connSource
func init() {
	cs = &testConnSource{}
}

// Listens for input and the produces i3 message output according
// to the req/resp data it is given. Will repeat for every testRequestResponse
// in the array, so make sure the IPC operation you're testing has the same
// number of IPC calls as there are elements in the array.
func startTestIPCSocket(data []testRequestResponse) {
	defer server.Close()
	var tmp []byte
	length := make([]byte, 4)
	mType := make([]byte, 4)

	// For every pre-programmed request/response pair we listen to the pipe and then respond
	for _, trr := range data {
		tmp = make([]byte, 256)
		ipcMsg := make([]byte, 0)

		binary.LittleEndian.PutUint32(length, uint32(len(trr.res.payload)))
		binary.LittleEndian.PutUint32(mType, uint32(trr.res.messageType))

		for _, a := range [][]byte{[]byte(_Magic), length, mType, []byte(trr.res.payload)} {
			ipcMsg = append(ipcMsg, a...)
		}

		server.Read(tmp)
		server.Write(ipcMsg)
	}
}

type testConnSource struct{}

// Creates a new pipe that is used to supply test values to the socket
func (dsf *testConnSource) newConn() (net.Conn, error) {
	// If we get called again after already opening a connection
	// (during the event subscription test for instance)
	// then make sure to close the existing connection (if there
	// is one) so we don't leak connections
	if server != nil {
		server.Close()
	}

	server, client = net.Pipe()
	return client, nil
}

// Test messages used in the various unit test. The source for most of the test JSON is: http://i3wm.org/docs/ipc.html
var testMessages = map[string][]testRequestResponse{
	"bar": {{
		i3Message{I3GetBarConfig, ""},
		i3Message{I3GetBarConfig, "[\"bar-bxuqzf\"]"},
	}, {
		i3Message{I3GetBarConfig, "bar-bxuqzf"},
		i3Message{I3GetBarConfig, "{\"id\": \"bar-bxuqzf\",\"mode\": \"dock\",\"position\": \"bottom\",\"status_command\": \"i3status\",\"font\": \"-misc-fixed-medium-r-normal--13-120-75-75-C-70-iso10646-1\",\"workspace_buttons\": true,\"binding_mode_indicator\": true,\"verbose\": false,\"colors\": {\"background\": \"#c0c0c0\",\"statusline\": \"#00ff00\",\"focused_workspace_text\": \"#ffffff\",\"focused_workspace_bg\": \"#000000\"}}"},
	}},
	"command": {{
		i3Message{I3Command, "exec /bin/true"},
		i3Message{I3Command, "[{ \"success\": true }]"},
	}},
	"marks": {{
		i3Message{I3GetMarks, ""},
		i3Message{I3GetMarks, "[]"},
	}},
	"outputs": {{
		i3Message{I3GetOutputs, ""},
		i3Message{I3GetOutputs, "[{\"name\": \"LVDS1\",\"active\": true,\"current_workspace\": \"4\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 800}},{\"name\": \"VGA1\",\"active\": true,\"current_workspace\": \"1\",\"rect\": {\"x\": 1280,\"y\": 0,\"width\": 1280,\"height\": 1024}}]"},
	}},
	"subscribe": {
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
		{
			i3Message{I3Subscribe, ""},
			i3Message{I3Subscribe, "{ \"success\": true }"},
		},
	},
	"tree": {{
		i3Message{I3GetTree, ""},
		i3Message{I3GetTree, "{\"id\": 6875648,\"name\": \"root\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 800},\"nodes\": [{\"id\": 6878320,\"name\": \"LVDS1\",\"layout\": \"output\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 800},\"nodes\": [{\"id\": 6878784,\"name\": \"topdock\",\"layout\": \"dockarea\",\"orientation\": \"vertical\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 0}},{\"id\": 6879344,\"name\": \"content\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 782},\"nodes\": [{\"id\": 6880464,\"name\": \"1\",\"orientation\": \"horizontal\",\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 782},\"floating_nodes\": [],\"nodes\": [{\"id\": 6929968,\"name\": \"#aa0000\",\"border\": \"normal\",\"percent\": 1,\"rect\": {\"x\": 0,\"y\": 18,\"width\": 1280,\"height\": 782}}]}]},{\"id\": 6880208,\"name\": \"bottomdock\",\"layout\": \"dockarea\",\"orientation\": \"vertical\",\"rect\": {\"x\": 0,\"y\": 782,\"width\": 1280,\"height\": 18},\"nodes\": [{\"id\": 6931312,\"name\": \"#00aa00\",\"percent\": 1,\"rect\": {\"x\": 0,\"y\": 782,\"width\": 1280,\"height\": 18}}]}]}]}"},
	}},
	"version": {{
		i3Message{I3GetVersion, ""},
		i3Message{I3GetVersion, "{\"human_readable\" : \"4.2-169-gf80b877 (2012-08-05, branch \\\"next\\\")\",\"loaded_config_file_name\" : \"/home/hwangcc23/.i3/config\",\"minor\" : 2,\"patch\" : 0,\"major\" : 4}"},
	}},
	"workspaces": {{
		i3Message{I3GetWorkspaces, ""},
		i3Message{I3GetWorkspaces, "[{\"num\": 0,\"name\": \"1\",\"visible\": true,\"focused\": true,\"urgent\": false,\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 800},\"output\": \"LVDS1\"},{\"num\": 1,\"name\": \"2\",\"visible\": false,\"focused\": false,\"urgent\": false,\"rect\": {\"x\": 0,\"y\": 0,\"width\": 1280,\"height\": 800},\"output\": \"LVDS1\"}]"},
	}},
}

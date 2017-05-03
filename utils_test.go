package i3ipc

import (
	"encoding/binary"
	"net"
)

type i3Message struct {
	messageType MessageType
	payload     string
}

type TestRequestResponse struct {
	req, res i3Message
}

func GetTestIPC(data []TestRequestResponse) *IPCSocket {
	server, client := net.Pipe()

	go func() {
		defer server.Close()
		var tmp []byte
		length := make([]byte, 4)
		mType := make([]byte, 4)

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
	}()

	ipc := &IPCSocket{}
	ipc.socket = client
	ipc.open = true

	return ipc
}

// Test messages used in the various unit test
var testMessages = map[string][]TestRequestResponse{
	"bar": {{
		i3Message{I3GetBarConfig, ""},
		i3Message{I3GetBarConfig, "[\"bar-bxuqzf\"]"},
	}, {
		i3Message{I3GetBarConfig, "bar-bxuqzf"},
		i3Message{I3GetBarConfig, "{\"id\": \"bar-bxuqzf\",\"mode\": \"dock\",\"position\": \"bottom\",\"status_command\": \"i3status\",\"font\": \"-misc-fixed-medium-r-normal--13-120-75-75-C-70-iso10646-1\",\"workspace_buttons\": true,\"binding_mode_indicator\": true,\"verbose\": false,\"colors\": {\"background\": \"#c0c0c0\",\"statusline\": \"#00ff00\",\"focused_workspace_text\": \"#ffffff\",\"focused_workspace_bg\": \"#000000\"}}"},
	}},
}

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
	"encoding/json"
)

// I3Bar represents the configuration of a bar. For documentation of the
// fields, refer to http://i3wm.org/docs/ipc.html#_bar_config_reply.
type I3Bar struct {
	ID                   string
	Mode                 string
	Position             string
	StatusCommand        string `json:"status_command"`
	Font                 string
	WorkspaceButtons     bool `json:"workspace_buttons"`
	BindingModeIndicator bool `json:"binding_mode_indicator"`
	Verbose              bool
	Colors               Colors
}

// Colors represents colors as used in I3Bar.
type Colors struct {
	Background              string
	Statusline              string
	Separator               string
	FocusedBackground       string `json:"focused_background"`
	FocusedStatusline       string `json:"focused_statusline"`
	FocusedSeparator        string `json:"focused_separator"`
	FocusedWorkspaceBg      string `json:"focused_workspace_bg"`
	FocusedWorkspaceBorder  string `json:"focused_workspace_border"`
	FocusedWorkspaceText    string `json:"focused_workspace_text"`
	ActiveWorkspaceBg       string `json:"active_workspace_bg"`
	ActiveWorkspaceBorder   string `json:"active_workspace_border"`
	ActiveWorkspaceText     string `json:"active_workspace_text"`
	InactiveWorkspaceBg     string `json:"inactive_workspace_bg"`
	InactiveWorkspaceBorder string `json:"inactive_workspace_border"`
	InactiveWorkspaceText   string `json:"inactive_workspace_text"`
	UrgentWorkspaceBg       string `json:"urgent_workspace_bg"`
	UrgentWorkspaceBorder   string `json:"urgent_workspace_border"`
	UrgentWorkspaceText     string `json:"urgent_workspace_text"`
	BindingModeBg           string `json:"binding_mode_bg"`
	BindingModeBorder       string `json:"binding_mode_border"`
	BindingModeText         string `json:"binding_mode_text"`
}

// GetBarIds fetches a list of IDs for all configured bars.
func (socket *IPCSocket) GetBarIds() (ids []string, err error) {
	jsonReply, err := socket.Raw(I3GetBarConfig, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &ids)
	return
}

// GetBarConfig returns the configuration of the bar with the given ID.
func (socket *IPCSocket) GetBarConfig(id string) (bar I3Bar, err error) {
	jsonReply, err := socket.Raw(I3GetBarConfig, id)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReply, &bar)
	return
}

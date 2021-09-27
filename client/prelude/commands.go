package prelude

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/rpcpb"
	"github.com/bishopfox/sliver/protobuf/sliverpb"
)

//RunCommand executes a given command
func RunCommand(message string, executor string, payload []byte, agent *AgentConfig, rpc rpcpb.SliverRPCClient, session *clientpb.Session) (string, int, int) {
	switch executor {
	case "keyword":
		task := splitMessage(message, '.')
		switch task[0] {
		case "bof":
			if len(task) != 2 {
				break
			}
			var bargs []bofArgs
			err := json.Unmarshal([]byte(task[1]), &bargs)
			if err != nil {
				return err.Error(), ErrorExitStatus, ErrorExitStatus
			}
			if payload == nil {
				return "missing BOF file", ErrorExitStatus, ErrorExitStatus
			}
			out, err := runBOF(session, rpc, payload, bargs)
			if err != nil {
				return err.Error(), ErrorExitStatus, ErrorExitStatus
			}
			return out, 0, 0
		case "exit":
			return shutdown(agent, rpc, session)
		default:
			return "Keyword selected not available for agent", ErrorExitStatus, ErrorExitStatus
		}
	default:
		bites, status, pid := execute(message, executor, agent, rpc, session)
		return string(bites), status, pid
	}
	return "", ErrorExitStatus, ErrorExitStatus
}

func execute(cmd string, executor string, agentConfig *AgentConfig, rpc rpcpb.SliverRPCClient, session *clientpb.Session) (string, int, int) {
	args := append(getCmdArg(executor), cmd)
	if executor == "psh" {
		executor = "powershell.exe"
	}
	execResp, err := rpc.Execute(context.Background(), &sliverpb.ExecuteReq{
		Path:    executor,
		Args:    args,
		Output:  true,
		Request: MakeRequest(session),
	})

	if err != nil {
		return fmt.Sprintf("Error: %s\n", err.Error()), 0, 0
	}

	if execResp.Response != nil && execResp.Response.Err != "" {
		return execResp.Response.Err, 0, 0
	}
	return execResp.Stdout, int(execResp.Status), int(execResp.Pid)
}

func getCmdArg(executor string) []string {
	var args []string
	switch executor {
	case "cmd":
		args = []string{"/C", "/S"}
	case "powershell", "psh":
		args = []string{"-execu", "ByPasS", "-C"}
	case "sh", "bash", "zsh":
		args = []string{"-c"}
	}
	return args
}

func splitMessage(message string, splitRune rune) []string {
	quoted := false
	values := strings.FieldsFunc(message, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && r == splitRune
	})
	return values
}

func shutdown(cfg *AgentConfig, rpc rpcpb.SliverRPCClient, session *clientpb.Session) (string, int, int) {
	_, err := rpc.KillSession(context.Background(), &sliverpb.KillSessionReq{
		Force:   false,
		Request: MakeRequest(session),
	})
	if err != nil {
		return err.Error(), ErrorExitStatus, ErrorExitStatus
	}
	return fmt.Sprintf("Terminated %s", session.Name), SuccessExitStatus, SuccessExitStatus
}

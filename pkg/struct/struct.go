package go_struct

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Name   string `json:"name"`
	Prompt string `json:"prompt"`
	Active bool   `json:"active"`
}

func (a *Agent) SetAgent(name, prompt string, active bool) {
	a.Name = name
	a.Prompt = prompt
	a.Active = active
}

func DemoStruct() {
	agent1 := Agent{
		Name:   "Agent Smith",
		Prompt: "Infiltrate the Matrix",
		Active: true,
	}
	fmt.Println("agent1:", agent1)

	var agent2 Agent
	agent2.SetAgent("Agent Johnson", "Protect the system", false)
	fmt.Println("agent2:", agent2)

	data, err := json.Marshal(agent1)
	if err != nil {
		fmt.Println("Error marshaling agent1:", err)
		return
	}
	fmt.Println("agent1 in JSON format:", string(data))
}

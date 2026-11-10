package gui

import "github.com/yetone/magpie/internal/agent"

// displayClients adds RavenAI's existing request identity for display only.
// Keeping it out of agent.Clients leaves gateway classification and routing
// rules untouched, including for records already named "RavenAI".
func displayClients() []*agent.Agent {
	return append(agent.Clients(), &agent.Agent{ID: "RavenAI", Name: "RavenAI", Icon: "ravenai"})
}

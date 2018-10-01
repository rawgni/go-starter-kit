package main

import (
	"github.com/srdhoni/go-starter-kit/agent"
	"github.com/srdhoni/go-starter-kit/agent/first"
	"github.com/srdhoni/go-starter-kit/agent/via"
	"github.com/srdhoni/go-starter-kit/apicancel"
)

func main() {

	m := make(map[agent.AgentID]interface{}, 0)

	m[agent.IDFirst] = first.AgentFirst{}
	m[agent.IDVia] = via.AgentVia{}

	apicancel.HandlerGet(m)
	// apicancel.HandlerEstimate()
}

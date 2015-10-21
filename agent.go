package goapi

import (
	"fmt"
	"strings"
)

func (c *Client) AgentList() ([]Agent, error) {
	v := []Agent{}
	err := c.api.Get(defaultContext(), c.pathTo("/agents"), nil, &v)
	return v, err
}

func (c *Client) AgentEnable(uuid string) error {
	path := c.pathTo("/agents/%s/enable", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) AgentDisable(uuid string) error {
	path := c.pathTo("/agents/%s/disable", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) AgentDelete(uuid string) error {
	path := c.pathTo("/agents/%s/delete", uuid)
	return c.api.Post(defaultContext(), path, nil, nil)
}

func (c *Client) AgentsDeleteLostAndMissing() ([]Agent, error) {
	allAgents := []Agent{}
	oldAgents := []Agent{}
	err := c.api.Get(defaultContext(), c.pathTo("/agents"), nil, &allAgents)
	for index, agent := range allAgents {
		fmt.Printf("%d: Agent %s (%s) is %s", index, agent.AgentName, agent.IpAddress, agent.Status)
		if strings.HasPrefix(agent.IpAddress, "172.17") && agent.Status == "Missing" || agent.Status == "LostContact" {
			fmt.Printf(" - trying to disable: ")
			c.AgentDisable(agent.Uuid)
			fmt.Printf("DONE - trying to delete: ")
			c.AgentDelete(agent.Uuid)
			fmt.Printf("DONE.\n")
			oldAgents = append(oldAgents, agent)
		} else {
			fmt.Printf(" - ignore keeping agent.\n")
		}
	}
	return oldAgents, err
}

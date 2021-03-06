package main

import "github.com/codegangsta/cli"

var stage = cli.Command{
	Name:  "stage",
	Usage: "commands for stages",
	Subcommands: []cli.Command{
		{
			Name:   "cancel",
			Usage:  "cancel a specific stage",
			Flags:  append(flags, flagPipelineName, flagStageName),
			Action: stageCancel,
		},
	},
}

func stageCancel(c *cli.Context) {
	client := newClient(c)

	pipelineName := c.String("pipeline")
	stageName := c.String("stage")
	err := client.StageCancel(pipelineName, stageName)
	assert(err)
}

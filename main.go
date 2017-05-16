package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "interact with the Jenkins API"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "abort-input",
			Usage: "abort a pipeline input",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "job-name, n",
					Usage: "name of the job where the input resides",
				},
				cli.StringFlag{
					Name:  "build-number, b",
					Usage: "the build number of the project awaiting input",
				},
				cli.StringFlag{
					Name:  "input-id, i",
					Usage: "the id of the input, specified in the Jenkinsfile",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("abort-input")
				return nil
			},
		},
		{
			Name:  "copy-job",
			Usage: "copy a job from another",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "old-job-name, o",
					Usage: "name of the job to copy",
				},
				cli.StringFlag{
					Name:  "new-job-name, n",
					Usage: "name of the new job to be created",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("copy-job")
				return nil
			},
		},
		{
			Name:  "create-job",
			Usage: "create a job",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "job-name, n",
					Usage: "name of the job to be created",
				},
				cli.StringFlag{
					Name:  "project-url, p",
					Usage: "browser link of project repository",
				},
				cli.StringFlag{
					Name:  "git-url, g",
					Usage: "git url of project",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("create-job")
				return nil
			},
		},
		{
			Name:  "env",
			Usage: "print environment information",
			Action: func(c *cli.Context) error {
				fmt.Println("env")
				return nil
			},
		},
		{
			Name:  "trigger-input",
			Usage: "trigger a pipeline input",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "job-name, n",
					Usage: "name of the job where the input resides",
				},
				cli.StringFlag{
					Name:  "build-number, b",
					Usage: "the build number of the build awaiting input",
				},
				cli.StringFlag{
					Name:  "input-id, i",
					Usage: "the id of the input, specified in the Jenkinsfile",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("input-id")
				return nil
			},
		},
	}

	app.Run(os.Args)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

type webHook struct {
	Ref        string   `json:"ref"`
	Commits    []string `json:"commits"`
	HeadCommit struct {
		ID string `json:"id"`
	} `json:"head_commit"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Deleted bool `json:"deleted"`
}

func (w webHook) Summary() string {
	return fmt.Sprintf("build on %s on branch %s (%s)",
		w.Repository.FullName,
		w.Ref,
		filepath.Base(w.HeadCommit.ID))
}

const webHookEndpoint = "https://build.grid.tf/hook/monitor-watch"

func main() {

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "organization,o",
			Value: "threefoldtech",
			Usage: "github organization",
		},
		cli.StringFlag{
			Name:  "repository,r",
			Usage: "repository name",
		},
		cli.StringFlag{
			Name:  "commit,c",
			Usage: "commit hash",
		},
		cli.StringFlag{
			Name:  "branch,b",
			Usage: "branch to use for the build",
			Value: "master",
		},
		cli.StringFlag{
			Name:  "path,p",
			Usage: "specified the path of the repository to use as source of information for the build",
			Value: "",
		},
	}

	app.Action = func(c *cli.Context) error {
		var (
			wh  webHook
			err error
		)

		if c.String("path") != "" {
			path := c.String("path")
			wh, err = Parse(path)
			if err != nil {
				log.Fatalf("fail to read the git information from %s: %v", path, err)
			}
		} else {
			wh.Repository.FullName = fmt.Sprintf("%s/%s", c.String("organization"), c.String("repository"))
			wh.Ref = fmt.Sprintf("ref/head/%s", c.String("branch"))
			wh.HeadCommit.ID = c.String("commit")
			wh.Commits = []string{}
		}

		fmt.Println(wh.Summary())

		body := bytes.Buffer{}
		if err := json.NewEncoder(&body).Encode(wh); err != nil {
			log.Fatalf("fail to serialize the webhook object: %v", err)
		}

		request, err := http.NewRequest("POST", webHookEndpoint, &body)
		if err != nil {
			log.Fatalf("error while creating the http request :%v", err)
		}
		request.Header.Set("X-Github-Event", "push")
		request.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Fatalf("fail to send request to autobuilder: %v", err)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("error code received from the auto builder: %v", resp.Status)
		}

		fmt.Println("build started")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

const (
	builderBaseURL  = "https://build.grid.tf"
	webHookEndpoint = builderBaseURL + "/hook/monitor-watch"
)

func main() {
	app := cli.NewApp()
	app.ArgsUsage = "[path|ghrepo]"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "commit,c",
			Usage: "commit hash",
		},
		cli.StringFlag{
			Name:  "branch,b",
			Usage: "branch to use for the build",
			Value: "master",
		},
	}

	app.Action = func(c *cli.Context) error {
		var (
			wh  webHook
			err error
		)

		// path can be either empty, a FS path to a GitHub Repo, or the fullname of a GitHub repo (organization/repo)
		path := c.Args().Get(0)
		switch {
		case path == "":
			path, err = os.Getwd()
			if err != nil {
				log.Fatalf("fail to get the current working directory as fs path: %v", err)
			}
			fallthrough
		case isDir(path):
			wh, err = Parse(path)
			if err != nil {
				log.Fatalf("fail to read the git information from %s: %v", path, err)
			}
		default: // assume path is in the form of organization/repo
			if strings.Count(path, "/") != 1 {
				log.Fatalf("invalid GitHub fullname: %s", path)
			}
			wh.Repository.FullName = path
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
		fmt.Printf("go to %s to follow the progress of your build\n", builderBaseURL)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.Mode().IsDir()
}

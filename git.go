package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

func Parse(path string) (wh webHook, err error) {
	wh.Commits = []string{}

	wh.HeadCommit.ID, err = commitID(path)
	if err != nil {
		return webHook{}, err
	}
	branch, err := branch(path)
	if err != nil {
		return webHook{}, err
	}
	wh.Ref = fmt.Sprintf("ref/head/%s", branch)

	repo, err := repository(path)
	if err != nil {
		return webHook{}, err
	}
	wh.Repository.FullName = repo

	return
}

func commitID(path string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func branch(path string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func repository(path string) (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = path
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	remote := string(out)

	i := strings.Index(remote, ":")
	if i == -1 {
		u, err := url.Parse(string(out))
		if err != nil {
			return "", err
		}
		repository := strings.SplitAfterN(u.Path, "/", 1)
		return strings.Join(repository, "/"), nil
	}

	return remote[i+1 : len(remote)-5], nil

}

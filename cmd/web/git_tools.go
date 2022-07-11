package main

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

type Commit struct {
	Hash     string `json:"hash"`
	Message  string `json:"message"`
	TaskCode string `json:"taskCode,omitempty"`
	PRNumber string `json:"PRNumber,omitempty"`
}

func parseGitLine(s string) Commit {
	c := Commit{}
	fields := strings.Fields(s)
	hash, rest := fields[0], fields[1:]
	c.Hash = hash
	c.Message = strings.Join(rest, " ")
	return c
}

func GetHistory(path string, maxEntries int) []Commit {
	cmd := exec.Command("git", "-C", path, "log", "--oneline", "--tags")
	val, err := cmd.Output()
	if err != nil {
		log.Fatalln("Could not execute cmd")
	}
	scanner := bufio.NewScanner(strings.NewReader(string(val)))
	var commits []Commit
	for scanner.Scan() {
		c := parseGitLine(scanner.Text())
		commits = append(commits, c)
		if len(commits) > maxEntries {
			break
		}
	}
	return commits
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
)

type GitLab struct {
	git           *gitlab.Client
	slackHook     string
	slackChannel  string
	slackUsername string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gitlab_url := os.Getenv("GITLAB_URL")
	gitlab_token := os.Getenv("GITLAB_TOKEN")
	gitlab_projects := strings.Split(os.Getenv("GITLAB_PROJECTS"), ",")
	slack_hook := os.Getenv("SLACK_HOOK")
	slack_channel := os.Getenv("SLACK_CHANNEL")

	client, err := gitlab.NewClient(gitlab_token, gitlab.WithBaseURL(gitlab_url+"/api/v4"))

	gl := &GitLab{
		git:          client,
		slackHook:    slack_hook,
		slackChannel: slack_channel,
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	for _, project := range gitlab_projects {
		gl.getOldMergeRequests(project)
	}

}

func (gl GitLab) getOldMergeRequests(project string) {

	now := time.Now()
	week_ago := now.AddDate(0, 0, -7)

	git := gl.git

	gitlab_project, _, err := git.Projects.GetProject(project, nil)

	if err != nil {
		log.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	opt := &gitlab.ListProjectMergeRequestsOptions{
		State:         gitlab.String("opened"),
		UpdatedBefore: &week_ago,
	}

	mergeRequests, _, err := git.MergeRequests.ListProjectMergeRequests(gitlab_project.ID, opt)

	if err != nil {
		log.Fatalf("Getting merge requests returns an error: %v", err)
	}

	for _, mr := range mergeRequests {

		layout := "2006-01-02 15:04:05.000 -0700 MST"
		t, _ := time.Parse(layout, mr.UpdatedAt.String())
		attachment1 := slack.Attachment{}

		attachment1.AddField(slack.Field{Title: "Author", Value: mr.Author.Name})
		attachment1.AddAction(slack.Action{Type: "button", Text: fmt.Sprintf("Merge %s", mr.Reference), Url: mr.WebURL, Style: "primary"})
		attachment1.AddField(slack.Field{Title: "Updated at", Value: t.Format("2006.01.02")})

		payload := slack.Payload{
			Text:        "*" + mr.Title + "*",
			Username:    gl.slackUsername,
			Channel:     gl.slackChannel,
			IconEmoji:   ":dansu:",
			Attachments: []slack.Attachment{attachment1},
		}

		err := slack.Send(gl.slackHook, "", payload)
		if len(err) > 0 {
			fmt.Printf("error: %s\n", err)
		}

	}

}

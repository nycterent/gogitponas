package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gitlab_url := os.Getenv("GITLAB_URL")
	gitlab_token := os.Getenv("GITLAB_TOKEN")
	gitlab_projects := strings.Split(os.Getenv("GITLAB_PROJECTS"), ",")

	// slack_hook := os.Getenv("SLACK_HOOK")
	// slack_channel := os.Getenv("SLACK_CHANNEL")

	g := gitlab.New(gitlab_token, gitlab_url)

	notification_targets := registry.New(strings.Split(os.Getenv("NOTIFICATION_TARGETS"), ","))

	notification_targets.Register("slack", &Slack{})
	notification_targets.Register("rocket", &Rocket{rc: rocketchat.New(os.Getenv("ROCKET_HOOK"))})

	for _, project := range gitlab_projects {
		for _, mr := range g.GetOldMergeRequests(project) {
			notification_targets.Send(mr)
		}
	}

}

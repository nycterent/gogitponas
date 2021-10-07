package main

import (
	"gogitponas/gitlab"
	"gogitponas/registry"
	"gogitponas/rocketchat"
	"gogitponas/slack"
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

	gitlabURL := os.Getenv("GITLAB_URL")
	gitlabToken := os.Getenv("GITLAB_TOKEN")
	gitlabProjects := strings.Split(os.Getenv("GITLAB_PROJECTS"), ",")

	g := gitlab.New(gitlabToken, gitlabURL)

	notificationTargets := registry.New(strings.Split(os.Getenv("NOTIFICATION_TARGETS"), ","))

	notificationTargets.Register("slack", &Slack{sc: slack.New(os.Getenv("SLACK_HOOK"))})
	notificationTargets.Register("rocket", &Rocket{rc: rocketchat.New(os.Getenv("ROCKET_HOOK"))})

	for _, project := range gitlabProjects {
		for _, mr := range g.GetOldMergeRequests(project) {
			notificationTargets.Send(mr)
		}
	}

}

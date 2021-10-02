package gitlab

import (
	"log"
	"time"

	"github.com/xanzy/go-gitlab"
)

func New(GitlabToken, GitlabUrl string) *GitLab {
	client, err := gitlab.NewClient(GitlabToken, gitlab.WithBaseURL(GitlabUrl+"/api/v4"))

	gl := &GitLab{
		git: client,
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return gl
}

type GitLab struct {
	git *gitlab.Client
}

func (gl GitLab) GetOldMergeRequests(project string) (OldMerges []GitlabMergeInformation) {

	now := time.Now()
	week_ago := now.AddDate(0, 0, 0)

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
		t, err := time.Parse(layout, mr.UpdatedAt.String())

		if err != nil {
			layout = "2006-01-02 15:04:05.00 -0700 MST"
			t, _ = time.Parse(layout, mr.UpdatedAt.String())
		}

		OldMerges = append(OldMerges, GitlabMergeInformation{
			Title:     mr.Title,
			Author:    mr.Author.Name,
			Reference: mr.Reference,
			MRUrl:     mr.WebURL,
			UpdatedAt: t.Format("2006.01.02"),
		})

		// attachment1 := slack.Attachment{}

		// attachment1.AddField(slack.Field{Title: "Author", Value: mr.Author.Name})
		// attachment1.AddAction(slack.Action{Type: "button", Text: fmt.Sprintf("Merge %s", mr.Reference), Url: mr.WebURL, Style: "primary"})
		// attachment1.AddField(slack.Field{Title: "Updated at", Value: t.Format("2006.01.02")})

		// payload := slack.Payload{
		// 	Text:        "*" + mr.Title + "*",
		// 	Username:    gl.slackUsername,
		// 	Channel:     gl.slackChannel,
		// 	IconEmoji:   ":dansu:",
		// 	Attachments: []slack.Attachment{attachment1},
		// }

		// err := slack.Send(gl.slackHook, "", payload)
		// if len(err) > 0 {
		// 	fmt.Printf("error: %s\n", err)
		// }

	}
	return
}

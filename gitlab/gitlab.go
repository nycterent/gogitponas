package gitlab

import (
	"log"
	"time"

	"github.com/xanzy/go-gitlab"
)

func New(GitlabToken, GitlabURL string) *Gitlab {
	client, err := gitlab.NewClient(GitlabToken, gitlab.WithBaseURL(GitlabURL+"/api/v4"))

	gl := &Gitlab{
		git: client,
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return gl
}

type Gitlab struct {
	git *gitlab.Client
}

func (gl Gitlab) GetOldMergeRequests(project string) (OldMerges []GitlabMergeInformation) {

	now := time.Now()
	week_ago := now.AddDate(0, 0, 0)

	git := gl.git

	gitlabProject, _, err := git.Projects.GetProject(project, nil)

	if err != nil {
		log.Fatalf("Projects.GetProject returns an error: %v", err)
	}

	opt := &gitlab.ListProjectMergeRequestsOptions{
		State:         gitlab.String("opened"),
		UpdatedBefore: &week_ago,
	}

	mergeRequests, _, err := git.MergeRequests.ListProjectMergeRequests(gitlabProject.ID, opt)

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
			MRURL:     mr.WebURL,
			UpdatedAt: t.Format("2006.01.02"),
		})
	}
	return
}

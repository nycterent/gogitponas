// attachment1.AddField(slack.Field{Title: "Author", Value: mr.Author.Name})
// attachment1.AddAction(slack.Action{Type: "button", Text: fmt.Sprintf("Merge %s", mr.Reference), Url: mr.WebURL, Style: "primary"})
// attachment1.AddField(slack.Field{Title: "Updated at", Value: t.Format("2006.01.02")})

package gitlab

type GitlabMergeInformation struct {
	Title     string
	Author    string
	Reference string
	MRUrl     string
	UpdatedAt string
}

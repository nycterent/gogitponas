package gitlab

// MergeInformation holds required information for Gitlab merge
type MergeInformation struct {
	Title     string
	Author    string
	Reference string
	MRURL     string
	UpdatedAt string
}

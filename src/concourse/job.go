package concourse

type Job struct {
	Name          string
	Url           string
	FinishedBuild FinishedBuild `json:"finished_build"`
}

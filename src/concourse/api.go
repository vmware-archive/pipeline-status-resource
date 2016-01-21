package concourse

import (
	"encoding/json"
	"fmt"
)

type Api interface {
	GetPipelines() []Pipeline
	GetJobs(pipeline string) []Job
}

type ApiClient struct {
	RequestDoer RequestDoer
}

func NewApi(requestDoer RequestDoer) Api {
	return &ApiClient{RequestDoer: requestDoer}
}

func (api *ApiClient) GetPipelines() []Pipeline {
	resp := api.RequestDoer.DoRequest("GET", "/v1/pipelines")
	defer resp.Body.Close()

	var pipelines []Pipeline
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&pipelines)
	if err != nil {
		panic(err)
	}

	return pipelines
}

func (api *ApiClient) GetJobs(pipeline string) []Job {
	resp := api.RequestDoer.DoRequest("GET", fmt.Sprintf("/v1/pipelines/%s/jobs", pipeline))
	defer resp.Body.Close()

	var jobs []Job
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&jobs)
	if err != nil {
		panic(err)
	}

	return jobs
}

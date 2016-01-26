package main

import (
  "concourse"
  "fmt"
  "io/ioutil"
  "flag"
  "encoding/json"
  "strconv"
)

func main() {
  hostPtr := flag.String("host", "", "Concourse URL")
  userPtr := flag.String("user", "", "Concourse username")
  passwordPtr := flag.String("password", "", "Concourse password")

  flag.Parse()

  CONCOURSE_URL := *hostPtr
  USER := *userPtr
  PASSWORD := *passwordPtr
  PIPELINE_WHITELIST := flag.Args()

  rd := concourse.NewRequestDoer(
    CONCOURSE_URL,
    USER,
    PASSWORD,
  )

  api := concourse.NewApi(rd)
  ps := api.GetPipelines()

  var failingBuilds []concourse.MetadataElement

  for _, p := range ps {
    if !p.Paused {
      for _, pipelineInWhitelist := range PIPELINE_WHITELIST {
        if p.Name == pipelineInWhitelist {
          jobs := api.GetJobs(p.Name)
          for _, j := range jobs {
            s := j.FinishedBuild.Status
            if s != "" && s != "succeeded" {
              buildId := strconv.Itoa(j.FinishedBuild.Id)
              statusAndName := j.FinishedBuild.Status + " = " + j.Name
              failingBuilds = append(failingBuilds, concourse.MetadataElement{buildId, statusAndName})
            }
          }
        }
      }
    }
  }

  if len(failingBuilds) == 0 {
    fmt.Println("All Jobs have Succeeded")
  } else {
    fmt.Println("Failing jobs detected:")
    for _, f := range failingBuilds {
      fmt.Println(f)
    }
  }

  bytes, _ := json.Marshal(failingBuilds)

  err := ioutil.WriteFile("failing-builds.json", bytes, 0644)

  if err != nil {
    panic(err)
  }
}

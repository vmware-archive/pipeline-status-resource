package main

import (
  "concourse"
  "fmt"
  "io/ioutil"
  "flag"
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

  var failingBuildNames []string
  var failingBuildIds []int

  for _, p := range ps {
    if !p.Paused {
      for _, pipelineInWhitelist := range PIPELINE_WHITELIST {
        if p == pipelineInWhitelist {
          jobs := api.GetJobs(p.Name)
          for _, j := range jobs {
            s := j.FinishedBuild.Status
            if s != "" && s != "succeeded" {
              failingBuildNames = append(failingBuildNames, fmt.Sprintf("%s = %s", j.Name, s))
              failingBuildIds = append(failingBuildIds, j.Id)
            }
          }
        }
      }
    }
  }

  var jsonString string

  if len(failures) == 0 {
    fmt.Println("All Jobs have Succeeded")
  } else {
    fmt.Println("Failing jobs detected:")
    for _, f := range failures {
      fmt.Println(f)
    }
  }

  jsonString = json.Marshal(failingBuildIds)

  bytes := []byte(jsonString)

  err := ioutil.WriteFile("failing-builds.json", bytes, 0644)

  if err != nil {
    panic(err)
  }
}

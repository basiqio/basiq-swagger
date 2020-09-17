package jobs

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/jobs"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"testing"
)

func TestGetJob(t *testing.T) {
	jobParams := &jobs.GetJobsParams{
		JobID:   "b1ef29c9-3f8c-45c9-a823-d819605989c6",
		Context: context.TODO(),
	}

	jobRsp, err := test.Client.Jobs.GetJobs(jobParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting job: %v", err)
	}

	e, err := json.Marshal(jobRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getJob.json", t)

	test.AssertJson(t, s, string(e))
}

// TODO This test needs to be done differently because the GET jobs are getting back jobs only for the previous 7 days

//func TestGetJobs(t *testing.T) {
//	jobsParams := &jobs.GetUserJobsParams{
//		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
//		Context: context.TODO(),
//	}
//
//	jobsRsp, err := test.Client.Jobs.GetUserJobs(jobsParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
//	if err != nil {
//		t.Fatalf("Error getting user jobs: %v", err)
//	}
//
//	e, err := json.Marshal(jobsRsp.GetPayload())
//	if err != nil {
//		t.Fatalf("Error: %v", err)
//	}
//
//	s := test.GetJsonResponse("./responses/getJobs.json", t)
//
//	test.AssertJson(t, s, string(e))
//}

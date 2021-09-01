package jobs

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/connections"
	"github.com/basiqio/basiq-swagger/dist/client/jobs"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
	"time"
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

func TestGetJobs(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := test.SetupUser(t)
	defer test.CleanupUser(t, userID)

	loginID := "Wentworth-Smith"
	password := "whislter"
	institutionID := "AU00000"
	institution := models.InstitutionModel{ID: &institutionID}

	connCreateParams := &connections.PostConnectionParams{
		UserID:                  userID,
		UserConnectionsPostData: &models.UserConnectionsPostData{Password: &password, LoginID: &loginID, Institution: &institution},
		Context:                 context.TODO(),
	}

	connectionPostRsp, err := test.Client.Connections.PostConnection(connCreateParams, token)

	if err != nil {
		t.Fatalf("Error posting connection, Error: %v", err)
	}

	var jobsPayload *models.JobsResponseResource
	for i := 0; i <= 15; i++ {
		jobsPayload, err = getJobsPayload(userID, token)
		if err != nil {
			t.Fatalf("Error getting user jobs: %v", err)
		}
		if *jobsPayload.Data[0].Steps[2].Status == models.JobsStepStatusSuccess {
			continue
		}
		time.Sleep(1 * time.Second)
	}

	s := test.GetJsonResponse("./responses/getJobs.json", t)
	s = strings.Replace(s, "5b3532e0-23c7-48cb-b7a4-ff37f46c3eb1", userID, 5)
	s = strings.Replace(s, "a826e470-83d5-482a-b19f-bb44985c14bb", *connectionPostRsp.GetPayload().ID, 2)
	s = strings.Replace(s, "2020-09-21T11:19:11Z", *jobsPayload.Data[0].Created, 1)
	s = strings.Replace(s, "2020-09-21T11:19:11Z", *jobsPayload.Data[0].Updated, 1)
	s = strings.Replace(s, "{source}", jobsPayload.Data[0].Links.Source, 1)
	s = strings.Replace(s, "{url}", jobsPayload.Data[0].Steps[0].Result.URL, 1)

	e, err := json.Marshal(jobsPayload)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	test.AssertJson(t, s, string(e))
}

func getJobsPayload(userID string, token runtime.ClientAuthInfoWriter) (*models.JobsResponseResource, error) {
	jobsParams := &jobs.GetUserJobsParams{
		UserID:  userID,
		Context: context.TODO(),
	}
	jobsRsp, err := test.Client.Jobs.GetUserJobs(jobsParams, token)
	if err != nil {
		return nil, err
	}

	return jobsRsp.GetPayload(), nil
}

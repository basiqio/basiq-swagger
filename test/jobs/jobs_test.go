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
	"github.com/stretchr/testify/assert"
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
	for i := 0; i < 15; i++ {
		jobsPayload, err = getJobsPayload(userID, token)
		if err != nil {
			t.Fatalf("Error getting user jobs: %v", err)
		}
		if validateSteps(jobsPayload.Data[0].Steps) {
			break
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

func TestPostJobsMfa(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := test.SetupUser(t)

	timeout := time.Minute * 2
	interval := time.Second * 5

	defer test.CleanupUser(t, userID)

	loginID := "gavinBelson"
	password := "hooli2016"
	institutionID := "AU00002"
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

	ticker := time.NewTicker(interval)
	end := time.NewTicker(timeout)
	allStepsSuccess := false

	for !allStepsSuccess {
		select {
		case <-end.C:
			t.Fatalf("MFA connection check timeouted")
		case <-ticker.C:
			jobParams := &jobs.GetJobsParams{
				JobID:   *connectionPostRsp.Payload.ID,
				Context: context.TODO(),
			}

			jobsResponse, err := test.Client.Jobs.GetJobs(jobParams, token)
			if err != nil {
				t.Fatalf("Error getting the job, Error: %v", err)
			}

			if isMfaStepInJob(jobsResponse.Payload.Steps) {
				postJobMfaParams := jobs.PostJobMfaParams{
					JobID:   *jobsResponse.Payload.ID,
					Context: context.TODO(),
					JobPostRequest: &models.JobPostRequest{
						MfaResponse: []string{"Hooli"},
					},
				}
				tokenClientScope := httptransport.BearerToken(test.GetTokenWithClientScope(t))

				postJobMfaResponse, err := test.Client.Jobs.PostJobMfa(&postJobMfaParams, tokenClientScope)
				if err != nil {
					t.Fatalf("Error on POST user job mfa: %v", err)
				}

				assert.Equal(t, jobsResponse.Payload.ID, postJobMfaResponse.Payload.ID, "Should be the same job ID")

				s := test.GetJsonResponse("./responses/postMfa.json", t)
				s = strings.Replace(s, "25689822-63db-4e51-bb67-6091c2d4f2e9", *postJobMfaResponse.Payload.ID, 2)
				s = strings.Replace(s, "{self}", *postJobMfaResponse.Payload.Links.Self, 1)

				e, err := json.Marshal(postJobMfaResponse.Payload)
				if err != nil {
					t.Fatalf("Error: %v", err)
				}
				test.AssertJson(t, s, string(e))

				jobRsp, err := test.Client.Jobs.GetJobs(jobParams, token)
				if err != nil {
					t.Fatalf("Error getting job: %v", err)
				}

				if isMfaStepSuccess(jobRsp.Payload.Steps) {
					allStepsSuccess = true
					ticker.Stop()
					end.Stop()
					break
				}
			}
			js, _ := json.Marshal(jobsResponse.Payload.Steps)
			t.Fatalf("No mfa-challenge found %s", string(js))
		}
	}
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

func validateSteps(steps []*models.JobsStep) bool {
	for _, s := range steps {
		if *s.Status != models.JobsStepStatusSuccess {
			return false
		}
	}
	return true
}

func isMfaStepSuccess(steps []*models.JobsStep) bool {
	for _, s := range steps {
		if s.Title == models.JobsStepTitleMfaDashChallenge && *s.Status == models.JobsStepStatusSuccess {
			return true
		}
	}
	return false
}

func isMfaStepInJob(steps []*models.JobsStep) bool {
	for _, s := range steps {
		if s.Title == models.JobsStepTitleMfaDashChallenge {
			return true
		}
	}
	return false
}

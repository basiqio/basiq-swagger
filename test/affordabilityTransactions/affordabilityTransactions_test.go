package affordabilityTransactions

import (
	"context"
	"github.com/basiqio/basiq-swagger/dist/client/affordability"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"testing"
)

func TestGetAffordabilityTransactions(t *testing.T) {
	_ = httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	snapshotID := "91b42a5b-2f9a-479d-9014-cdc211155f74"

	_ = &affordability.GetAffordabilitySnapshotTransactionsParams{
		UserID:     userID,
		SnapshotID: snapshotID,
		Context:    context.TODO(),
	}

}

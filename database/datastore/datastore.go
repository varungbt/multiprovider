package datastore

import (
	"context"
	"fmt"
	"multiprovider/database"
	"multiprovider/database/driver"
	"os"

	"cloud.google.com/go/datastore"
)

// Provider is google
const Provider = "google"

func init() {
	database.DefaultProviderMux().RegisterDbClient(Provider, new(datastoreConnOpener))
}

// DatastoreDbClient ...
type datastoreDbClient struct {
	kind   string
	client *datastore.Client
}

// DatastoreConnOpener ...
type datastoreConnOpener struct {
	client *datastore.Client
}

func (co *datastoreConnOpener) OpenConnection(ctx context.Context) (*database.DbClient, error) {

	project := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := datastore.NewClient(ctx, project)

	if err != nil {
		return nil, err
	}

	datastoreClient := &datastoreDbClient{
		kind:   "converter-jobs",
		client: client,
	}

	return database.NewDbClient(datastoreClient), nil

}

// RegisterJob registers a job in google cloud datastore
func (jd *datastoreDbClient) RegisterJob(jobID string, jobType string) bool {
	fmt.Println("Registering a job in Google Cloud datastore ")
	return true
}

// PutJob ...
func (jd *datastoreDbClient) PutJob(jobID string) bool {
	return true
}

// DeleteJob ...
func (jd *datastoreDbClient) DeleteJob(jobID string) bool {
	return true
}

// GetJob ...
func (jd *datastoreDbClient) GetJob(jobID string) driver.Job {
	return driver.Job{
		JobID: "jobid",
		Type:  "conversionJob",
	}
}

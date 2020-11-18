package datastore

import (
	"multiprovider/database/driver"

	"cloud.google.com/go/datastore"
)

// GcpJobsDatabase ...
type GcpJobsDatabase struct {
	kind   string
	client *datastore.Client
}

// RegisterJob registers a job in google cloud datastore
func (jd *GcpJobsDatabase) RegisterJob(jobID string, jobType string) bool {
	return true
}

// PutJob ...
func (jd *GcpJobsDatabase) PutJob(jobID string) bool {
	return true
}

// DeleteJob ...
func (jd *GcpJobsDatabase) DeleteJob(jobID string) bool {
	return true
}

// GetJob ...
func (jd *GcpJobsDatabase) GetJob(jobID string) driver.Job {
	return driver.Job{
		JobID: "jobid",
		Type:  "conversionJob",
	}
}

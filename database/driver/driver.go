package driver

// Job ...
type Job struct {
	JobID string
	Type  string
}

// DbClient  is a interface to be implemented by every cloud provider
type DbClient interface {
	RegisterJob(jobID string, jobType string) bool

	PutJob(jobID string) bool

	DeleteJob(jobID string) bool

	GetJob(jobID string) Job
}

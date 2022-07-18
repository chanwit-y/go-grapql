// memory is a in memory data storage solution for Job
package job

import (
	"errors"
	"sync"
)

// InMemoryRepository is a storage for jobs that uses a map to store them
type InMemoryRepository struct {
	// jobs is used to store jobs
	jobs map[string][]Job
	sync.Mutex
}

// NewMemoryRepository initializes a memory with mock data
func NewMemoryRepository() *InMemoryRepository {
	jobs := make(map[string][]Job)

	jobs["1"] = []Job{
		{
			ID:         "123-123",
			EmployeeID: "1",
			Company:    "Google",
			Title:      "Logo",
			Start:      "2021-01-01",
			End:        "",
		},
	}
	jobs["2"] = []Job{
		{
			ID:         "124-124",
			EmployeeID: "2",
			Company:    "Google",
			Title:      "Janitor",
			Start:      "2021-05-03",
			End:        "",
		}, {
			ID:         "125-125",
			EmployeeID: "2",
			Company:    "Microsoft",
			Title:      "Janitor",
			Start:      "1980-03-04",
			End:        "2021-05-02",
		},
	}
	return &InMemoryRepository{
		jobs: jobs,
	}
}

// GetJobs returns all jobs for a certain Employee
func (imr *InMemoryRepository) GetJobs(employeeID string) ([]Job, error) {

	if jobs, ok := imr.jobs[employeeID]; ok {
		return jobs, nil
	}
	return nil, errors.New("no such employee exist")

}

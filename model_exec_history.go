/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type ExecHistory struct {
	// the status of purge job.
	JobStatus string `json:"job_status,omitempty"`
	// the update time of purge job.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// the job parameters of purge job.
	JobParameters string       `json:"job_parameters,omitempty"`
	Schedule      *ScheduleObj `json:"schedule,omitempty"`
	// if purge job was deleted.
	Deleted bool `json:"deleted,omitempty"`
	// the job kind of purge job.
	JobKind string `json:"job_kind,omitempty"`
	// the creation time of purge job.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// the id of purge job.
	Id int32 `json:"id,omitempty"`
	// the job name of purge job.
	JobName string `json:"job_name,omitempty"`
}

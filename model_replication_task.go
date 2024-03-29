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

// The replication task
type ReplicationTask struct {
	// The status of the task
	Status string `json:"status,omitempty"`
	// The ID of the underlying job that the task related to
	JobId string `json:"job_id,omitempty"`
	// The start time of the task
	StartTime time.Time `json:"start_time,omitempty"`
	// The destination resource that the task operates
	DstResource string `json:"dst_resource,omitempty"`
	// The source resource that the task operates
	SrcResource string `json:"src_resource,omitempty"`
	// The type of the resource that the task operates
	ResourceType string `json:"resource_type,omitempty"`
	// The operation of the task
	Operation string `json:"operation,omitempty"`
	// The ID of the task
	Id int32 `json:"id,omitempty"`
	// The ID of the execution that the task belongs to
	ExecutionId int32 `json:"execution_id,omitempty"`
	// The end time of the task
	EndTime time.Time `json:"end_time,omitempty"`
}

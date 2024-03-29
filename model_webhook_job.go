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

// The webhook job.
type WebhookJob struct {
	// The webhook job status.
	Status string `json:"status,omitempty"`
	// The webhook job update time.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The webhook job event type.
	EventType string `json:"event_type,omitempty"`
	// The webhook job creation time.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The webhook job notify detailed data.
	JobDetail string `json:"job_detail,omitempty"`
	// The webhook job ID.
	Id int64 `json:"id,omitempty"`
	// The webhook job notify type.
	NotifyType string `json:"notify_type,omitempty"`
	// The webhook policy ID.
	PolicyId int64 `json:"policy_id,omitempty"`
}

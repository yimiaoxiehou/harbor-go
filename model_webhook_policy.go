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

// The webhook policy object
type WebhookPolicy struct {
	// The update time of the webhook policy.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The description of webhook policy.
	Description string `json:"description,omitempty"`
	// The creator of the webhook policy.
	Creator string `json:"creator,omitempty"`
	// The create time of the webhook policy.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// Whether the webhook policy is enabled or not.
	Enabled    bool                  `json:"enabled,omitempty"`
	Targets    []WebhookTargetObject `json:"targets,omitempty"`
	EventTypes []string              `json:"event_types,omitempty"`
	// The project ID of webhook policy.
	ProjectId int32 `json:"project_id,omitempty"`
	// The webhook policy ID.
	Id int64 `json:"id,omitempty"`
	// The name of webhook policy.
	Name string `json:"name,omitempty"`
}

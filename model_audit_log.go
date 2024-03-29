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

type AuditLog struct {
	// Username of the user in this log entry.
	Username string `json:"username,omitempty"`
	// Name of the repository in this log entry.
	Resource string `json:"resource,omitempty"`
	// The operation against the repository in this log entry.
	Operation string `json:"operation,omitempty"`
	// The time when this operation is triggered.
	OpTime time.Time `json:"op_time,omitempty"`
	// The ID of the audit log entry.
	Id int32 `json:"id,omitempty"`
	// Tag of the repository in this log entry.
	ResourceType string `json:"resource_type,omitempty"`
}

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

type Schedule struct {
	// The status of the schedule.
	Status string `json:"status,omitempty"`
	// the update time of the schedule.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The parameters of schedule job
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Schedule   *ScheduleObj           `json:"schedule,omitempty"`
	// the creation time of the schedule.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The id of the schedule.
	Id int32 `json:"id,omitempty"`
}

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

// the schedule task info
type ScheduleTask struct {
	// the cron of the current schedule task
	Cron string `json:"cron,omitempty"`
	// the update time of the schedule task
	UpdateTime time.Time `json:"update_time,omitempty"`
	// the vendor id of the current task
	VendorId int32 `json:"vendor_id,omitempty"`
	// the vendor type of the current schedule task
	VendorType string `json:"vendor_type,omitempty"`
	// the id of the Schedule task
	Id int32 `json:"id,omitempty"`
}
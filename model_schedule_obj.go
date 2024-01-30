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

type ScheduleObj struct {
	// A cron expression, a time-based job scheduler.
	Cron string `json:"cron,omitempty"`
	// The next time to schedule to run the job.
	NextScheduledTime time.Time `json:"next_scheduled_time,omitempty"`
	// The schedule type. The valid values are 'Hourly', 'Daily', 'Weekly', 'Custom', 'Manual' and 'None'. 'Manual' means to trigger it right away and 'None' means to cancel the schedule.
	Type_ string `json:"type,omitempty"`
}
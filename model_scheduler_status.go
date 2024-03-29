/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// the scheduler status
type SchedulerStatus struct {
	// if the scheduler is paused
	Paused bool `json:"paused,omitempty"`
}

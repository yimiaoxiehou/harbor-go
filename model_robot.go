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

type Robot struct {
	// The update time of the robot.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The description of the robot
	Description string `json:"description,omitempty"`
	// The level of the robot, project or system
	Level string `json:"level,omitempty"`
	// The editable status of the robot
	Editable bool `json:"editable,omitempty"`
	// The creation time of the robot.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The expiration date of the robot
	ExpiresAt int64 `json:"expires_at,omitempty"`
	// The name of the robot
	Name string `json:"name,omitempty"`
	// The secret of the robot
	Secret string `json:"secret,omitempty"`
	// The disable status of the robot
	Disable bool `json:"disable,omitempty"`
	// The duration of the robot in days
	Duration int64 `json:"duration,omitempty"`
	// The ID of the robot
	Id          int64             `json:"id,omitempty"`
	Permissions []RobotPermission `json:"permissions,omitempty"`
}

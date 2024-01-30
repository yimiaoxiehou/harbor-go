/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The request for robot account creation.
type RobotCreate struct {
	// The secret of the robot
	Secret string `json:"secret,omitempty"`
	// The disable status of the robot
	Disable bool `json:"disable,omitempty"`
	// The name of the robot
	Name string `json:"name,omitempty"`
	// The level of the robot, project or system
	Level string `json:"level,omitempty"`
	// The duration of the robot in days
	Duration int64 `json:"duration,omitempty"`
	// The description of the robot
	Description string            `json:"description,omitempty"`
	Permissions []RobotPermission `json:"permissions,omitempty"`
}

/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StringConfigItem struct {
	// The configure item can be updated or not
	Editable bool `json:"editable,omitempty"`
	// The string value of current config item
	Value string `json:"value,omitempty"`
}
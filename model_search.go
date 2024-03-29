/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Search struct {
	// Search results of the projects that matched the filter keywords.
	Project []Project `json:"project,omitempty"`
	// Search results of the repositories that matched the filter keywords.
	Repository []SearchRepository `json:"repository,omitempty"`
}

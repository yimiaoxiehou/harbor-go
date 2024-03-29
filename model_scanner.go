/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Scanner struct {
	// Version of the scanner adapter
	Version string `json:"version,omitempty"`
	// Name of the scanner provider
	Vendor string `json:"vendor,omitempty"`
	// Name of the scanner
	Name string `json:"name,omitempty"`
}

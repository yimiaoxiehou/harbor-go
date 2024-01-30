/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Icon struct {
	// The base64 encoded content of the icon
	Content string `json:"content,omitempty"`
	// The content type of the icon
	ContentType string `json:"content-type,omitempty"`
}
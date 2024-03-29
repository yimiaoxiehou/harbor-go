/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Endpoint struct {
	// The URL of OIDC endpoint to be tested.
	Url string `json:"url,omitempty"`
	// Whether the certificate should be verified
	VerifyCert bool `json:"verify_cert,omitempty"`
}

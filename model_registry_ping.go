/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RegistryPing struct {
	// The registry access key.
	AccessKey string `json:"access_key,omitempty"`
	// Credential type of the registry, e.g. 'basic'.
	CredentialType string `json:"credential_type,omitempty"`
	// The registry access secret.
	AccessSecret string `json:"access_secret,omitempty"`
	// The registry URL.
	Url string `json:"url,omitempty"`
	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`
	// Type of the registry, e.g. 'harbor'.
	Type_ string `json:"type,omitempty"`
	// The registry ID.
	Id int64 `json:"id,omitempty"`
}

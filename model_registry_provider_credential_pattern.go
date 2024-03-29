/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The registry credential pattern
type RegistryProviderCredentialPattern struct {
	// The access key type
	AccessKeyType string `json:"access_key_type,omitempty"`
	// The access secret data
	AccessSecretData string `json:"access_secret_data,omitempty"`
	// The access secret type
	AccessSecretType string `json:"access_secret_type,omitempty"`
	// The access key data
	AccessKeyData string `json:"access_key_data,omitempty"`
}

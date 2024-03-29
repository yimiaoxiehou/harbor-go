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

type GeneralInfo struct {
	// The flag to indicate whether Harbor is in readonly mode.
	ReadOnly bool `json:"read_only,omitempty"`
	// The current time of the server.
	CurrentTime time.Time `json:"current_time,omitempty"`
	// The setting of auth proxy this is only available when Harbor relies on authproxy for authentication.
	AuthproxySettings *AuthproxySetting `json:"authproxy_settings,omitempty"`
	// The build version of Harbor.
	HarborVersion string `json:"harbor_version,omitempty"`
	// The flag to indicate whether notification mechanism is enabled on Harbor instance.
	NotificationEnable bool `json:"notification_enable,omitempty"`
	// The auth mode of current Harbor instance.
	AuthMode string `json:"auth_mode,omitempty"`
	// The flag to indicate whether the current auth mode should consider as a primary one.
	PrimaryAuthMode bool `json:"primary_auth_mode,omitempty"`
	// Indicate whether the Harbor instance enable user to register himself.
	SelfRegistration bool `json:"self_registration,omitempty"`
	// The external URL of Harbor, with protocol.
	ExternalUrl string `json:"external_url,omitempty"`
	// Indicate who can create projects, it could be 'adminonly' or 'everyone'.
	ProjectCreationRestriction string `json:"project_creation_restriction,omitempty"`
	// Indicate whether there is a ca root cert file ready for download in the file system.
	HasCaRoot bool `json:"has_ca_root,omitempty"`
	// If the Harbor instance is deployed with nested notary.
	WithNotary bool `json:"with_notary,omitempty"`
	// The storage provider's name of Harbor registry
	RegistryStorageProviderName string `json:"registry_storage_provider_name,omitempty"`
	// The url of registry against which the docker command should be issued.
	RegistryUrl string `json:"registry_url,omitempty"`
}

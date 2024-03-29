/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ConfigurationsResponse struct {
	// Verify the OIDC provider's certificate'
	OidcVerifyCert *BoolConfigItem `json:"oidc_verify_cert,omitempty"`
	// Whether skip the audit log in database
	SkipAuditLogDatabase *BoolConfigItem `json:"skip_audit_log_database,omitempty"`
	// Extra parameters to add when redirect request to OIDC provider
	OidcExtraRedirectParms *StringConfigItem `json:"oidc_extra_redirect_parms,omitempty"`
	// The auth mode of current system, such as \"db_auth\", \"ldap_auth\", \"oidc_auth\"
	AuthMode *StringConfigItem `json:"auth_mode,omitempty"`
	// Whether the Harbor instance supports self-registration.  If it''s set to false, admin need to add user to the instance.
	SelfRegistration *BoolConfigItem `json:"self_registration,omitempty"`
	// The token review endpoint
	HttpAuthproxyTokenreviewEndpoint *StringConfigItem `json:"http_authproxy_tokenreview_endpoint,omitempty"`
	// The DN of the user to do the search.
	LdapSearchDn *StringConfigItem `json:"ldap_search_dn,omitempty"`
	// The storage quota per project
	StoragePerProject *IntegerConfigItem                   `json:"storage_per_project,omitempty"`
	ScanAllPolicy     *ConfigurationsResponseScanAllPolicy `json:"scan_all_policy,omitempty"`
	// Verify the HTTP auth provider's certificate
	HttpAuthproxyVerifyCert *BoolConfigItem `json:"http_authproxy_verify_cert,omitempty"`
	// The filter to search the ldap group
	LdapGroupSearchFilter *StringConfigItem `json:"ldap_group_search_filter,omitempty"`
	// The client id of UAA
	UaaClientId *StringConfigItem `json:"uaa_client_id,omitempty"`
	// Timeout in seconds for connection to LDAP server
	LdapTimeout *IntegerConfigItem `json:"ldap_timeout,omitempty"`
	// The Base DN for LDAP binding.
	LdapBaseDn *StringConfigItem `json:"ldap_base_dn,omitempty"`
	// The filter for LDAP search
	LdapFilter *StringConfigItem `json:"ldap_filter,omitempty"`
	// The flag to indicate whether Harbor is in readonly mode.
	ReadOnly *BoolConfigItem `json:"read_only,omitempty"`
	// The robot account token duration in days
	RobotTokenDuration *IntegerConfigItem `json:"robot_token_duration,omitempty"`
	// Auto onboard the OIDC user
	OidcAutoOnboard *BoolConfigItem `json:"oidc_auto_onboard,omitempty"`
	// The certificate of the HTTP auth provider
	HttpAuthproxyServerCertificate *StringConfigItem `json:"http_authproxy_server_certificate,omitempty"`
	// The OIDC provider name
	OidcName *StringConfigItem `json:"oidc_name,omitempty"`
	// Enable quota per project
	QuotaPerProjectEnable *BoolConfigItem `json:"quota_per_project_enable,omitempty"`
	// The URL of LDAP server
	LdapUrl *StringConfigItem `json:"ldap_url,omitempty"`
	// The endpoint of the audit log forwarder
	AuditLogForwardEndpoint *StringConfigItem `json:"audit_log_forward_endpoint,omitempty"`
	// Indicate who can create projects, it could be ''adminonly'' or ''everyone''.
	ProjectCreationRestriction *StringConfigItem `json:"project_creation_restriction,omitempty"`
	// The client secret of the UAA
	UaaClientSecret *StringConfigItem `json:"uaa_client_secret,omitempty"`
	// The attribute which is used as identity for the LDAP binding, such as \"CN\" or \"SAMAccountname\"
	LdapUid *StringConfigItem `json:"ldap_uid,omitempty"`
	// Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate.
	LdapVerifyCert *BoolConfigItem `json:"ldap_verify_cert,omitempty"`
	// The client ID of the OIDC provider
	OidcClientId *StringConfigItem `json:"oidc_client_id,omitempty"`
	// The base DN to search LDAP group.
	LdapGroupBaseDn *StringConfigItem `json:"ldap_group_base_dn,omitempty"`
	// The attribute which is used as identity of the LDAP group, default is cn.'
	LdapGroupAttributeName *StringConfigItem `json:"ldap_group_attribute_name,omitempty"`
	// Specify the ldap group which have the same privilege with Harbor admin
	LdapGroupAdminDn *StringConfigItem `json:"ldap_group_admin_dn,omitempty"`
	// The OIDC group filter which filters out the group doesn't match the regular expression
	OidcGroupFilter *StringConfigItem `json:"oidc_group_filter,omitempty"`
	// The usernames which has the harbor admin privileges
	HttpAuthproxyAdminUsernames *StringConfigItem `json:"http_authproxy_admin_usernames,omitempty"`
	// The flag to indicate whether the current auth mode should consider as a primary one.
	PrimaryAuthMode *BoolConfigItem `json:"primary_auth_mode,omitempty"`
	// The group which has the harbor admin privileges
	HttpAuthproxyAdminGroups *StringConfigItem `json:"http_authproxy_admin_groups,omitempty"`
	// The endpoint of the OIDC provider
	OidcEndpoint *StringConfigItem `json:"oidc_endpoint,omitempty"`
	// The endpoint of the HTTP auth
	HttpAuthproxyEndpoint *StringConfigItem `json:"http_authproxy_endpoint,omitempty"`
	// The rebot account name prefix
	RobotNamePrefix *StringConfigItem `json:"robot_name_prefix,omitempty"`
	// The OIDC group which has the harbor admin privileges
	OidcAdminGroup *StringConfigItem `json:"oidc_admin_group,omitempty"`
	// The scope to search ldap users,'0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE'
	LdapScope *IntegerConfigItem `json:"ldap_scope,omitempty"`
	// The endpoint of the UAA
	UaaEndpoint *StringConfigItem `json:"uaa_endpoint,omitempty"`
	// Search user before onboard
	HttpAuthproxySkipSearch *BoolConfigItem `json:"http_authproxy_skip_search,omitempty"`
	// Whether or not to skip update the pull time for scanner
	ScannerSkipUpdatePulltime *BoolConfigItem `json:"scanner_skip_update_pulltime,omitempty"`
	// The user attribute to identify the group membership
	LdapGroupMembershipAttribute *StringConfigItem `json:"ldap_group_membership_attribute,omitempty"`
	// The scope of the OIDC provider
	OidcScope *StringConfigItem `json:"oidc_scope,omitempty"`
	// The expiration time of the token for internal Registry, in minutes.
	TokenExpiration *IntegerConfigItem `json:"token_expiration,omitempty"`
	// Enable notification
	NotificationEnable *BoolConfigItem `json:"notification_enable,omitempty"`
	// The attribute claims the username
	OidcUserClaim *StringConfigItem `json:"oidc_user_claim,omitempty"`
	// The attribute claims the group name
	OidcGroupsClaim *StringConfigItem `json:"oidc_groups_claim,omitempty"`
	// The session timeout in minutes
	SessionTimeout *IntegerConfigItem `json:"session_timeout,omitempty"`
	// The scope to search ldap group. ''0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE''
	LdapGroupSearchScope *IntegerConfigItem `json:"ldap_group_search_scope,omitempty"`
	// Verify the certificate in UAA server
	UaaVerifyCert *BoolConfigItem `json:"uaa_verify_cert,omitempty"`
}

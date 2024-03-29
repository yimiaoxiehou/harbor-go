# Configurations

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcVerifyCert** | **bool** | Verify the OIDC provider&#39;s certificate&#39; | [optional] [default to null]
**SkipAuditLogDatabase** | **bool** | Skip audit log database | [optional] [default to null]
**OidcExtraRedirectParms** | **string** | Extra parameters to add when redirect request to OIDC provider | [optional] [default to null]
**AuthMode** | **string** | The auth mode of current system, such as \&quot;db_auth\&quot;, \&quot;ldap_auth\&quot;, \&quot;oidc_auth\&quot; | [optional] [default to null]
**SelfRegistration** | **bool** | Whether the Harbor instance supports self-registration.  If it&#39;&#39;s set to false, admin need to add user to the instance. | [optional] [default to null]
**HttpAuthproxyTokenreviewEndpoint** | **string** | The token review endpoint | [optional] [default to null]
**LdapSearchDn** | **string** | The DN of the user to do the search. | [optional] [default to null]
**StoragePerProject** | **int32** | The storage quota per project | [optional] [default to null]
**HttpAuthproxyVerifyCert** | **bool** | Verify the HTTP auth provider&#39;s certificate | [optional] [default to null]
**LdapSearchPassword** | **string** | The password of the ldap search dn | [optional] [default to null]
**LdapGroupSearchFilter** | **string** | The filter to search the ldap group | [optional] [default to null]
**UaaClientId** | **string** | The client id of UAA | [optional] [default to null]
**LdapTimeout** | **int32** | Timeout in seconds for connection to LDAP server | [optional] [default to null]
**LdapBaseDn** | **string** | The Base DN for LDAP binding. | [optional] [default to null]
**LdapFilter** | **string** | The filter for LDAP search | [optional] [default to null]
**ReadOnly** | **bool** | The flag to indicate whether Harbor is in readonly mode. | [optional] [default to null]
**RobotTokenDuration** | **int32** | The robot account token duration in days | [optional] [default to null]
**OidcAutoOnboard** | **bool** | Auto onboard the OIDC user | [optional] [default to null]
**HttpAuthproxyServerCertificate** | **string** | The certificate of the HTTP auth provider | [optional] [default to null]
**OidcName** | **string** | The OIDC provider name | [optional] [default to null]
**QuotaPerProjectEnable** | **bool** | Enable quota per project | [optional] [default to null]
**LdapUrl** | **string** | The URL of LDAP server | [optional] [default to null]
**AuditLogForwardEndpoint** | **string** | The audit log forward endpoint | [optional] [default to null]
**ProjectCreationRestriction** | **string** | Indicate who can create projects, it could be &#39;&#39;adminonly&#39;&#39; or &#39;&#39;everyone&#39;&#39;. | [optional] [default to null]
**UaaClientSecret** | **string** | The client secret of the UAA | [optional] [default to null]
**LdapUid** | **string** | The attribute which is used as identity for the LDAP binding, such as \&quot;CN\&quot; or \&quot;SAMAccountname\&quot; | [optional] [default to null]
**LdapVerifyCert** | **bool** | Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate. | [optional] [default to null]
**OidcClientId** | **string** | The client ID of the OIDC provider | [optional] [default to null]
**LdapGroupBaseDn** | **string** | The base DN to search LDAP group. | [optional] [default to null]
**LdapGroupAttributeName** | **string** | The attribute which is used as identity of the LDAP group, default is cn.&#39; | [optional] [default to null]
**LdapGroupAdminDn** | **string** | Specify the ldap group which have the same privilege with Harbor admin | [optional] [default to null]
**OidcGroupFilter** | **string** | The OIDC group filter which filters out the group name doesn&#39;t match the regular expression | [optional] [default to null]
**HttpAuthproxyAdminUsernames** | **string** | The username which has the harbor admin privileges | [optional] [default to null]
**PrimaryAuthMode** | **bool** | The flag to indicate whether the current auth mode should consider as a primary one. | [optional] [default to null]
**HttpAuthproxyAdminGroups** | **string** | The group which has the harbor admin privileges | [optional] [default to null]
**OidcEndpoint** | **string** | The endpoint of the OIDC provider | [optional] [default to null]
**HttpAuthproxyEndpoint** | **string** | The endpoint of the HTTP auth | [optional] [default to null]
**OidcClientSecret** | **string** | The OIDC provider secret | [optional] [default to null]
**OidcAdminGroup** | **string** | The OIDC group which has the harbor admin privileges | [optional] [default to null]
**LdapScope** | **int32** | The scope to search ldap users,&#39;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#39; | [optional] [default to null]
**UaaEndpoint** | **string** | The endpoint of the UAA | [optional] [default to null]
**HttpAuthproxySkipSearch** | **bool** | Search user before onboard | [optional] [default to null]
**ScannerSkipUpdatePulltime** | **bool** | Whether or not to skip update pull time for scanner | [optional] [default to null]
**LdapGroupMembershipAttribute** | **string** | The user attribute to identify the group membership | [optional] [default to null]
**OidcScope** | **string** | The scope of the OIDC provider | [optional] [default to null]
**TokenExpiration** | **int32** | The expiration time of the token for internal Registry, in minutes. | [optional] [default to null]
**NotificationEnable** | **bool** | Enable notification | [optional] [default to null]
**OidcUserClaim** | **string** | The attribute claims the username | [optional] [default to null]
**OidcGroupsClaim** | **string** | The attribute claims the group name | [optional] [default to null]
**SessionTimeout** | **int32** | The session timeout for harbor, in minutes. | [optional] [default to null]
**LdapGroupSearchScope** | **int32** | The scope to search ldap group. &#39;&#39;0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE&#39;&#39; | [optional] [default to null]
**RobotNamePrefix** | **string** | The rebot account name prefix | [optional] [default to null]
**UaaVerifyCert** | **bool** | Verify the certificate in UAA server | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



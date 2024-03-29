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

type UserResp struct {
	Comment  string `json:"comment,omitempty"`
	Username string `json:"username,omitempty"`
	// The update time of the user.
	UpdateTime   time.Time     `json:"update_time,omitempty"`
	UserId       int32         `json:"user_id,omitempty"`
	Realname     string        `json:"realname,omitempty"`
	OidcUserMeta *OidcUserInfo `json:"oidc_user_meta,omitempty"`
	// The creation time of the user.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// indicate the admin privilege is grant by authenticator (LDAP), is always false unless it is the current login user
	AdminRoleInAuth bool   `json:"admin_role_in_auth,omitempty"`
	SysadminFlag    bool   `json:"sysadmin_flag,omitempty"`
	Email           string `json:"email,omitempty"`
}

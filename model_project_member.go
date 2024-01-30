/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProjectMember struct {
	// The role id 1 for projectAdmin, 2 for developer, 3 for guest, 4 for maintainer
	RoleId      int32       `json:"role_id,omitempty"`
	MemberGroup *UserGroup  `json:"member_group,omitempty"`
	MemberUser  *UserEntity `json:"member_user,omitempty"`
}
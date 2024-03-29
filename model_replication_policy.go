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

type ReplicationPolicy struct {
	// The update time of the policy.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The description of the policy.
	Description string `json:"description,omitempty"`
	// Specify how many path components will be replaced by the provided destination namespace. The default value is -1 in which case the legacy mode will be applied.
	DestNamespaceReplaceCount int32 `json:"dest_namespace_replace_count,omitempty"`
	// The create time of the policy.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The source registry.
	SrcRegistry *Registry `json:"src_registry,omitempty"`
	// Whether to replicate the deletion operation.
	ReplicateDeletion bool `json:"replicate_deletion,omitempty"`
	// The replication policy filter array.
	Filters []ReplicationFilter `json:"filters,omitempty"`
	// speed limit for each task
	Speed int32 `json:"speed,omitempty"`
	// The policy ID.
	Id int64 `json:"id,omitempty"`
	// Whether to enable copy by chunk.
	CopyByChunk bool `json:"copy_by_chunk,omitempty"`
	// The policy name.
	Name string `json:"name,omitempty"`
	// The destination registry.
	DestRegistry *Registry `json:"dest_registry,omitempty"`
	// Whether the policy is enabled or not.
	Enabled bool `json:"enabled,omitempty"`
	// The destination namespace.
	DestNamespace string              `json:"dest_namespace,omitempty"`
	Trigger       *ReplicationTrigger `json:"trigger,omitempty"`
	// Deprecated, use \"replicate_deletion\" instead. Whether to replicate the deletion operation.
	Deletion bool `json:"deletion,omitempty"`
	// Whether to override the resources on the destination registry.
	Override bool `json:"override,omitempty"`
}

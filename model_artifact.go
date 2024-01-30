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

type Artifact struct {
	// The size of the artifact
	Size int64 `json:"size,omitempty"`
	// The push time of the artifact
	PushTime time.Time `json:"push_time,omitempty"`
	// The overview of the scan result.
	ScanOverview *ScanOverview `json:"scan_overview,omitempty"`
	Tags         []Tag         `json:"tags,omitempty"`
	// The latest pull time of the artifact
	PullTime    time.Time   `json:"pull_time,omitempty"`
	Labels      []Label     `json:"labels,omitempty"`
	Accessories []Accessory `json:"accessories,omitempty"`
	References  []Reference `json:"references,omitempty"`
	// The manifest media type of the artifact
	ManifestMediaType string      `json:"manifest_media_type,omitempty"`
	ExtraAttrs        *ExtraAttrs `json:"extra_attrs,omitempty"`
	// The ID of the artifact
	Id int64 `json:"id,omitempty"`
	// The digest of the artifact
	Digest string `json:"digest,omitempty"`
	// The digest of the icon
	Icon string `json:"icon,omitempty"`
	// The ID of the repository that the artifact belongs to
	RepositoryId  int64          `json:"repository_id,omitempty"`
	AdditionLinks *AdditionLinks `json:"addition_links,omitempty"`
	// The media type of the artifact
	MediaType string `json:"media_type,omitempty"`
	// The ID of the project that the artifact belongs to
	ProjectId int64 `json:"project_id,omitempty"`
	// The type of the artifact, e.g. image, chart, etc
	Type_       string       `json:"type,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
}
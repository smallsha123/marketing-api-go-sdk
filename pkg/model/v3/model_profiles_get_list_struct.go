/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ProfilesGetListStruct struct {
	OwnerId          *int64      `json:"owner_id,omitempty"`
	ProfileType      ProfileType `json:"profile_type,omitempty"`
	ProfileId        *int64      `json:"profile_id,omitempty"`
	HeadImageId      *string     `json:"head_image_id,omitempty"`
	HeadImageUrl     *string     `json:"head_image_url,omitempty"`
	ProfileName      *string     `json:"profile_name,omitempty"`
	Description      *string     `json:"description,omitempty"`
	CreatedTime      *int64      `json:"created_time,omitempty"`
	LastModifiedTime *int64      `json:"last_modified_time,omitempty"`
	ProfileUrl       *string     `json:"profile_url,omitempty"`
	SystemStatus     SysStatus   `json:"system_status,omitempty"`
}

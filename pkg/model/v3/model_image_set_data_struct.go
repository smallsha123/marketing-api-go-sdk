/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 门店图片 ID
type ImageSetDataStruct struct {
	ImageId  *string   `json:"image_id,omitempty"`
	Status   SysStatus `json:"status,omitempty"`
	AuditMsg *string   `json:"audit_msg,omitempty"`
}

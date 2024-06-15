/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 元素粒度审核结果
type ElementResultListStruct struct {
	ImageId                 *string                                       `json:"image_id,omitempty"`
	VideoId                 *string                                       `json:"video_id,omitempty"`
	ElementName             *string                                       `json:"element_name,omitempty"`
	ElementValue            *string                                       `json:"element_value,omitempty"`
	ElementFingerprint      *string                                       `json:"element_fingerprint,omitempty"`
	ComponentInfo           *ComponentInfoCanEmpty                        `json:"component_info,omitempty"`
	ElementType             ReviewElementType                             `json:"element_type,omitempty"`
	ReviewStatus            ReviewResultStatus                            `json:"review_status,omitempty"`
	ElementRejectDetailInfo *[]ComponentElementRejectDetailInfoListStruct `json:"element_reject_detail_info,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 素材标签信息
type MaterialLabelStruct struct {
	ImageId   *string              `json:"image_id,omitempty"`
	MediaId   *string              `json:"media_id,omitempty"`
	LabelList *[]CustomLabelStruct `json:"label_list,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 图集
type ImageListStruct struct {
	JumpInfo *JumpinfoStruct  `json:"jump_info,omitempty"`
	List     *[]ImageListItem `json:"list,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 选择按钮结构
type ChosenBaseButtonStruct struct {
	Text     *string         `json:"text,omitempty"`
	JumpInfo *JumpinfoStruct `json:"jump_info,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 咨询组件
type ConsultComponentValueStruct struct {
	Id       *int64                  `json:"id,omitempty"`
	JumpInfo *[]LandingPageStructure `json:"jump_info,omitempty"`
}
/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 获取定向标签的条件
type TagSpec struct {
	BehaviorSpec *BehaviorTargetingTagSpec `json:"behavior_spec,omitempty"`
	InterestSpec *InterestTargetingTagSpec `json:"interest_spec,omitempty"`
}

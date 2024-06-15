/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ObjectCommentFlagUpdateRequest struct {
	AccountId    *int64                     `json:"account_id,omitempty"`
	AdcreativeId *int64                     `json:"adcreative_id,omitempty"`
	OpType       SetObjectCommentFlagOpType `json:"op_type,omitempty"`
	CommentId    *string                    `json:"comment_id,omitempty"`
	CommentLevel *int64                     `json:"comment_level,omitempty"`
}

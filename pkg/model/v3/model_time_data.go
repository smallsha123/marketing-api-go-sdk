/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// TimeData返回结构
type TimeData struct {
	Date                  *string    `json:"date,omitempty"`
	Hour                  *int64     `json:"hour,omitempty"`
	Domain                *string    `json:"domain,omitempty"`
	ActionType            ActionType `json:"action_type,omitempty"`
	CustomAction          *string    `json:"custom_action,omitempty"`
	RawActionCount        *int64     `json:"raw_action_count,omitempty"`
	IdentifiedActionCount *int64     `json:"identified_action_count,omitempty"`
	IdentifiedUserCount   *int64     `json:"identified_user_count,omitempty"`
}

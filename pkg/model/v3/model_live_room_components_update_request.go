/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LiveRoomComponentsUpdateRequest struct {
	AccountId     *int64                       `json:"account_id,omitempty"`
	ComponentId   *int64                       `json:"component_id,omitempty"`
	ComponentType CreativeComponentType        `json:"component_type,omitempty"`
	ComponentSpec *LiveRoomComponentSpecStruct `json:"component_spec,omitempty"`
}

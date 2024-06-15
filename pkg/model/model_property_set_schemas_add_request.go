/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PropertySetSchemasAddRequest struct {
	AccountId     *int64                `json:"account_id,omitempty"`
	PropertySetId *int64                `json:"property_set_id,omitempty"`
	UserIdType    PropertySetUserIdType `json:"user_id_type,omitempty"`
	Schemas       *[]Schema             `json:"schemas,omitempty"`
}

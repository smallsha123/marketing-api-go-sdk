/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProgrammedUpdateRequest struct {
	AccountId                        *int64                       `json:"account_id,omitempty"`
	MaterialDeriveId                 *int64                       `json:"material_derive_id,omitempty"`
	AutoDerivedProgramCreativeSwitch *bool                        `json:"auto_derived_program_creative_switch,omitempty"`
	StandardSwitch                   *bool                        `json:"standard_switch,omitempty"`
	UpdateMaterialGroups             *[]MaterialGroupUpdateStruct `json:"update_material_groups,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CampaignsUpdateConfiguredStatusRequest struct {
	AccountId                  *int64                                                         `json:"account_id,omitempty"`
	UpdateConfiguredStatusSpec *[]CampaignsUpdateConfiguredStatusUpdateConfiguredStatusStruct `json:"update_configured_status_spec,omitempty"`
}

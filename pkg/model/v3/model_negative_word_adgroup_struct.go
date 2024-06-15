/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告否定词
type NegativeWordAdgroupStruct struct {
	AdgroupId           *int64    `json:"adgroup_id,omitempty"`
	PhraseNegativeWords *[]string `json:"phrase_negative_words,omitempty"`
	ExactNegativeWords  *[]string `json:"exact_negative_words,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 单个否定词长度超过限制导致失败的否定词列表
type ExceedLengthNegativeWordStruct struct {
	PhraseNegativeWords *[]string `json:"phrase_negative_words,omitempty"`
	ExactNegativeWords  *[]string `json:"exact_negative_words,omitempty"`
}

/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type KeywordRecommendGetResponseData struct {
	TotalNum          *int64                  `json:"total_num,omitempty"`
	RecommendWordList *[]RecommendWordStructs `json:"recommend_word_list,omitempty"`
}

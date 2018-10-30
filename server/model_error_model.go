/*
 * Pseudo Service
 *
 * Deterministic or pseudo-random data generator
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"encoding/json"
	"fmt"
)

const (
	_ = iota
	ERROR_APIKEY
	ERROR_MISSING_COUNT
	ERROR_GENERATING_USERS
	ERROR_WRONG_COUNT
)

type ErrorModel struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewErrorJson(code int32, message string) []byte {
	asStruct := ErrorModel{code, message}

	asByte, err := json.Marshal(asStruct)
	if err != nil {
		fmt.Printf("error marshalling error json %s", message)
	}
	return asByte
}

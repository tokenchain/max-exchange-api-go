/*
 * MAX RESTful API List
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// get details of a specific deposit
type Deposit struct {

	// currency id
	Currency string `json:"currency,omitempty"`

	// deposit amount
	Amount string `json:"amount,omitempty"`

	// deposit fee
	Fee string `json:"fee,omitempty"`

	// unique transaction id
	Txid string `json:"txid,omitempty"`

	// received timestamp (second)
	CreatedAt int32 `json:"created_at,omitempty"`

	// confirmations for crypto currency
	Confirmations int32 `json:"confirmations,omitempty"`

	// lastest updated timestamp (second)
	UpdatedAt int32 `json:"updated_at,omitempty"`

	// current state
	State string `json:"state,omitempty"`
}

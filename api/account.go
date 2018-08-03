/*
 * MAX RESTful API List
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type Account struct {

	// currency id, e.g. twd, btc, ...
	Currency string `json:"currency,omitempty"`

	// available balance
	Balance string `json:"balance,omitempty"`

	// locked funds
	Locked string `json:"locked,omitempty"`
}

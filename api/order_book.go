/*
 * MAX RESTful API List
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// get the order book of a specified market
type OrderBook struct {

	Asks []Order `json:"asks,omitempty"`

	Bids []Order `json:"bids,omitempty"`
}

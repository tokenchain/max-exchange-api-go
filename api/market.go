/*
 * MAX RESTful API List
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// get all available markets.
type Market struct {

	// unique market id, check /api/v2/markets for available markets
	Id string `json:"id,omitempty"`

	// market name
	Name string `json:"name,omitempty"`

	// base unit
	BaseUnit string `json:"base_unit,omitempty"`

	// fixed precision of base unit
	BaseUnitPrecision int32 `json:"base_unit_precision,omitempty"`

	// quote unit
	QuoteUnit string `json:"quote_unit,omitempty"`

	// fixed precision of quote unit
	QuoteUnitPrecision int32 `json:"quote_unit_precision,omitempty"`
}

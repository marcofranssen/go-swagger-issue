package api

import "github.com/marcofranssen/go-swagger-issue/domain"

// ExampleResponse the response of an example resource
// swagger:model ExampleResponse
type ExampleResponse struct {
	*domain.Example
}

// ExampleListResponse the response of a list of example resources
// swagger:model ExampleListResponse
type ExampleListResponse struct {
	Examples []ExampleResponse `json:"examples"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

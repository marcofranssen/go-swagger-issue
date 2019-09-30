// Package api Exposes this swagger-issue-example.
//
// the purpose of this api is to showcase the swagger doc generation issue.
//
//     Schemes: http
//     Host: localhost:5000
//     BasePath: /api/v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package api

import "net/http"

//go:generate swagger generate spec -m -o ../docs/swagger.json

// New creates a new instance of the http.ServeMux
func New() (*http.ServeMux, error) {
	router := http.NewServeMux()
	// register http.Handlers

	// swagger:operation GET /examples retrieves examples
	//
	// Get a collection of examples including paging
	//
	// This will return you a page of examples.
	//
	// ---
	// consumes:
	// - application/json
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: collection of paginated examples
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/ExampleListResponse"

	return router, nil
}

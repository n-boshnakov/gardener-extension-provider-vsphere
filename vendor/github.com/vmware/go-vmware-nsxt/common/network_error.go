/*
 * NSX-T Manager API
 *
 * VMware NSX-T Manager REST API
 *
 * API version: 3.0.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package common

// Network error related to container objects.
type NetworkError struct {
	// Error code of network related error.
	ErrorCode string `json:"error_code,omitempty"`
	// Detailed message of network related error.
	ErrorMessage string `json:"error_message,omitempty"`
	// Additional error information in json format.
	Spec string `json:"spec,omitempty"`
}

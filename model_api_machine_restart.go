/*
 * Fly Machines API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApiMachineRestart struct {
	// MaxRetries is only relevant with the on-failure policy.
	MaxRetries int32 `json:"max_retries,omitempty"`
	Policy string `json:"policy,omitempty"`
}
/*
 * Fly Machines API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MachineVersion struct {
	UserConfig *ApiMachineConfig `json:"user_config,omitempty"`
	Version string `json:"version,omitempty"`
}

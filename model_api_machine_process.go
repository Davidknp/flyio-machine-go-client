/*
 * Fly Machines API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApiMachineProcess struct {
	Cmd []string `json:"cmd,omitempty"`
	Entrypoint []string `json:"entrypoint,omitempty"`
	Env map[string]string `json:"env,omitempty"`
	Exec []string `json:"exec,omitempty"`
	User string `json:"user,omitempty"`
}

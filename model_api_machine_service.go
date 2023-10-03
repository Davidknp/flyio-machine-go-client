/*
 * Fly Machines API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApiMachineService struct {
	Autostart bool `json:"autostart,omitempty"`
	Autostop bool `json:"autostop,omitempty"`
	Checks []ApiMachineCheck `json:"checks,omitempty"`
	Concurrency *ApiMachineServiceConcurrency `json:"concurrency,omitempty"`
	ForceInstanceDescription string `json:"force_instance_description,omitempty"`
	ForceInstanceKey string `json:"force_instance_key,omitempty"`
	InternalPort int32 `json:"internal_port,omitempty"`
	MinMachinesRunning int32 `json:"min_machines_running,omitempty"`
	Ports []ApiMachinePort `json:"ports,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

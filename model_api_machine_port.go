/*
 * Fly Machines API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApiMachinePort struct {
	EndPort int32 `json:"end_port,omitempty"`
	ForceHttps bool `json:"force_https,omitempty"`
	Handlers []string `json:"handlers,omitempty"`
	HttpOptions *ApiHttpOptions `json:"http_options,omitempty"`
	Port int32 `json:"port,omitempty"`
	ProxyProtoOptions *ApiProxyProtoOptions `json:"proxy_proto_options,omitempty"`
	StartPort int32 `json:"start_port,omitempty"`
	TlsOptions *ApiTlsOptions `json:"tls_options,omitempty"`
}

/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AuthproxySetting struct {
	// The certificate to be pinned when connecting auth proxy.
	ServerCertificate string `json:"server_certificate,omitempty"`
	// The fully qualified URI of token review endpoint of authproxy, such as 'https://192.168.1.2:8443/tokenreview'
	TokenreivewEndpoint string `json:"tokenreivew_endpoint,omitempty"`
	// The fully qualified URI of login endpoint of authproxy, such as 'https://192.168.1.2:8443/login'
	Endpoint string `json:"endpoint,omitempty"`
	// The flag to determine whether Harbor should verify the certificate when connecting to the auth proxy.
	VerifyCert bool `json:"verify_cert,omitempty"`
	// The flag to determine whether Harbor can skip search the user/group when adding him as a member.
	SkipSearch bool `json:"skip_search,omitempty"`
}

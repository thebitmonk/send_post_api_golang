/*
 * SendPost API
 *
 * SendPost API to send transactional emails reliably
 *
 * API version: 1.0.0
 * Contact: hello@sendx.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelsDomain struct {
	Created int64 `json:"created,omitempty"`
	Dkim *ModelsDnsRecord `json:"dkim,omitempty"`
	DkimConfig string `json:"dkimConfig,omitempty"`
	DkimVerified bool `json:"dkimVerified,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ReturnPath *ModelsDnsRecord `json:"returnPath,omitempty"`
	ReturnPathVerified bool `json:"returnPathVerified,omitempty"`
	Track *ModelsDnsRecord `json:"track,omitempty"`
	TrackVerified bool `json:"trackVerified,omitempty"`
}
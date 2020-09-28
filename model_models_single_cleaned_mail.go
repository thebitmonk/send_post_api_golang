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

type ModelsSingleCleanedMail struct {
	Err string `json:"err,omitempty"`
	MailId string `json:"mailId,omitempty"`
	Reason string `json:"reason,omitempty"`
	Valid bool `json:"valid,omitempty"`
}

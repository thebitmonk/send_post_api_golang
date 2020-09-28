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

type ModelsWMessage struct {
	AccountID int64 `json:"AccountID,omitempty"`
	Groups *interface{} `json:"Groups,omitempty"`
	IPID int64 `json:"IPID,omitempty"`
	MessageType string `json:"MessageType,omitempty"`
	MessageUUID string `json:"MessageUUID,omitempty"`
	SMTPCode int64 `json:"SMTPCode,omitempty"`
	SMTPDescription string `json:"SMTPDescription,omitempty"`
	SubAccountID int64 `json:"SubAccountID,omitempty"`
	To string `json:"To,omitempty"`
	Type_ *ModelsEventType `json:"Type,omitempty"`
}
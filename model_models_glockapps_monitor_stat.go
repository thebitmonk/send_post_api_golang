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

type ModelsGlockappsMonitorStat struct {
	Day int64 `json:"day,omitempty"`
	ListedCount int64 `json:"listedCount,omitempty"`
	ListedIn []ModelsGlockappsBlacklist `json:"listedIn,omitempty"`
	Month int64 `json:"month,omitempty"`
	SenderScore int64 `json:"senderScore,omitempty"`
	Year int64 `json:"year,omitempty"`
}
package models

import "time"

type Events struct {
	RespStatus int       `json:"respStatus" bson:"respStatus"`
	RespBody   string    `json:"respBody" bson:"respBody"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	IsTestLogs bool      `json:"isTestLogs" bson:"isTestLogs"`
}

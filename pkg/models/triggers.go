package models

import "encoding/json"

const (
	ScheduleTrigger TriggerType = "schedule"
	ApiTrigger      TriggerType = "api"
)

type TriggerType string

type Trigger struct {
	Name        string          `json:"name" bson:"name"`
	Endpoint    string          `json:"endpoint" bson:"endpoint"`
	TriggerType TriggerType     `json:"triggerType" bson:"triggerType"`
	Payload     json.RawMessage `json:"payload" bson:"payload"`
	MethodType  string          `json:"methodType" bson:"methodType"`
	Schedule    string          `json:"schedule" bson:"schedule"`
}

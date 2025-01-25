package models

const (
	ScheduleTrigger TriggerType = "schedule"
	ApiTrigger      TriggerType = "api"
)

type TriggerType string

type Event struct {
	Endpoint    string      `json:"endpoint" bson:"endpoint"`
	TriggerType TriggerType `json:"triggerType" bson:"triggerType"`
	Payload     string      `json:"payload" bson:"payload"`
	MethodType  string      `json:"methodType" bson:"methodType"`
}

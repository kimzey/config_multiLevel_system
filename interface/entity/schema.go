package model

import (
	"encoding/json"
)

type Schema struct {
	ID          uint
	Level       string
	ServiceName *string
	SchemaData  json.RawMessage
}

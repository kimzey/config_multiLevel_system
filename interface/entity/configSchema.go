package model

import (
	"encoding/json"
)

type Configuration struct {
	ID         uint
	LevelName  string
	EntityID   string
	ConfigData json.RawMessage
}

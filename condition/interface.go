package condition

import (
	"encoding/json"
	"github.com/goodluckxu-go/notice/code"
)

type Sign uint8

type Condition interface {
	Verify() bool
	SetMetadata(map[string]*code.Metadata)
	MarshalSign() Sign
	Unmarshal(list []any) error
	json.Marshaler
}

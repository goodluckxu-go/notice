package condition

import (
	"github.com/goodluckxu-go/notice/code"
)

type Condition interface {
	Verify() bool
	SetMetadata(map[string]*code.Metadata)
}

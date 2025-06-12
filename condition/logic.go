package condition

import (
	"github.com/goodluckxu-go/notice/code"
)

type Add []Condition

func (a *Add) Verify() bool {
	for _, item := range *a {
		if !item.Verify() {
			return false
		}
	}
	return true
}

func (a *Add) SetMetadata(metadata map[string]*code.Metadata) {
	for _, item := range *a {
		item.SetMetadata(metadata)
	}
}

type Or []Condition

func (o *Or) Verify() bool {
	for _, item := range *o {
		if item.Verify() {
			return true
		}
	}
	return false
}

func (o *Or) SetMetadata(metadata map[string]*code.Metadata) {
	for _, item := range *o {
		item.SetMetadata(metadata)
	}
}

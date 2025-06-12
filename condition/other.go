package condition

import (
	"github.com/goodluckxu-go/notice/code"
)

type In struct {
	Field    string
	Value    []any
	metadata map[string]*code.Metadata
}

func (i *In) Verify() bool {
	if i.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(i.metadata[i.Field], i.Value...)
	if err != nil {
		return false
	}
	for _, v := range val {
		if v == data {
			return true
		}
	}
	return false
}

func (i *In) SetMetadata(metadata map[string]*code.Metadata) {
	i.metadata = metadata
}

type NotIn struct {
	Field    string
	Value    []any
	metadata map[string]*code.Metadata
}

func (n *NotIn) Verify() bool {
	if n.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(n.metadata[n.Field], n.Value...)
	if err != nil {
		return false
	}
	for _, v := range val {
		if v == data {
			return false
		}
	}
	return true
}

func (n *NotIn) SetMetadata(metadata map[string]*code.Metadata) {
	n.metadata = metadata
}

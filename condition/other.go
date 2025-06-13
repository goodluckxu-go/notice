package condition

import (
	"encoding/json"
	"errors"
	"github.com/goodluckxu-go/notice/code"
)

func init() {
	RegisterCondition(&In{})
	RegisterCondition(&NotIn{})
}

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

func (i *In) MarshalSign() Sign {
	return 8
}

func (i *In) MarshalJSON() ([]byte, error) {
	m := []any{i.MarshalSign(), i.Field, i.Value}
	return json.Marshal(m)
}

func (i *In) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'In' unmarshal fail")
	}
	var ok bool
	if i.Field, ok = list[1].(string); !ok {
		return errors.New("'In' unmarshal fail")
	}
	if i.Value, ok = list[2].([]any); !ok {
		return errors.New("'In' unmarshal fail")
	}
	return nil
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

func (n *NotIn) MarshalSign() Sign {
	return 9
}

func (n *NotIn) MarshalJSON() ([]byte, error) {
	m := []any{n.MarshalSign(), n.Field, n.Value}
	return json.Marshal(m)
}

func (n *NotIn) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'NotIn' unmarshal fail")
	}
	var ok bool
	if n.Field, ok = list[1].(string); !ok {
		return errors.New("'NotIn' unmarshal fail")
	}
	if n.Value, ok = list[2].([]any); !ok {
		return errors.New("'NotIn' unmarshal fail")
	}
	return nil
}

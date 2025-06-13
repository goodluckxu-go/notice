package condition

import (
	"encoding/json"
	"errors"
	"github.com/goodluckxu-go/notice/code"
)

func init() {
	RegisterCondition(&Add{})
	RegisterCondition(&Or{})
}

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

func (a *Add) MarshalSign() Sign {
	return 0
}

func (a *Add) MarshalJSON() ([]byte, error) {
	m := []any{a.MarshalSign(), *a}
	return json.Marshal(m)
}

func (a *Add) Unmarshal(list []any) error {
	if len(list) != 2 {
		return errors.New("'Add' unmarshal fail")
	}
	adds, _ := list[1].([]any)
	*a = make([]Condition, len(adds))
	var itemList []any
	var ok bool
	for i, item := range adds {
		if itemList, ok = item.([]any); !ok {
			return errors.New("'Add' unmarshal fail")
		}
		var val Condition
		if err := UnmarshalerConditionList(itemList, &val); err != nil {
			return err
		}
		(*a)[i] = val
	}
	return nil
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

func (o *Or) MarshalSign() Sign {
	return 1
}

func (o *Or) MarshalJSON() ([]byte, error) {
	m := []any{o.MarshalSign(), *o}
	return json.Marshal(m)
}

func (o *Or) Unmarshal(list []any) error {
	if len(list) != 2 {
		return errors.New("'Or' unmarshal fail")
	}
	ors, _ := list[1].([]any)
	*o = make([]Condition, len(ors))
	var itemList []any
	var ok bool
	for i, item := range ors {
		if itemList, ok = item.([]any); !ok {
			return errors.New("'Or' unmarshal fail")
		}
		var val Condition
		if err := UnmarshalerConditionList(itemList, &val); err != nil {
			return err
		}
		(*o)[i] = val
	}
	return nil
}

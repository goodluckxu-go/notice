package condition

import (
	"errors"
	"log"
	"reflect"
	"sync"
)

type conditionInfo struct {
	list  []Condition
	types []reflect.Type
	mux   sync.Mutex
}

func (c *conditionInfo) Add(condition Condition) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	if c.inArray(condition, c.list) {
		return errors.New("duplicate condition registered")
	}
	c.list = append(c.list, condition)
	c.types = append(c.types, reflect.TypeOf(condition).Elem())
	return nil
}

func (c *conditionInfo) Search(sign Sign) reflect.Type {
	c.mux.Lock()
	defer c.mux.Unlock()
	for i, v := range c.list {
		if v.MarshalSign() == sign {
			return c.types[i]
		}
	}
	return nil
}

func (c *conditionInfo) inArray(val Condition, list []Condition) bool {
	for _, v := range list {
		if v.MarshalSign() == val.MarshalSign() {
			return true
		}
	}
	return false
}

var conditions = &conditionInfo{}

// RegisterCondition 注册条件
func RegisterCondition(condition Condition) {
	if conditions.Add(condition) != nil {
		log.Fatal("duplicate condition registered")
	}
}

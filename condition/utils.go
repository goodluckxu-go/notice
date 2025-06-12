package condition

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goodluckxu-go/notice/code"
	"strings"
)

func covertMetadata(metadata *code.Metadata, values ...any) (any, []any, error) {
	if metadata == nil {
		return nil, nil, errors.New("undefined metadata")
	}
	switch data := metadata.Value.(type) {
	case *code.Metadata_Int:
		var rs []any
		for _, value := range values {
			switch val := value.(type) {
			case int:
				rs = append(rs, int64(val))
			case int8:
				rs = append(rs, int64(val))
			case int16:
				rs = append(rs, int64(val))
			case int32:
				rs = append(rs, int64(val))
			case int64:
				rs = append(rs, val)
			case uint:
				rs = append(rs, int64(val))
			case uint8:
				rs = append(rs, int64(val))
			case uint16:
				rs = append(rs, int64(val))
			case uint32:
				rs = append(rs, int64(val))
			case uint64:
				rs = append(rs, int64(val))
			case float32:
				rs = append(rs, int64(val))
			case float64:
				rs = append(rs, int64(val))
			default:
				return nil, nil, errors.New("unknown type")
			}
		}
		return data.Int, rs, nil
	case *code.Metadata_Uint:
		var rs []any
		for _, value := range values {
			switch val := value.(type) {
			case uint:
				rs = append(rs, uint64(val))
			case uint8:
				rs = append(rs, uint64(val))
			case uint16:
				rs = append(rs, uint64(val))
			case uint32:
				rs = append(rs, uint64(val))
			case uint64:
				rs = append(rs, val)
			case int:
				rs = append(rs, uint64(val))
			case int8:
				rs = append(rs, uint64(val))
			case int16:
				rs = append(rs, uint64(val))
			case int32:
				rs = append(rs, uint64(val))
			case int64:
				rs = append(rs, uint64(val))
			case float32:
				rs = append(rs, uint64(val))
			case float64:
				rs = append(rs, uint64(val))
			default:
				return nil, nil, errors.New("unknown type")
			}
		}
		return data.Uint, rs, nil
	case *code.Metadata_Float:
		var rs []any
		for _, value := range values {
			switch val := value.(type) {
			case float32:
				rs = append(rs, float64(val))
			case float64:
				rs = append(rs, val)
			case int:
				rs = append(rs, float64(val))
			case int8:
				rs = append(rs, float64(val))
			case int16:
				rs = append(rs, float64(val))
			case int32:
				rs = append(rs, float64(val))
			case int64:
				rs = append(rs, float64(val))
			case uint:
				rs = append(rs, float64(val))
			case uint8:
				rs = append(rs, float64(val))
			case uint16:
				rs = append(rs, float64(val))
			case uint32:
				rs = append(rs, float64(val))
			case uint64:
				rs = append(rs, float64(val))
			default:
				return nil, nil, errors.New("unknown type")
			}
		}
		return data.Float, rs, nil
	case *code.Metadata_String_:
		var rs []any
		for _, value := range values {
			switch val := value.(type) {
			case string:
				rs = append(rs, val)
			default:
				return nil, nil, errors.New("unknown type")
			}
		}
		return data.String_, rs, nil
	case *code.Metadata_Bool:
		var rs []any
		for _, value := range values {
			switch val := value.(type) {
			case bool:
				rs = append(rs, val)
			default:
				return nil, nil, errors.New("unknown type")
			}
		}
		return data.Bool, rs, nil
	}
	return nil, nil, errors.New("invalid metadata")
}

func covertFloat64(val any) (float64, error) {
	switch data := val.(type) {
	case float64:
		return data, nil
	case float32:
		return float64(data), nil
	case int:
		return float64(data), nil
	case int8:
		return float64(data), nil
	case int16:
		return float64(data), nil
	case int32:
		return float64(data), nil
	case int64:
		return float64(data), nil
	case uint:
		return float64(data), nil
	case uint8:
		return float64(data), nil
	case uint16:
		return float64(data), nil
	case uint32:
		return float64(data), nil
	case uint64:
		return float64(data), nil
	}
	return 0, errors.New("unsupported float64")
}

func MarshalerCondition(condition Condition) ([]byte, error) {
	if condition == nil {
		return nil, nil
	}
	switch cond := condition.(type) {
	case *Add:
		var list []string
		for _, item := range *cond {
			buf, err := MarshalerCondition(item)
			if err != nil {
				return nil, err
			}
			list = append(list, string(buf))
		}
		return []byte(fmt.Sprintf(`{"type":"add","list":[` + strings.Join(list, ",") + `]}`)), nil
	case *Or:
		var list []string
		for _, item := range *cond {
			buf, err := MarshalerCondition(item)
			if err != nil {
				return nil, err
			}
			list = append(list, string(buf))
		}
		return []byte(fmt.Sprintf(`{"type":"or","list":[` + strings.Join(list, ",") + `]}`)), nil
	case *Eq:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"eq","field":"%v","value":%v}`, cond.Field, val)), nil
	case *Neq:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"neq","field":"%v","value":%v}`, cond.Field, val)), nil
	case *Gt:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"gt","field":"%v","value":%v}`, cond.Field, val)), nil
	case *Gte:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"gte","field":"%v","value":%v}`, cond.Field, val)), nil
	case *Lt:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"lt","field":"%v","value":%v}`, cond.Field, val)), nil
	case *Lte:
		val := fmt.Sprintf("%v", cond.Value)
		if _, ok := cond.Value.(string); ok {
			val = `"` + val + `"`
		}
		return []byte(fmt.Sprintf(`{"type":"lte","field":"%v","value":%v}`, cond.Field, val)), nil
	case *In:
		var list []string
		for _, item := range cond.Value {
			val := fmt.Sprintf("%v", item)
			if _, ok := item.(string); ok {
				val = `"` + val + `"`
			}
			list = append(list, val)
		}
		return []byte(fmt.Sprintf(`{"type":"in","field":"%v","list":[`+strings.Join(list, ",")+`]}`, cond.Field)), nil
	case *NotIn:
		var list []string
		for _, item := range cond.Value {
			val := fmt.Sprintf("%v", item)
			if _, ok := item.(string); ok {
				val = `"` + val + `"`
			}
			list = append(list, val)
		}
		return []byte(fmt.Sprintf(`{"type":"not_in","field":"%v","list":[`+strings.Join(list, ",")+`]}`, cond.Field)), nil
	}
	return nil, errors.New("unsupported marshaler condition")
}

func UnmarshalerCondition(data []byte, val *Condition) error {
	if len(data) == 0 {
		return nil
	}
	m := map[string]any{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	return unmarshalerCondition(m, val)
}

func unmarshalerConditionMap(m map[string]any, val *Condition) error {
	switch m["type"] {
	case "add":
		*val = &Add{}
		return unmarshalerCondition(m["list"], val, "add")
	case "or":
		*val = &Or{}
		return unmarshalerCondition(m["list"], val, "or")
	case "eq":
		*val = &Eq{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "neq":
		*val = &Neq{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "gt":
		*val = &Gt{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "gte":
		*val = &Gte{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "lt":
		*val = &Lt{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "lte":
		*val = &Lte{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: m["value"],
		}
		return nil
	case "in":
		list, ok := m["list"].([]any)
		if !ok {
			return errors.New("invalid condition")
		}
		*val = &In{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: list,
		}
		return nil
	case "not_in":
		list, ok := m["list"].([]any)
		if !ok {
			return errors.New("invalid condition")
		}
		*val = &NotIn{
			Field: fmt.Sprintf("%v", m["field"]),
			Value: list,
		}
		return nil
	}
	return errors.New("unsupported unmarshaler condition")
}

func unmarshalerCondition(m any, val *Condition, tps ...string) error {
	switch v := m.(type) {
	case map[string]any:
		return unmarshalerConditionMap(v, val)
	case []any:
		tp := ""
		if len(tps) > 0 {
			tp = tps[0]
		}
		if tp == "" {
			return errors.New("unsupported unmarshaler condition")
		}
		list := make([]Condition, len(v))
		for i, item := range v {
			var itemVal Condition
			if err := unmarshalerCondition(item, &itemVal); err != nil {
				return err
			}
			list[i] = itemVal
		}
		switch tp {
		case "add":
			a := Add(list)
			*val = &a
		case "or":
			o := Or(list)
			*val = &o
		}
		return nil
	}
	return errors.New("unsupported unmarshaler condition")
}

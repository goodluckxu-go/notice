package condition

import (
	"encoding/json"
	"errors"
	"github.com/goodluckxu-go/notice/code"
	"reflect"
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

func UnmarshalerCondition(data []byte, val *Condition) error {
	if len(data) == 0 {
		return nil
	}
	var list []any
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}
	return UnmarshalerConditionList(list, val)
}

func UnmarshalerConditionList(list []any, val *Condition) error {
	if len(list) == 0 {
		return nil
	}
	sign, ok := list[0].(float64)
	if !ok {
		return errors.New("unsupported Condition")
	}
	vType := conditions.Search(Sign(sign))
	if vType == nil {
		return errors.New("unsupported Condition")
	}
	*val = reflect.New(vType).Interface().(Condition)
	return (*val).Unmarshal(list)
}

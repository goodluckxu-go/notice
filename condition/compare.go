package condition

import (
	"encoding/json"
	"errors"
	"github.com/goodluckxu-go/notice/code"
)

func init() {
	RegisterCondition(&Eq{})
	RegisterCondition(&Neq{})
	RegisterCondition(&Gt{})
	RegisterCondition(&Gte{})
	RegisterCondition(&Lt{})
	RegisterCondition(&Lte{})
}

type Eq struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (e *Eq) Verify() bool {
	if e.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(e.metadata[e.Field], e.Value)
	if err != nil {
		return false
	}
	return data == val[0]
}

func (e *Eq) SetMetadata(metadata map[string]*code.Metadata) {
	e.metadata = metadata
}

func (e *Eq) MarshalSign() Sign {
	return 2
}

func (e *Eq) MarshalJSON() ([]byte, error) {
	m := []any{e.MarshalSign(), e.Field, e.Value}
	return json.Marshal(m)
}

func (e *Eq) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Eq' unmarshal fail")
	}
	var ok bool
	if e.Field, ok = list[1].(string); !ok {
		return errors.New("'Eq' unmarshal fail")
	}
	e.Value = list[2]
	return nil
}

type Neq struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (n *Neq) Verify() bool {
	if n.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(n.metadata[n.Field], n.Value)
	if err != nil {
		return false
	}
	return data != val[0]
}

func (n *Neq) SetMetadata(metadata map[string]*code.Metadata) {
	n.metadata = metadata
}

func (n *Neq) MarshalSign() Sign {
	return 3
}

func (n *Neq) MarshalJSON() ([]byte, error) {
	m := []any{n.MarshalSign(), n.Field, n.Value}
	return json.Marshal(m)
}

func (n *Neq) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Neq' unmarshal fail")
	}
	var ok bool
	if n.Field, ok = list[1].(string); !ok {
		return errors.New("'Neq' unmarshal fail")
	}
	n.Value = list[2]
	return nil
}

type Gt struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (g *Gt) Verify() bool {
	if g.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(g.metadata[g.Field], g.Value)
	if err != nil {
		return false
	}
	var floatDate, floatVal float64
	if floatDate, err = covertFloat64(data); err != nil {
		return false
	}
	if floatVal, err = covertFloat64(val[0]); err != nil {
		return false
	}
	return floatDate > floatVal
}

func (g *Gt) SetMetadata(metadata map[string]*code.Metadata) {
	g.metadata = metadata
}

func (g *Gt) MarshalSign() Sign {
	return 4
}

func (g *Gt) MarshalJSON() ([]byte, error) {
	m := []any{g.MarshalSign(), g.Field, g.Value}
	return json.Marshal(m)
}

func (g *Gt) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Gt' unmarshal fail")
	}
	var ok bool
	if g.Field, ok = list[1].(string); !ok {
		return errors.New("'Gt' unmarshal fail")
	}
	g.Value = list[2]
	return nil
}

type Gte struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (g *Gte) Verify() bool {
	if g.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(g.metadata[g.Field], g.Value)
	if err != nil {
		return false
	}
	var floatDate, floatVal float64
	if floatDate, err = covertFloat64(data); err != nil {
		return false
	}
	if floatVal, err = covertFloat64(val[0]); err != nil {
		return false
	}
	return floatDate >= floatVal
}

func (g *Gte) SetMetadata(metadata map[string]*code.Metadata) {
	g.metadata = metadata
}

func (g *Gte) MarshalSign() Sign {
	return 5
}

func (g *Gte) MarshalJSON() ([]byte, error) {
	m := []any{g.MarshalSign(), g.Field, g.Value}
	return json.Marshal(m)
}

func (g *Gte) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Gte' unmarshal fail")
	}
	var ok bool
	if g.Field, ok = list[1].(string); !ok {
		return errors.New("'Gte' unmarshal fail")
	}
	g.Value = list[2]
	return nil
}

type Lt struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (l *Lt) Verify() bool {
	if l.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(l.metadata[l.Field], l.Value)
	if err != nil {
		return false
	}
	var floatDate, floatVal float64
	if floatDate, err = covertFloat64(data); err != nil {
		return false
	}
	if floatVal, err = covertFloat64(val[0]); err != nil {
		return false
	}
	return floatDate < floatVal
}

func (l *Lt) SetMetadata(metadata map[string]*code.Metadata) {
	l.metadata = metadata
}

func (l *Lt) MarshalSign() Sign {
	return 6
}

func (l *Lt) MarshalJSON() ([]byte, error) {
	m := []any{l.MarshalSign(), l.Field, l.Value}
	return json.Marshal(m)
}

func (l *Lt) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Lt' unmarshal fail")
	}
	var ok bool
	if l.Field, ok = list[1].(string); !ok {
		return errors.New("'Lt' unmarshal fail")
	}
	l.Value = list[2]
	return nil
}

type Lte struct {
	Field    string
	Value    any
	metadata map[string]*code.Metadata
}

func (l *Lte) Verify() bool {
	if l.metadata == nil {
		return false
	}
	data, val, err := covertMetadata(l.metadata[l.Field], l.Value)
	if err != nil {
		return false
	}
	var floatDate, floatVal float64
	if floatDate, err = covertFloat64(data); err != nil {
		return false
	}
	if floatVal, err = covertFloat64(val[0]); err != nil {
		return false
	}
	return floatDate <= floatVal
}

func (l *Lte) SetMetadata(metadata map[string]*code.Metadata) {
	l.metadata = metadata
}

func (l *Lte) MarshalSign() Sign {
	return 7
}

func (l *Lte) MarshalJSON() ([]byte, error) {
	m := []any{l.MarshalSign(), l.Field, l.Value}
	return json.Marshal(m)
}

func (l *Lte) Unmarshal(list []any) error {
	if len(list) != 3 {
		return errors.New("'Lte' unmarshal fail")
	}
	var ok bool
	if l.Field, ok = list[1].(string); !ok {
		return errors.New("'Lte' unmarshal fail")
	}
	l.Value = list[2]
	return nil
}

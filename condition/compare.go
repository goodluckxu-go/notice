package condition

import (
	"encoding/json"
	"fmt"
	"github.com/goodluckxu-go/notice/code"
)

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

func (n *Neq) MarshalJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Neq) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, n)
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
	fmt.Println(data, val, err)
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

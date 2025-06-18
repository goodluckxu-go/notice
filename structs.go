package notice

import (
	"fmt"
	"github.com/goodluckxu-go/notice/code"
	"github.com/goodluckxu-go/notice/condition"
	"sync"
)

type service struct {
	recv code.Notice_RecvMessageServer
}

func (s *service) close() {
}

type services struct {
	list []*service
	mux  sync.RWMutex
}

func (s *services) add(ser *service) (no uint32) {
	s.mux.Lock()
	defer s.mux.Unlock()
	for ; no < uint32(len(s.list)) && s.list[no] != nil; no++ {
	}
	if uint32(len(s.list)) == no {
		s.list = append(s.list, ser)
	} else {
		s.list[no] = ser
	}
	return no
}

func (s *services) modify(no uint32, vals ...any) {
	s.mux.Lock()
	defer s.mux.Unlock()
	ser := s.list[no]
	if ser == nil {
		return
	}
	for _, val := range vals {
		switch v := val.(type) {
		case code.Notice_RecvMessageServer:
			ser.recv = v
		}
	}
	s.list[no] = ser
}

func (s *services) exists(no uint32) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if no >= uint32(len(s.list)) {
		return false
	}
	return s.list[no] != nil
}

func (s *services) get(no uint32) *service {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.list[no]
}

func (s *services) del(no uint32) {
	s.mux.Lock()
	defer s.mux.Unlock()
	ser := s.list[no]
	if ser == nil {
		return
	}
	ser.close()
	s.list[no] = nil
}

type Client struct {
	ID       string
	No       uint32
	Metadata map[string]*code.Metadata
}

type clients struct {
	list []*Client
	mux  sync.RWMutex
}

func (c *clients) Add(cl *Client) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	var no = -1
	for i, val := range c.list {
		if val == nil {
			no = i
		} else if val.ID == cl.ID {
			return fmt.Errorf("client %s already exists", cl.ID)
		}
	}
	if no == -1 {
		c.list = append(c.list, cl)
	} else {
		c.list[no] = cl
	}
	return nil
}

func (c *clients) List() []*Client {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.list
}

func (c *clients) Del(id string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for i := 0; i < len(c.list); i++ {
		if c.list[i] != nil && c.list[i].ID == id {
			c.list[i] = nil
		}
	}
}

func (c *clients) DelService(no uint32) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, cl := range c.list {
		if cl.No == no {
			c.list[no] = nil
		}
	}
}

func (c *clients) Search(idList []string, condition condition.Condition) (rs []*Client, err error) {
	for _, cl := range c.list {
		if cl == nil {
			continue
		}
		if len(idList) != 0 && !inArray(cl.ID, idList) {
			continue
		}
		if condition != nil {
			condition.SetMetadata(cl.Metadata)
			if !condition.Verify() {
				continue
			}
		}
		rs = append(rs, cl)
	}
	return
}

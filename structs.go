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
	s.list = append(s.list, ser)
	return uint32(len(s.list)) - 1
}

func (s *services) modify(no uint32, vals ...any) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if no >= uint32(len(s.list)) {
		return
	}
	ser := s.list[no]
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
	return true
}

func (s *services) get(no uint32) *service {
	s.mux.RLock()
	defer s.mux.RUnlock()
	if no >= uint32(len(s.list)) {
		return nil
	}
	return s.list[no]
}

func (s *services) del(no uint32) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if no >= uint32(len(s.list)) {
		return
	}
	tmp := make([]*service, len(s.list)-1)
	n := copy(tmp, s.list[:no])
	copy(tmp[n:], s.list[no+1:])
	s.list = tmp
	cliList.DelService(no)
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
	no := 0
	for ; no < len(c.list) && c.list[no].ID != cl.ID; no++ {
	}
	if no < len(c.list) {
		return fmt.Errorf("client %s already exists", cl.ID)
	}
	c.list = append(c.list, cl)
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
	no := 0
	for ; no < len(c.list) && c.list[no].ID != id; no++ {
	}
	if no < len(c.list) {
		tmp := make([]*Client, len(c.list)-1)
		n := copy(tmp, c.list[:no])
		copy(tmp[n:], c.list[no+1:])
		c.list = tmp
	}
}

func (c *clients) DelService(no uint32) {
	for _, cl := range c.list {
		if cl.No == no {
			c.Del(cl.ID)
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

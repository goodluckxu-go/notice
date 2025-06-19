package notice

import (
	"errors"
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
	m   map[string]*service
	mux sync.RWMutex
}

func (s *services) add(id string, ser *service) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, ok := s.m[id]; ok {
		return errors.New("service exists")
	}
	s.m[id] = ser
	return nil
}

func (s *services) modify(id string, vals ...any) {
	s.mux.Lock()
	defer s.mux.Unlock()
	var ser *service
	var ok bool
	if ser, ok = s.m[id]; !ok {
		return
	}
	for _, val := range vals {
		switch v := val.(type) {
		case code.Notice_RecvMessageServer:
			ser.recv = v
		}
	}
	s.m[id] = ser
}

func (s *services) exists(id string) bool {
	s.mux.RLock()
	defer s.mux.RUnlock()
	_, ok := s.m[id]
	return ok
}

func (s *services) get(id string) *service {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.m[id]
}

func (s *services) del(id string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.m, id)
	cliList.delService(id)
}

type Client struct {
	ID        string
	ServiceID string
	Metadata  map[string]*code.Metadata
}

type clients struct {
	list []*Client
	mux  sync.RWMutex
}

func (c *clients) add(cl *Client) error {
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

func (c *clients) del(id string) {
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

func (c *clients) delService(serviceID string) {
	for _, cl := range c.list {
		if cl.ServiceID == serviceID {
			c.del(cl.ID)
		}
	}
}

func (c *clients) search(idList []string, condition condition.Condition) (rs map[string][]string, err error) {
	rs = map[string][]string{}
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
		rs[cl.ServiceID] = append(rs[cl.ServiceID], cl.ID)
	}
	return
}

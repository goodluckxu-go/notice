package notice

import (
	"errors"
	"github.com/goodluckxu-go/notice/code"
	"github.com/goodluckxu-go/notice/condition"
	"sync"
)

type ServerChan struct {
	message  []byte
	clientID string
}

type Server struct {
	m   map[string]code.Notice_RecvMessageServer
	mux sync.Mutex
}

func (s *Server) IsRegistered(id string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.m[id]
	return ok
}

func (s *Server) Add(id string, server code.Notice_RecvMessageServer) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[id] = server
}

func (s *Server) Del(id string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	clients.DelServer(id)
	delete(s.m, id)
}

func (s *Server) Get(id string) (code.Notice_RecvMessageServer, bool) {
	ser, ok := s.m[id]
	return ser, ok
}

var servers = &Server{m: map[string]code.Notice_RecvMessageServer{}}

type Client struct {
	id       string
	serverID string
	metadata map[string]*code.Metadata
}

type Clients struct {
	list []Client
	mux  sync.Mutex
}

func (t *Clients) Add(id, serverID string, metadata map[string]*code.Metadata) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	var idList []string
	for _, v := range t.list {
		idList = append(idList, v.id)
	}
	if inArray(id, idList) {
		return errors.New("client is exist")
	}
	t.list = append(t.list, Client{id, serverID, metadata})
	return nil
}

func (t *Clients) Del(id, serverID string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	index := 0
	for index < len(t.list) {
		item := t.list[index]
		if item.id == id && item.serverID == serverID {
			t.list = append(t.list[:index], t.list[index+1:]...)
		} else {
			index++
		}
	}
}

func (t *Clients) DelServer(serverID string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	index := 0
	for index < len(t.list) {
		item := t.list[index]
		if item.serverID == serverID {
			t.list = append(t.list[:index], t.list[index+1:]...)
		} else {
			index++
		}
	}
}

type Condition struct {
	Field string
	Type  string
	Value any
}

func (t *Clients) Search(idList []string, condition condition.Condition) (rs []Client, err error) {
	for _, cli := range t.list {
		if len(idList) != 0 && !inArray(cli.id, idList) {
			continue
		}
		if condition != nil {
			condition.SetMetadata(cli.metadata)
			if !condition.Verify() {
				continue
			}
		}
		rs = append(rs, cli)
	}
	return
}

var clients = &Clients{}

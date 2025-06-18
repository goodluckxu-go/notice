package notice

import "github.com/goodluckxu-go/notice/condition"

type ClientInterface interface {
	Add(cl *Client) error
	List() []*Client
	Del(id string)
	DelService(no uint32)
	Search(idList []string, condition condition.Condition) (rs []*Client, err error)
}

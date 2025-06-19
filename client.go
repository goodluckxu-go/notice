package notice

import (
	"context"
	"encoding/json"
	pb "github.com/goodluckxu-go/notice/code"
	cond "github.com/goodluckxu-go/notice/condition"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"time"
)

type NoticeClient struct {
	serviceID string
	conn      *grpc.ClientConn
	client    pb.NoticeClient
	readyRecv chan struct{}
}

func (c *NoticeClient) AddClient(id string, metadata map[string]any) (err error) {
	mData := map[string]*pb.Metadata{}
	for k, v := range metadata {
		mData[k] = toMetadata(v)
	}
	err = cliList.add(&Client{ID: id, ServiceID: c.serviceID, Metadata: mData})
	if err != nil {
		return
	}
	var add pb.Notice_AddClientClient
	if add, err = c.client.AddClient(context.Background()); err != nil {
		return
	}
	err = add.Send(&pb.ClientReq{ServiceID: c.serviceID, Id: id, Metadata: mData})
	if err != nil {
		return
	}
	_, _ = add.CloseAndRecv()
	return
}

func (c *NoticeClient) DelClient(id string) (err error) {
	cliList.del(id)
	_, err = c.client.DelClient(context.Background(), &pb.ClientReq{ServiceID: c.serviceID, Id: id})
	return
}

func (c *NoticeClient) SendMessage(message []byte, idList []string, condition cond.Condition) (err error) {
	var buf []byte
	if buf, err = json.Marshal(condition); err != nil {
		return
	}
	_, err = c.client.SendMessage(context.Background(), &pb.SendReq{ServiceID: c.serviceID, Message: message,
		IdList: idList, Condition: buf})
	return
}

func (c *NoticeClient) RecvMessage(cb func(id string, message []byte)) error {
	for {
		select {
		case <-c.readyRecv:
			steam, err := c.client.RecvMessage(context.Background(), &pb.Service{ServiceID: c.serviceID})
			if err != nil {
				return err
			}
		steamPos:
			for {
				select {
				case <-steam.Context().Done():
					break steamPos
				default:
					var recv *pb.RecvResp
					for {
						recv, err = steam.Recv()
						if err != nil {
							break steamPos
						}
						go func(recv *pb.RecvResp) {
							for _, id := range recv.IdList {
								cb(id, recv.Message)
							}
						}(recv)
					}
				}
			}
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func (c *NoticeClient) Close() {
	_ = c.conn.Close()
}

func (c *NoticeClient) handleReady(isReady bool) {
	if !isReady {
		return
	}
	// 注册服务
	registerResp, err := c.client.Register(context.Background(), nil)
	if err != nil {
		return
	}
	c.serviceID = registerResp.ServiceID
	// 添加客户端
	var add pb.Notice_AddClientClient
	if add, err = c.client.AddClient(context.Background()); err != nil {
		return
	}
	for _, client := range cliList.list {
		if client == nil {
			continue
		}
		_ = add.Send(&pb.ClientReq{ServiceID: c.serviceID, Id: client.ID, Metadata: client.Metadata})
	}
	_, _ = add.CloseAndRecv()
	c.readyRecv <- struct{}{}
}

func (c *NoticeClient) checkStatus() {
	c.handleReady(true)
	isReady := false
	var state connectivity.State
	for {
		c.conn.WaitForStateChange(context.Background(), state)
		state = c.conn.GetState()
		if state == connectivity.Idle {
			c.conn.Connect()
		} else if state == connectivity.Ready {
			go c.handleReady(isReady)
			isReady = true
		}
	}
}

func Dail(addr string, opts ...grpc.DialOption) (*NoticeClient, error) {
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}
	c := &NoticeClient{conn: conn, client: pb.NewNoticeClient(conn), readyRecv: make(chan struct{})}
	go c.checkStatus()
	return c, nil
}

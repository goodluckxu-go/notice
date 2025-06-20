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
	serviceID    string
	conn         *grpc.ClientConn
	client       pb.NoticeClient
	sendMessages *sendMsgs
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
	var clientList map[string][]string
	if clientList, err = cliList.search(idList, condition); err != nil {
		return
	}
	c.sendMessages.add(message, clientList[c.serviceID]...)
	_, err = c.client.SendMessage(context.Background(), &pb.SendReq{ServiceID: c.serviceID, Message: message,
		IdList: idList, Condition: buf})
	return
}

func (c *NoticeClient) RecvMessage(cb func(id string, message []byte)) error {
	for {
		time.Sleep(10 * time.Millisecond)
		list := c.sendMessages.list()
		if len(list) == 0 {
			continue
		}
		for _, msg := range list {
			cb(msg.id, msg.message)
		}
	}
}

func (c *NoticeClient) recvMessage(steam pb.Notice_RecvMessageClient) {
	for {
		recv, err := steam.Recv()
		if err != nil {
			return
		}
		c.sendMessages.add(recv.Message, recv.IdList...)
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
	_, err := c.client.Register(context.Background(), &pb.Service{ServiceID: c.serviceID})
	if err != nil {
		return
	}
	// 接受消息
	var steam pb.Notice_RecvMessageClient
	steam, err = c.client.RecvMessage(context.Background(), &pb.Service{ServiceID: c.serviceID})
	if err != nil {
		return
	}
	go c.recvMessage(steam)
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
	c := &NoticeClient{serviceID: getUUID(), conn: conn, client: pb.NewNoticeClient(conn), sendMessages: &sendMsgs{}}
	go c.checkStatus()
	return c, nil
}

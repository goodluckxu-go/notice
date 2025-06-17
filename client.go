package notice

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/goodluckxu-go/notice/code"
	cond "github.com/goodluckxu-go/notice/condition"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"time"
)

type NoticeClient struct {
	serverID  string
	conn      *grpc.ClientConn
	client    pb.NoticeClient
	readyRecv chan struct{}
}

func (c *NoticeClient) AddClient(id string, metadata map[string]any) (err error) {
	mData := map[string]*pb.Metadata{}
	for k, v := range metadata {
		mData[k] = toMetadata(v)
	}
	clients.Add(id, c.serverID, mData)
	_, err = c.client.AddClient(context.Background(), &pb.ClientReq{Server: &pb.ServerReq{Id: c.serverID},
		Id: id, Metadata: mData})
	return
}

func (c *NoticeClient) DelClient(id string) (err error) {
	clients.Del(id, c.serverID)
	_, err = c.client.DelClient(context.Background(), &pb.ClientReq{Server: &pb.ServerReq{Id: c.serverID},
		Id: id})
	return
}

func (c *NoticeClient) SendMessage(message []byte, idList []string, condition cond.Condition) (err error) {
	var buf []byte
	if buf, err = json.Marshal(condition); err != nil {
		return
	}
	_, err = c.client.SendMessage(context.Background(), &pb.SendReq{Server: &pb.ServerReq{Id: c.serverID},
		Message: message, IdList: idList, Condition: buf})
	return
}

func (c *NoticeClient) RecvMessage(cb func(id string, message []byte)) error {
	for {
		select {
		case <-c.readyRecv:
			steam, err := c.client.RecvMessage(context.Background(), &pb.ServerReq{Id: c.serverID})
			if err != nil {
				return err
			}
			var recv *pb.RecvResp
			for {
				recv, err = steam.Recv()
				if err != nil {
					break
				}
				if recv.Heartbeat {
					continue
				}
				cb(recv.ClientID, recv.Message)
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
	_, err := c.client.Register(context.Background(), &pb.ServerReq{Id: c.serverID})
	if err != nil {
		return
	}
	// 添加客户端
	for _, client := range clients.list {
		_, err = c.client.AddClient(context.Background(), &pb.ClientReq{Server: &pb.ServerReq{Id: c.serverID},
			Id: client.id, Metadata: client.metadata})
	}
	c.readyRecv <- struct{}{}
}

func (c *NoticeClient) checkStatus() {
	c.handleReady(true)
	isReady := false
	var state connectivity.State
	for {
		c.conn.WaitForStateChange(context.Background(), state)
		state = c.conn.GetState()
		fmt.Println("state:", state)
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
	c := &NoticeClient{serverID: getUUID(), conn: conn, client: pb.NewNoticeClient(conn), readyRecv: make(chan struct{})}
	go c.checkStatus()
	return c, nil
}

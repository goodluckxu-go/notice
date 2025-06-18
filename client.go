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
	no        int
	conn      *grpc.ClientConn
	client    pb.NoticeClient
	readyRecv chan struct{}
}

func (c *NoticeClient) AddClient(id string, metadata map[string]any) (err error) {
	mData := map[string]*pb.Metadata{}
	for k, v := range metadata {
		mData[k] = toMetadata(v)
	}
	err = cliList.Add(&Client{ID: id, No: uint32(c.no), Metadata: mData})
	if err != nil {
		return
	}
	var add pb.Notice_AddClientClient
	if add, err = c.client.AddClient(context.Background()); err != nil {
		return
	}
	err = add.Send(&pb.ClientReq{No: uint32(c.no),
		Id: id, Metadata: mData})
	if err != nil {
		return
	}
	_ = add.CloseSend()
	return
}

func (c *NoticeClient) DelClient(id string) (err error) {
	cliList.Del(id)
	_, err = c.client.DelClient(context.Background(), &pb.ClientReq{No: uint32(c.no), Id: id})
	return
}

func (c *NoticeClient) SendMessage(message []byte, idList []string, condition cond.Condition) (err error) {
	var buf []byte
	if buf, err = json.Marshal(condition); err != nil {
		return
	}
	_, err = c.client.SendMessage(context.Background(), &pb.SendReq{No: uint32(c.no), Message: message,
		IdList: idList, Condition: buf})
	return
}

func (c *NoticeClient) RecvMessage(cb func(id string, message []byte)) error {
	for {
		select {
		case <-c.readyRecv:
			steam, err := c.client.RecvMessage(context.Background(), &pb.Number{No: uint32(c.no)})
			if err != nil {
				return err
			}
			var recv *pb.RecvResp
			for {
				recv, err = steam.Recv()
				if err != nil {
					break
				}
				cb(recv.Id, recv.Message)
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
	numResp, err := c.client.Register(context.Background(), nil)
	if err != nil {
		return
	}
	c.no = int(numResp.No)
	// 添加客户端
	var add pb.Notice_AddClientClient
	if add, err = c.client.AddClient(context.Background()); err != nil {
		return
	}
	for _, client := range cliList.List() {
		if client == nil {
			continue
		}
		_ = add.Send(&pb.ClientReq{No: uint32(c.no), Id: client.ID, Metadata: client.Metadata})
	}
	_ = add.CloseSend()
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

func SetClient(client ClientInterface) {
	cliList = client
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

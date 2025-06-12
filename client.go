package notice

import (
	"context"
	pb "github.com/goodluckxu-go/notice/code"
	cond "github.com/goodluckxu-go/notice/condition"
	"google.golang.org/grpc"
)

type NoticeClient struct {
	serverID string
	conn     *grpc.ClientConn
	client   pb.NoticeClient
}

func (c *NoticeClient) AddClient(id string, metadata map[string]any) (err error) {
	mData := map[string]*pb.Metadata{}
	for k, v := range metadata {
		mData[k] = toMetadata(v)
	}
	_, err = c.client.AddClient(context.Background(), &pb.ClientReq{Server: &pb.ServerReq{Id: c.serverID},
		Id: id, Metadata: mData})
	return
}

func (c *NoticeClient) DelClient(id string) (err error) {
	_, err = c.client.DelClient(context.Background(), &pb.ClientReq{Server: &pb.ServerReq{Id: c.serverID},
		Id: id})
	return
}

func (c *NoticeClient) SendMessage(message []byte, idList []string, condition cond.Condition) (err error) {
	var buf []byte
	if buf, err = cond.MarshalerCondition(condition); err != nil {
		return
	}
	_, err = c.client.SendMessage(context.Background(), &pb.SendReq{Server: &pb.ServerReq{Id: c.serverID},
		Message: message, IdList: idList, Condition: buf})
	return
}

func (c *NoticeClient) RecvMessage(cb func(id string, message []byte)) error {
	steam, err := c.client.RecvMessage(context.Background(), &pb.ServerReq{Id: c.serverID})
	if err != nil {
		return err
	}
	var recv *pb.RecvResp
	for {
		recv, err = steam.Recv()
		if err != nil {
			return err
		}
		if recv.Heartbeat {
			continue
		}
		cb(recv.ClientID, recv.Message)
	}
}

func (c *NoticeClient) Close() {
	_ = c.conn.Close()
}

func Dail(addr string) (*NoticeClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	serverID := getUUID()
	client := pb.NewNoticeClient(conn)
	_, err = client.Register(context.Background(), &pb.ServerReq{Id: serverID})
	if err != nil {
		return nil, err
	}
	return &NoticeClient{serverID: serverID, conn: conn, client: client}, nil
}

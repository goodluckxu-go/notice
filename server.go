package notice

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/goodluckxu-go/notice/code"
	"github.com/goodluckxu-go/notice/condition"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"time"
)

type NoticeServer struct {
	pb.UnimplementedNoticeServer
}

func (c *NoticeServer) Register(ctx context.Context, req *pb.ServerReq) (*emptypb.Empty, error) {
	if servers.IsRegistered(req.GetId()) {
		return nil, errors.New("server is already registered")
	}
	servers.Add(req.GetId(), nil)
	return nil, nil
}

func (c *NoticeServer) AddClient(ctx context.Context, req *pb.ClientReq) (*emptypb.Empty, error) {
	if !servers.IsRegistered(req.GetServer().GetId()) {
		return nil, errors.New("server not found")
	}
	err := clients.Add(req.GetId(), req.GetServer().GetId(), req.Metadata)
	return nil, err
}

func (c *NoticeServer) DelClient(ctx context.Context, req *pb.ClientReq) (*emptypb.Empty, error) {
	if !servers.IsRegistered(req.GetServer().GetId()) {
		return nil, errors.New("server not found")
	}
	clients.Del(req.GetId(), req.GetServer().GetId())
	return nil, nil
}

func (c *NoticeServer) SendMessage(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	var cond condition.Condition
	if err := condition.UnmarshalerCondition(req.Condition, &cond); err != nil {
		fmt.Println(err)
		return nil, err
	}
	clientList, err := clients.Search(req.IdList, cond)
	if err != nil {
		return nil, err
	}
	for _, client := range clientList {
		if ser, ok := servers.Get(client.serverID); ok {
			go ser.Send(&pb.RecvResp{
				ClientID: client.id,
				Message:  req.Message,
			})
		}
	}
	return nil, nil
}

func (c *NoticeServer) RecvMessage(req *pb.ServerReq, stream pb.Notice_RecvMessageServer) error {
	if _, ok := servers.Get(req.GetId()); !ok {
		return errors.New("server not found")
	}
	servers.Add(req.GetId(), stream)
	select {
	case <-stream.Context().Done():
	}
	return nil
}

func (c *NoticeServer) heartbeat(serverID string, stream pb.Notice_RecvMessageServer, close chan struct{}) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			err := stream.Send(&pb.RecvResp{Heartbeat: true})
			if err != nil {
				servers.Del(serverID)
				close <- struct{}{}
				ticker.Stop()
			}
		}
	}
}

func Listen(addr string, opts ...grpc.ServerOption) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	server := grpc.NewServer(opts...)
	pb.RegisterNoticeServer(server, &NoticeServer{})
	return server.Serve(listener)
}

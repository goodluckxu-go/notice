package notice

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/goodluckxu-go/notice/code"
	"github.com/goodluckxu-go/notice/condition"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"net"
)

type NoticeServer struct {
	pb.UnimplementedNoticeServer
}

func (c *NoticeServer) Register(context.Context, *emptypb.Empty) (*pb.Number, error) {
	no := serList.add(&service{})
	return &pb.Number{No: no}, nil
}

func (c *NoticeServer) AddClient(stream pb.Notice_AddClientServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if !serList.exists(req.No) {
			return errors.New("server not found")
		}
		_ = cliList.Add(&Client{
			No:       req.GetNo(),
			ID:       req.GetId(),
			Metadata: req.GetMetadata(),
		})
	}
	return nil
}

func (c *NoticeServer) DelClient(ctx context.Context, req *pb.ClientReq) (*emptypb.Empty, error) {
	if !serList.exists(req.GetNo()) {
		return nil, errors.New("server not found")
	}
	cliList.Del(req.GetId())
	return nil, nil
}

func (c *NoticeServer) SendMessage(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	var cond condition.Condition
	if err := condition.UnmarshalerCondition(req.Condition, &cond); err != nil {
		return nil, err
	}
	clientList, err := cliList.Search(req.IdList, cond)
	if err != nil {
		return nil, err
	}
	for _, client := range clientList {
		if ser := serList.get(client.No); ser != nil {
			go ser.recv.Send(&pb.RecvResp{
				Id:      client.ID,
				Message: req.Message,
			})
		}
	}
	return nil, nil
}

func (c *NoticeServer) RecvMessage(req *pb.Number, stream pb.Notice_RecvMessageServer) error {
	if !serList.exists(req.No) {
		return fmt.Errorf("server not found")
	}
	serList.modify(req.No, stream)
	select {
	case <-stream.Context().Done():
		serList.del(req.No)
	}
	return nil
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

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

func (c *NoticeServer) Register(context.Context, *emptypb.Empty) (*pb.Service, error) {
	id := getUUID()
	err := serList.add(id, &service{})
	if err != nil {
		return nil, err
	}
	return &pb.Service{ServiceID: id}, nil
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
		if !serList.exists(req.GetServiceID()) {
			return errors.New("server not found")
		}
		_ = cliList.add(&Client{
			ServiceID: req.GetServiceID(),
			ID:        req.GetId(),
			Metadata:  req.GetMetadata(),
		})
	}
	return nil
}

func (c *NoticeServer) DelClient(ctx context.Context, req *pb.ClientReq) (*emptypb.Empty, error) {
	if !serList.exists(req.GetServiceID()) {
		return nil, errors.New("server not found")
	}
	cliList.del(req.GetId())
	return nil, nil
}

func (c *NoticeServer) SendMessage(ctx context.Context, req *pb.SendReq) (*emptypb.Empty, error) {
	var cond condition.Condition
	if err := condition.UnmarshalerCondition(req.GetCondition(), &cond); err != nil {
		return nil, err
	}
	clientList, err := cliList.search(req.GetIdList(), cond)
	if err != nil {
		return nil, err
	}
	for serviceID, idList := range clientList {
		if ser := serList.get(serviceID); ser != nil {
			go ser.recv.Send(&pb.RecvResp{
				IdList:  idList,
				Message: req.GetMessage(),
			})
		}
	}
	return nil, nil
}

func (c *NoticeServer) RecvMessage(req *pb.Service, stream pb.Notice_RecvMessageServer) error {
	if !serList.exists(req.GetServiceID()) {
		return fmt.Errorf("server not found")
	}
	serList.modify(req.GetServiceID(), stream)
	select {
	case <-stream.Context().Done():
		serList.del(req.GetServiceID())
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

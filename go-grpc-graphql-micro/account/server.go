package account

import (
	"context"
	"fmt"
	"net"

	"github.com/Dacoband/Go-gRPC-GraphQL-Microservice/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountRespone, error) {
	a, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   a.ID,
		Name: a.Name,
	}}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountRespone, error) {
	a, err := s.service.GetAccount(ctx, r.id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountRespone{
		Account: &pb.Account{
			Id:   a.ID,
			Name: a.Name,
		}}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsRespone, error) {
	res, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(
			accounts,
			&pb.Account{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetAccountsRespone{Accounts: accounts}, nil
}

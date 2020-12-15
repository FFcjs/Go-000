package service

import (
	"context"
	"log"

	pb "Go-000/Week04/api/myservice"
	"Go-000/Week04/internal/myservice/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	pb.UnimplementedGreeterServer
	u *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(u *biz.UserUsecase) *GreeterService {
	return &GreeterService{u:u}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//dto->do
	log.Printf("Received: %v", in.GetName())
	o:=&biz.User{Name: in.GetName()}
	s.u.CreateUser(o)
	return &pb.HelloReply{Message: "create " + in.GetName()+"success"}, nil
}


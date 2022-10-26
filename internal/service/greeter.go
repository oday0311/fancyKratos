package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "fancyKrato/api/helloworld/v1"
	"fancyKrato/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) GetAccount(ctx context.Context, in *v1.AccountRequest) (*v1.AccountResponse, error) {

	//get data from DB.
	//g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	log.Debug("GetAccount: ", in.GetName())

	return &v1.AccountResponse{Message: "Hello " + " this is the account model"}, nil
}

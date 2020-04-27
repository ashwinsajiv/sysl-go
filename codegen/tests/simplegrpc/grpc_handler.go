// Code generated by sysl DO NOT EDIT.
package simplegrpc

import (
	"context"

	"github.com/anz-bank/sysl-go/codegen/tests/simple"
	pb "github.com/anz-bank/sysl-go/codegen/tests/simplepb"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/validator"
	"google.golang.org/grpc"
)

// GrpcServiceHandler for SimpleGrpc API
type GrpcServiceHandler struct {
	genCallback         core.GrpcGenCallback
	serviceInterface    *GrpcServiceInterface
	unimpl              *pb.UnimplementedSimpleGrpcServer
	simpleSimpleService simple.Service
}

// NewGrpcServiceHandler for SimpleGrpc
func NewGrpcServiceHandler(genCallback core.GrpcGenCallback, serviceInterface *GrpcServiceInterface, simpleSimpleService simple.Service) *GrpcServiceHandler {
	return &GrpcServiceHandler{genCallback, serviceInterface, &(pb.UnimplementedSimpleGrpcServer{}), simpleSimpleService}
}

// RegisterServer registers the SimpleGrpc gRPC service
func (s *GrpcServiceHandler) RegisterServer(ctx context.Context, server *grpc.Server) {
	pb.RegisterSimpleGrpcServer(server, s)
}

// GetStuff ...
func (s *GrpcServiceHandler) GetStuff(ctx context.Context, req *pb.GetStuffRequest) (*pb.GetStuffResponse, error) {
	if s.serviceInterface.GetStuff == nil {
		return s.unimpl.GetStuff(ctx, req)
	}

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	client := GetStuffClient{
		PostStuff: s.simpleSimpleService.PostStuff,
	}

	return s.serviceInterface.GetStuff(ctx, req, client)
}

// Config ...
func (s *GrpcServiceHandler) Config() validator.Validator {
	return s.genCallback.Config()
}

// Name ...
func (s *GrpcServiceHandler) Name() string {
	return "SimpleGrpc"
}

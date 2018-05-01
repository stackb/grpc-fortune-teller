package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"path"

	"github.com/stackb/fortune-teller/app/fileutil"
	proto "github.com/stackb/fortune-teller/proto/fortune"
	"github.com/vromero/gofortune/lib/fortune"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
)

func main() {

	baseDir := "/tmp/fortune-teller"
	fileutil.MustMkdirAll(baseDir)

	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(200),
	}

	grpcServer := grpc.NewServer(opts...)

	fortuneTeller := &FortuneTeller{
		fs: createFortuneFilesystemNodeDescriptor(baseDir),
	}
	proto.RegisterFortuneTellerServer(grpcServer, fortuneTeller)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Error while starting grpc server: %v\n", err)
	}

	log.Printf("Listening for gRPC requests at %d\n", grpcPort)
	grpcServer.Serve(lis)
}

// FortuneTeller - struct that will implement the grpc service interface.
type FortuneTeller struct {
	fs *fortune.FileSystemNodeDescriptor
}

// Predict - implementation for the grpc unary request method.
func (f *FortuneTeller) Predict(ctx context.Context, r *proto.PredictionRequest) (*proto.PredictionResponse, error) {
	_, data, err := fortune.GetRandomFortune(*f.fs)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Unable to render fortune: %v", err)
	}
	return &proto.PredictionResponse{
		Message: data,
	}, nil
}

func createFortuneFilesystemNodeDescriptor(baseDir string) *fortune.FileSystemNodeDescriptor {

	// Restore the packed fortune data
	fortuneDir := path.Join(baseDir, "usr/share/games/fortunes")

	fileutil.MustRestore(baseDir, fortuneFiles, nil)

	// init gofortune fs
	fs, err := fortune.LoadPaths([]fortune.ProbabilityPath{
		{Path: fortuneDir},
	})
	if err != nil {
		log.Fatalf("Unable to load fortune paths: %v", err)
	}

	fortune.SetProbabilities(&fs, true) // consider all equal probabilities
	return &fs
}

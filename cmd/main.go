package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spoonscen/go-proto-grpc-demo/foo"
	"google.golang.org/grpc"
)

// FooServer is an implementation of the proto Foo service.
type FooServer struct {
}

// NewFooServer returns a pointer to a FooServer.
func NewFooServer() *FooServer {
	return &FooServer{}
}

// SayHello implements the SayHello rpc method from the proto Foo service.
func (f *FooServer) SayHello(ctx context.Context, req *foo.HelloRequest) (*foo.HelloResponse, error) {
	return &foo.HelloResponse{
		Greeting: "Hello " + req.GetName(),
	}, nil
}

func main() {
	s := grpc.NewServer()
	foo.RegisterFooServer(s, NewFooServer())

	l, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalf("listen tcp %s", err.Error())
	}

	go func() {
		log.Printf("Listening on 4444")
		s.Serve(l)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-stop
}

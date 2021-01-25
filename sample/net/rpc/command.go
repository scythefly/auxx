package rpc

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	return &String{Value: "hello:" + args.GetValue()}, nil
}

func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &String{Value: "stream hello:" + args.GetValue()}
		if err = stream.Send(reply); err != nil {
			return err
		}
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////

type TaskServiceImpl struct{}

func (p *TaskServiceImpl) Channel(stream TaskService_ChannelServer) error {
	var (
		err    error
		str    *String
		status int32
	)

	for {
		status++
		str, err = stream.Recv()
		fmt.Println("TaskService channel recv", str.GetValue())
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			goto END
		}
		if status > 5 {
			goto END
		}
	}
END:
	fmt.Println("TaskService channel SendAndClose")
	return stream.SendAndClose(&Response{Status: status})
}

// ////////////////////////////////////////////////////////////////////////////////////////

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "rpc",
		RunE: runGRpc,
	}
	return cmd
}

func runGRpc(_ *cobra.Command, _ []string) error {
	server := grpc.NewServer()
	RegisterHelloServiceServer(server, new(HelloServiceImpl))
	RegisterTaskServiceServer(server, new(TaskServiceImpl))
	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	go runClient(ctx)
	go func() {
		server.Serve(l)
	}()
	time.Sleep(15 * time.Second)
	fmt.Println(">>> server stop")
	server.Stop()
	select {
	case <-ctx.Done():
	}
	return nil
}

func runClient(ctx context.Context) {
	time.Sleep(5 * time.Second)
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := NewHelloServiceClient(conn)
	stream, err := client.Channel(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	go clientStreamRecv(stream)
	go clientStreamSend(stream)
	go func() {
		for i := 0; i < 5; i++ {
			reply, err := client.Hello(context.Background(), &String{Value: fmt.Sprintf("hello_%d", i)})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(reply.GetValue())
			time.Sleep(time.Second)
		}
	}()

	tc := NewTaskServiceClient(conn)
	tstream, err := tc.Channel(context.Background())
	go tclientStreamSend(tstream)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(10 * time.Second)
}

func clientStreamSend(stream HelloService_ChannelClient) {
	var idx int
	for {
		idx++
		if err := stream.Send(&String{Value: fmt.Sprintf("stream hi_%d", idx)}); err != nil {
			fmt.Println("client stream send:", err)
			return
		}
		time.Sleep(time.Second)
	}
}

func clientStreamRecv(stream HelloService_ChannelClient) {
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("client stream recv:", err)
			return
		}
		fmt.Println(">> client stream recv:", reply.GetValue())
	}
}

func tclientStreamSend(stream TaskService_ChannelClient) {
	var idx int
	for {
		idx++
		if err := stream.Send(&String{Value: fmt.Sprintf("task stream hi_%d", idx)}); err != nil {
			fmt.Println("task client stream send:", err)
			return
		}
		time.Sleep(time.Second)
	}
}

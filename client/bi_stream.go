package main

import (
	"context"
	pb "github.com/Renewdxin/GO-grpc/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) error {
	// 打印消息，表示双向流已开始
	log.Printf("Bidirectional streaming started")

	// 创建双向流
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	// 创建一个用于通知主协程和接收协程完成任务的通道
	waitc := make(chan struct{})

	// 启动接收协程
	go func() {
		// 循环接收服务器发送的消息
		for {
			message, err := stream.Recv()
			// 如果遇到 io.EOF，表示服务器关闭了发送流，退出循环
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		// 关闭通知通道，通知主协程接收任务已完成
		close(waitc)
	}()

	// 在主协程中，遍历 NamesList 中的名称并发送到服务器
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	// 关闭发送流，告知服务器不会再发送更多消息
	stream.CloseSend()

	// 等待接收协程完成任务
	<-waitc

	// 打印消息，表示双向流已结束
	log.Printf("Bidirectional streaming finished")

	return nil
}

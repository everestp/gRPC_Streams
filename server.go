package main

import (
	mainapipb "grpc_stream/proto/gen"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)


type server struct{
	 mainapipb.UnimplementedCalculatorServer
}

func (s  *server) GenerateFibonacci(req *mainapipb.FibonacciRequest , stream mainapipb.Calculator_GenerateFibbonacciServer) error{
 n := req.N
 a,b := 0,1

 for i := 0 ; i <int(n); i++{
	err := stream.Send(&mainapipb.FibonacciResponse{
		Number: int32(a),
	})
	if err != nil{
		return  err
	}
	a,b =b,a+b
	time.Sleep(time.Second)
 }
 return nil
}

func main(){
lis , err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	mainapipb.RegisterCalculatorServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil{
		log.Fatal(err)
	}

}
package encoder

import (
	"errors"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	grpcr "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func New(addr string) (*Client, error) {
	resolveTypes(addr)
	return &Client{
		encoder: &Encoder{
			types: make(map[string]*typeInfo),
		},
	}, nil
}


func resolveTypes(addr string) (map[string]*typeInfo, error){
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	rcl := grpcr.NewServerReflectionClient(conn)

	client, err := rcl.ServerReflectionInfo(context.Background())
	if err != nil {
		return nil, err
	}
	err = client.Send(&grpcr.ServerReflectionRequest{
		MessageRequest: &grpcr.ServerReflectionRequest_ListServices{},
	})
	if err != nil {
		return nil, err
	}

	resp, err := client.Recv()
	if err != nil {
		return nil, err
	}

	if err := resp.GetErrorResponse(); err != nil {
		return nil, fmt.Errorf("ListServices error: %v", err.GetErrorMessage())
	}

	lresp := resp.GetListServicesResponse()


	res := make(map[string]*typeInfo)


	for _, s := range lresp.Service {
		err = client.Send(&grpcr.ServerReflectionRequest{
			MessageRequest: &grpcr.ServerReflectionRequest_FileContainingSymbol{
				FileContainingSymbol: s.GetName(),
			},
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.Recv()
		if err != nil {
			return nil, err
		}

		if err := resp.GetErrorResponse(); err != nil {
			return nil, fmt.Errorf("FileContainingSymbol error: %v", err.GetErrorMessage())
		}

		fresp := resp.GetFileDescriptorResponse()



		for _, desc := range fresp.FileDescriptorProto {
			file := descriptor.FileDescriptorProto{}
			if err := proto.Unmarshal(desc, &file); err != nil {
				return nil, err
			}

			for _, tp := range file.MessageType {
				log.Printf("%+v", tp.GetName())
			}

			for _, service := range file.GetService() {
				for _, m := range service.GetMethod() {
					log.Printf("%+v => %+v", m.GetInputType(), m.GetOutputType())

				}
			}
		}
	}


	return res, nil
}

type Client struct {
	encoder *Encoder
}

func (c *Client) Invoke(req []*Field) (interface{}, error) {
	return nil, errors.New("not implemented")
}

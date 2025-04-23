package client

import (
	"context"
	"github.com/9triver/cfn/proto"
	"github.com/9triver/cfn/proto/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	address string
	conn    *grpc.ClientConn
	/* service clients */
	functionServiceClient data.FunctionServiceClient
}

func NewClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		address: address,
		conn:    conn,
	}

	/* init service clients */
	client.functionServiceClient = data.NewFunctionServiceClient(conn)
	return client, nil
}

func (client *Client) DeployPyFunc(appendPyFunc *data.AppendPyFunc) (*proto.ServiceReplay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	replay, err := client.functionServiceClient.DeployPyFunc(ctx, appendPyFunc)
	if err != nil {
		return nil, err
	}
	return replay, nil
}

func (client *Client) Close() {
	client.conn.Close()
}

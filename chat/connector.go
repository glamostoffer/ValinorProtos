package chat

import (
	"context"
	"github.com/glamostoffer/ValinorProtos/chat/client_chat"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connector struct {
	ClientChat client_chat.ClientChatServiceClient
	Config     Config
	conn       *grpc.ClientConn
}

func New(cfg Config) *Connector {
	return &Connector{
		Config: cfg,
	}
}

func (c *Connector) Start(ctx context.Context) (err error) {
	if c.Config.DontRun {
		return nil
	}

	c.conn, err = grpc.DialContext(
		ctx,
		c.Config.Address,
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.FailOnNonTempDialError(true),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	c.ClientChat = client_chat.NewClientChatServiceClient(c.conn)
	//c.AdminAuth = admin_auth.NewAdminAuthServiceClient(c.conn)

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

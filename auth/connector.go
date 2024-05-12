package auth

import (
	"context"
	"github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	"github.com/glamostoffer/ValinorProtos/auth/client_auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

type Connector struct {
	ClientAuth client_auth.ClientAuthServiceClient
	AdminAuth  admin_auth.AdminAuthServiceClient
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

	c.ClientAuth = client_auth.NewClientAuthServiceClient(c.conn)
	c.AdminAuth = admin_auth.NewAdminAuthServiceClient(c.conn)

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

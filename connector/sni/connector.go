package sni

import (
	"context"
	"net"

	ctxvalue "github.com/AbnerEarl/gost-x/ctx"
	"github.com/AbnerEarl/gost-x/registry"
	"github.com/go-gost/core/connector"
	md "github.com/go-gost/core/metadata"
)

func init() {
	registry.ConnectorRegistry().Register("sni", NewConnector)
}

type sniConnector struct {
	md      metadata
	options connector.Options
}

func NewConnector(opts ...connector.Option) connector.Connector {
	options := connector.Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return &sniConnector{
		options: options,
	}
}

func (c *sniConnector) Init(md md.Metadata) (err error) {
	return c.parseMetadata(md)
}

func (c *sniConnector) Connect(ctx context.Context, conn net.Conn, network, address string, opts ...connector.ConnectOption) (net.Conn, error) {
	log := c.options.Logger.WithFields(map[string]any{
		"remote":  conn.RemoteAddr().String(),
		"local":   conn.LocalAddr().String(),
		"network": network,
		"address": address,
		"sid":     string(ctxvalue.SidFromContext(ctx)),
	})
	log.Debugf("connect %s/%s", address, network)

	return &sniClientConn{Conn: conn, host: c.md.host}, nil
}

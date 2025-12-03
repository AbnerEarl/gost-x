package v4

import (
	"time"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	connectTimeout time.Duration
	disable4a      bool
}

func (c *socks4Connector) parseMetadata(md mdata.Metadata) (err error) {
	const (
		connectTimeout = "timeout"
		disable4a      = "disable4a"
	)

	c.md.connectTimeout = mdutil.GetDuration(md, connectTimeout)
	c.md.disable4a = mdutil.GetBool(md, disable4a)

	return
}

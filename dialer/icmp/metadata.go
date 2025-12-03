package quic

import (
	"time"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	keepAlivePeriod  time.Duration
	maxIdleTimeout   time.Duration
	handshakeTimeout time.Duration
}

func (d *icmpDialer) parseMetadata(md mdata.Metadata) (err error) {
	if mdutil.GetBool(md, "keepalive") {
		d.md.keepAlivePeriod = mdutil.GetDuration(md, "ttl")
		if d.md.keepAlivePeriod <= 0 {
			d.md.keepAlivePeriod = 10 * time.Second
		}
	}
	d.md.handshakeTimeout = mdutil.GetDuration(md, "handshakeTimeout")
	d.md.maxIdleTimeout = mdutil.GetDuration(md, "maxIdleTimeout")

	return
}

package tun

import (
	"math"
	"time"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

const (
	defaultKeepAlivePeriod = 10 * time.Second
	MaxMessageSize         = math.MaxUint16
)

type metadata struct {
	keepAlivePeriod time.Duration
	passphrase      string
	p2p             bool
}

func (h *tunHandler) parseMetadata(md mdata.Metadata) (err error) {
	if mdutil.GetBool(md, "tun.keepalive", "keepalive") {
		h.md.keepAlivePeriod = mdutil.GetDuration(md, "tun.ttl", "ttl")
		if h.md.keepAlivePeriod <= 0 {
			h.md.keepAlivePeriod = defaultKeepAlivePeriod
		}
	}

	h.md.passphrase = mdutil.GetString(md, "tun.token", "token", "passphrase")
	h.md.p2p = mdutil.GetBool(md, "tun.p2p", "p2p")
	return
}

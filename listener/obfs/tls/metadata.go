package tls

import (
	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	md "github.com/go-gost/core/metadata"
)

type metadata struct {
	mptcp bool
}

func (l *obfsListener) parseMetadata(md md.Metadata) (err error) {
	l.md.mptcp = mdutil.GetBool(md, "mptcp")

	return
}

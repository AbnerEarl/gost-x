package tcp

import (
	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	tproxy bool
	mptcp  bool
}

func (l *redirectListener) parseMetadata(md mdata.Metadata) (err error) {
	l.md.tproxy = mdutil.GetBool(md, "tproxy")
	l.md.mptcp = mdutil.GetBool(md, "mptcp")

	return
}

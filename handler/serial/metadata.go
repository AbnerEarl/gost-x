package serial

import (
	"time"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	timeout time.Duration
}

func (h *serialHandler) parseMetadata(md mdata.Metadata) (err error) {
	h.md.timeout = mdutil.GetDuration(md, "timeout", "serial.timeout")
	return
}

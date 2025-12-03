package serial

import (
	"time"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	md "github.com/go-gost/core/metadata"
)

type metadata struct {
	timeout time.Duration
}

func (l *serialListener) parseMetadata(md md.Metadata) (err error) {
	l.md.timeout = mdutil.GetDuration(md, "timeout", "serial.timeout", "listener.serial.timeout")

	return
}

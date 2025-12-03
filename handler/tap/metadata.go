package tap

import (
	"math"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

const (
	MaxMessageSize = math.MaxUint16
)

type metadata struct {
	key string
}

func (h *tapHandler) parseMetadata(md mdata.Metadata) (err error) {
	const (
		key = "key"
	)

	h.md.key = mdutil.GetString(md, key)
	return
}

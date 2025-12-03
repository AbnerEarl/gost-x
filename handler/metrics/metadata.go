package metrics

import (
	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

const (
	DefaultPath = "/metrics"
)

type metadata struct {
	path string
}

func (h *metricsHandler) parseMetadata(md mdata.Metadata) (err error) {
	h.md.path = mdutil.GetString(md, "metrics.path", "path")
	if h.md.path == "" {
		h.md.path = DefaultPath
	}
	return
}

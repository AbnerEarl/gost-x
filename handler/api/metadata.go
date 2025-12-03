package api

import (
	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	accesslog  bool
	pathPrefix string
}

func (h *apiHandler) parseMetadata(md mdata.Metadata) (err error) {
	h.md.accesslog = mdutil.GetBool(md, "api.accessLog", "accessLog")
	h.md.pathPrefix = mdutil.GetString(md, "api.pathPrefix", "pathPrefix")
	return
}

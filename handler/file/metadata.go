package file

import (
	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	dir string
}

func (h *fileHandler) parseMetadata(md mdata.Metadata) (err error) {
	h.md.dir = mdutil.GetString(md, "file.dir", "dir")
	return
}

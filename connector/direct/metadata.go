package direct

import (
	"strings"

	mdutil "github.com/AbnerEarl/gost-x/metadata/util"
	mdata "github.com/go-gost/core/metadata"
)

type metadata struct {
	action string
}

func (c *directConnector) parseMetadata(md mdata.Metadata) (err error) {
	c.md.action = strings.ToLower(mdutil.GetString(md, "action"))
	return
}

package api

import (
	"net/http"

	"github.com/AbnerEarl/gost-x/config"
	"github.com/AbnerEarl/gost-x/config/loader"
	"github.com/AbnerEarl/gost-x/config/parsing/parser"
	"github.com/AbnerEarl/gost-x/registry"
	"github.com/gin-gonic/gin"
)

// swagger:parameters reloadConfigRequest
type reloadConfigRequest struct{}

// successful operation.
// swagger:response reloadConfigResponse
type reloadConfigResponse struct {
	Data Response
}

func reloadConfig(ctx *gin.Context) {
	// swagger:route POST /config/reload Reload reloadConfigRequest
	//
	// Hot reload config.
	//
	//     Security:
	//       basicAuth: []
	//
	//     Responses:
	//       200: reloadConfigResponse

	cfg, err := parser.Parse()
	if err != nil {
		writeError(ctx, NewError(http.StatusBadRequest, ErrCodeInvalid, err.Error()))
		return
	}

	config.Set(cfg)

	if err := loader.Load(cfg); err != nil {
		writeError(ctx, NewError(http.StatusBadRequest, ErrCodeInvalid, err.Error()))
		return
	}

	for _, svc := range registry.ServiceRegistry().GetAll() {
		svc := svc
		go func() {
			svc.Serve()
		}()
	}

	ctx.JSON(http.StatusOK, Response{
		Msg: "OK",
	})
}

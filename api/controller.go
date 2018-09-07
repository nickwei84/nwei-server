package api

import (
	"net/http"

	"github.com/n4wei/nwei-server/api/healthcheck"
	"github.com/n4wei/nwei-server/api/weight"
	"github.com/n4wei/nwei-server/lib/logger"
)

type controller struct {
	router *http.ServeMux
	logger logger.Logger
}

func NewController(logger logger.Logger) *controller {
	router := http.NewServeMux()
	router.Handle("/healthcheck", chain(healthcheck.Handler, WithLogging(logger)))
	router.Handle("/weight", chain(weight.Handler, WithLogging(logger)))

	return &controller{
		router: router,
		logger: logger,
	}
}

func (c *controller) Handler() http.Handler {
	return c.router
}
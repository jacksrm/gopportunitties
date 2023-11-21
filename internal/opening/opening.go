package opening

import (
	"github.com/jacksrm/gopportunitties/internal/opening/controller"
	"github.com/jacksrm/gopportunitties/internal/opening/repository/sqlite"
	"github.com/jacksrm/gopportunitties/internal/opening/service"
)

type Module struct {
	Service    *service.Service
	Controller *controller.Controller
}

func New() *Module {
	repo := sqlite.New()
	service := service.New(repo)
	controller := controller.New(service)

	return &Module{service, controller}
}

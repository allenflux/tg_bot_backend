package management

import "tg_bot_backend/internal/service"

type sManagement struct{}

func New() *sManagement {
	return &sManagement{}
}

func init() {
	service.RegisterManagement(New())
}

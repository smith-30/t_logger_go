package sample

import (
	"github.com/smith-30/t_logger/zap_logger/handler"
	"go.uber.org/zap"
)

type (
	sampleHandler struct {
		logger *zap.Logger
	}

	Config struct {
		Logger *zap.Logger
	}
)

func NewHandler(c Config) handler.Handler {
	return &sampleHandler{
		logger: c.Logger,
	}
}

func (h *sampleHandler) Do(param interface{}) error {
	switch param.(type) {
	case int:
		h.logger.Info("Did Int", zap.Int("param", param.(int)))
	case string:
		h.logger.Info("Did String", zap.String("param", param.(string)))
	default:
		h.logger.Info("Did Any", zap.Any("param", param))
	}
	return nil
}

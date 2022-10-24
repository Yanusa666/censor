package censor

import (
	"github.com/rs/zerolog"
	"sf-censor/internal/config"
)

type Censor struct {
	cfg *config.Config
	lgr zerolog.Logger
}

func NewCensor(cfg *config.Config, lgr zerolog.Logger) *Censor {
	c := &Censor{
		cfg: cfg,
		lgr: lgr,
	}

	return c
}

func (c *Censor) Shutdown() error {
	return nil
}

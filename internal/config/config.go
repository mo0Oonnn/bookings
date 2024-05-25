package config

import (
	"github.com/CloudyKit/jet"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache *jet.Set
	InProduction  bool
	Session       *scs.SessionManager
}

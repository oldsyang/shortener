package global

import (
	ut "github.com/go-playground/universal-translator"
	"shortener/config"
)

var (
	ServerConfig    *config.ServerConfig = &config.ServerConfig{}
	Trans           ut.Translator
)

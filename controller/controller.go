package controller

import (
	"github.com/lazypanda2004/VLC_controller/config"
)

type Controller interface {
	Connect(connCfg *ConnCfg)
}




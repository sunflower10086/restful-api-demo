package impl

import (
	"log"
	"os"

	"github.com/sunflower10086/restful-api-demo/apps/host"
)

var _ host.Service = &HostServiceImpl{}

type HostServiceImpl struct {
	l *log.Logger
}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l: log.New(os.Stderr, "  [Host] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

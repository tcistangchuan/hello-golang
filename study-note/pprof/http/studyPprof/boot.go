package studyPprof

import (
	"log"
	_ "net/http/pprof"
)

func BootPprof() {
	log.Println("start pprof")
}

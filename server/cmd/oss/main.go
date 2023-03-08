package main

import (
	oss "github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen/oss/ossservice"
	"log"
)

func main() {
	svr := oss.NewServer(new(OssServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

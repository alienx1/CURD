package main

import "api/src/controller"

func main() {
	r := controller.ServerRouter()
	r.Run()
}

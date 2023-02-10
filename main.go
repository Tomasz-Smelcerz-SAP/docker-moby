package main

import (
	"fmt"

	docker "github.com/docker/docker/pkg/tarsum"
	moby "github.com/moby/moby/pkg/tarsum"
)

func debug(obj interface{}) string {
	return fmt.Sprintf("%T %v %#v", obj, obj, obj)
}
func main() {
	fmt.Println("start")
	var dv docker.Version = 1
	var mv moby.Version = 2

	fmt.Println(debug(dv))
	fmt.Println(debug(mv))

	fmt.Println("done")
}

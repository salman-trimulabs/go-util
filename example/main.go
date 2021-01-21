package main

import (
	utils "github.com/salman-trimulabs/go-util"
)

func main(){
	x :=  utils.Init{
		"docker://quay.io/buildah/stable", "~/Documents/test/",
	}
	x.CopyImageFromDockerRepository()
}
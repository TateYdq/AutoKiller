package main

import (
	"fmt"
	"os"
	"github.com/TateYdq/AutoKiller/util/tool"
	"os/exec"
	"github.com/gin-gonic/gin"
)


func Test(){

	res := exec.Command("ifconfig")
	err := res.Run()
	if err != nil {
		fmt.Errorf("error:%s\n",err)
		os.Exit(1)
	}
	out := res.Stdout
	fmt.Println(out)
}


func main()  {
	r := gin.Default()
	r.GET("/ping",util.Pong)
}


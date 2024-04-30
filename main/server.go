package main

import (
	"fmt"

	"github.com/koki-mekonnen/Discordbot-Go-Echo/bot"
	"github.com/koki-mekonnen/Discordbot-Go-Echo/config"
	"github.com/labstack/echo"
)


func main(){

	e:=echo.New()
	err:=config.ReadConfig()

	if err!=nil{
		fmt.Println(err.Error())
		return
	}

   bot.Start()

	e.Start(":8080")

}
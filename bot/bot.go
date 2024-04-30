package bot

import (
	"github.com/koki-mekonnen/Discordbot-Go-Echo/config"

	"fmt"

	"github.com/bwmarrin/discordgo"
)


var BotID string
var goBot *discordgo.Session


func Start(){

	goBot,err :=discordgo.New("Bot "+ config.Token)

	
	if err!=nil{
	fmt.Println((err.Error()))
		return 
	}

	u,err :=goBot.User("@me")

	
	if err!=nil{
	fmt.Println((err.Error()))
		return 
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err=goBot.Open()

	
	if err!=nil{
	fmt.Println((err.Error()))
		return 
	}
	

	fmt.Println("bot is running ")

}


func messageHandler(s *discordgo.Session,m *discordgo.MessageCreate){
	if m.Author.ID ==BotID{
		return
	}

	if m.Content =="ping"{
		_,_=s.ChannelMessageSend(m.ChannelID,"PONG")
	}





}
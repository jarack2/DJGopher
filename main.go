package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const token string = "NzcwMDAyMzExODc4OTM0NTI4.X5XOiQ.Z9F3_0y55l_VScYv7qx_zbV38rg"
var BotID string

func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	err = dg.Open()

  if err != nil {
	  fmt.Println(err.Error())
    return
  }

	fmt.Println("Bot is running!")
}

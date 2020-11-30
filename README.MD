# DJ Gopher <img src="https://pbs.twimg.com/media/DTYVjS-X4AEgy7p.png" width="100" height="100">
----------
## Overview
The GO language is a language created by google in order to fullfill scalability, space, and run time needs.

DJ Gopher is a discord bot written in GO that has a couple of different functions which include:

* Games
  * Hangman (1+ players)
  * Connect Four (2 players)
  * Trivia (2-6 players)
* Music
  * Pop
  * Punk
  * Rock 
  * Sound Effects

The bot is only interactable in Discord, a chat client, and must be running on a server or computer in order to connect to Discord and function properly. 

Once the program is running, no further input besides that in Discord chatbox is needed. 


## Installing Environment:

To get started with GO, head over to the official downloads page and click on the OS being used. https://golang.org/doc/install

NOTE: For the purposes of the README the Windows version will be used.

Verify GO has been correctly installed with the command.

 $ go version

In order to write in GO, a text editor is needed. Visual Studio Code is popular editor with plugin support for GO. Although most editors will suffice.

Download Visual Studio Code and then install the GO plugin
 https://code.visualstudio.com/download
 https://marketplace.visualstudio.com/items?itemName=golang.Go


Finally, the latest version of Git should be installed. Git is a way to push, pull, and merge code, mostly on Github. This is how many packages will be imported into the project and how the project itself will be published.

 https://git-scm.com/download/win


All the necessary components are installed to effectively work with GO lang in the Windows environment

## Adding Bot to Server:

A discord account is needed to proceed. Create an account on https://discord.com/ then proceed.

DJGopher is an application that must be added then run on a Discord server. To add the bot first create a server. 

<img src="https://media.discordapp.net/attachments/772600665301581855/782809967043411978/addServer.png" width="1200" height="600">

Then add the bot to the server by clicking this link 
https://discord.com/api/oauth2/authorize?client_id=770002311878934528&permissions=1849163584&scope=bot

<img src="https://media.discordapp.net/attachments/772600665301581855/782817670331498516/addBot.png">

Now that the bot is added to the server, we can run it to play games or music.

## Running The Bot:

To start, clone the repository that includes the bot. This is done by using the command 

$ git clone https://github.com/JaredRackley/DJGopher.git

After this is done run the following command

$ go get github.com/bwmarrin/discordgo

The Discord bot is now ready to run. Use the command

$ go build  

This will create an executable DJGopher.exe that will run the bot when executed or you can use the command below to run it as well

$ go run main.go


## Usage:

The following commands will invoke the bot
           
![User Commands](https://media.discordapp.net/attachments/765801734547177505/782807971377446932/Usage.png)




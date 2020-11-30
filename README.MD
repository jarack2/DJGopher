# DJ Gopher <img src="https://pbs.twimg.com/media/DTYVjS-X4AEgy7p.png" width="100" height="100">
----------
## Overview
The GO language is a language created by google in order to fullfill scalability, space, and run time needs.

DJ Gopher is a discord bot written in GO that has a couple of different functions which include:

* Games/
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


## Configuring The Bot:

To start, clone the repository that includes the bot. This is done by using the command 

$ git clone github.com/JaredRackley/DJGopher/

After this is done run the following command

$ go get github.com/bwmarrin/discordgo

The Discord bot is now ready to run. Use the command

$ go build main.go  

This will create an executable DJGopher.exe that will run the bot when executed

## Adding Bot to Your Server:

To add a bot to a server you created and own use this link https://discord.com/api/oauth2/authorize?client_id=770002311878934528&permissions=1849163584&scope=bot


## Usage:

The following commands will invoke the bot
 
![User Commands](https://github.com/JaredRackley/DJGopher/blob/main/UserInput.png)




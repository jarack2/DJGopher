DJ Gopher
----------

To start, clone the repository.

Then you must first run the following command: go get github.com/bwmarrin/discordgo

To convert from mp3 to dca use:

ffmpeg -i test.mp3 -f s16le -ar 48000 -ac 2 pipe:1 | dca > test.dca

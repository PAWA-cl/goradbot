
# About

**GoRadBot**, a radio bot player for Discord. Made by Patricio Pérez (me) for PAWA.cl

# Why a bot?

Well, right now there a so many bots used for music player in Discord, but i found the following problems:

- Some of the players don't have support for external links, because of security measures (which are fine, i can understand it).
- Others only offer this as a premium feature.
- Others are a little laggy, or have a latency, causing that the music sounds chopped, or incomplete.
- Others, offer radio as part of a very large suite with other unnecessary commands.
- And finally, most of them offer no open-source alternatve.

## But there are options!

Yeah! There is a popular option, [Fredboat](https://github.com/Frederikam/FredBoat), that is implemented in Java. The problem: I need to run this in a **very** cheap alternative (like, a raspberry), so the memory usage is crucial.
 
## And then, Gopher speaks

Well, since i'm interested to develop this as quick as possible, the most reasonable option (for me) was Golang. There are many libraries for Golang, that there is already one that supports Discord and also mp3 plays!

# Features

This initial release has only the basic support:

- A Helper command.
- Bot Token configuration through a YAML file.
- Embed templates for the bot, written in YAML for convenience.
- Obviously, a play url/stop feature.

# Requirements

For the current implementation, you need to install `ffmpeg`. This is used for the mp3 conversion to dca.

## For developing

If you want to test it in your workspace, then you also need to install the following go packages:

    go get -u github.com/bwmarrin/discordgo
    go get -u github.com/ghodss/yaml
    go get -u github.com/jonas747/dca

# Improvements

Since this is an initial release, there are many things to solve and document, so i'll keep improving during the following weeks. But, for now, here are my TODO's

 - [ ] Add Documentation in Wiki
 - [ ] Polish templates
 - [ ] Write tutorial for initial Discord Users
 - [ ] Add Session support
 - [ ] Make a docker file

## Additional Goals
 - [ ] Launch a bot officially in top.gg
 - [ ] ?
 - [ ] Profit

# Support the maintainer
Finally, if you think that my effort deserves a prize, or if you want to help to keep me motivated, i'll be always grateful if you can support the project. Please check the **Sponsor** button! Thanks in advance.

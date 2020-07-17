# Slack-ProfileApp

[![Build Status](https://travis-ci.org/jasonbronson/slack-profileapp.svg?branch=master)](https://travis-ci.org/jasonbronson/slack-profileapp)

Slack-Profileapp is a Golang based app which can be set to run on a crontab so that you can automatically change your slack profile picture every X interval. 



# Features!

  - Reads from a directory list of images provided
  - Randomly pics off an image to use
  - Token can be passed to the app to allow more than one slack team to be used. (crontab.sh file)


### Installation

To build the Slack-Profileapp just download repo then type in "go build". The binary file profileapp will appear in the root directory. 

Then run it with the following:

``` ./profileapp -token=123 -path=./pics```

Make sure you get a token from slack here by creating a slack app on your slack team. 
https://api.slack.com/apps





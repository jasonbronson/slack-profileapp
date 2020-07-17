#!/bin/bash
#eval $(egrep -v '^#' .env | xargs) myfunction
export $(egrep -v '^#' .env | xargs)
#echo ${token}
#echo ${path}

/Users/jasonbronson/Downloads/profileapp/profileapp -token=${token} -path=${path} 
/Users/jasonbronson/Downloads/profileapp/profileapp -token=${token2} -path=${path}

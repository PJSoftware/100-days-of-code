#!bash

cd cmd/TweetCommit
go build
ls -l TweetCommit.exe
cd ../..

cp cmd/TweetCommit/TweetCommit.exe ..

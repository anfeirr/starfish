#! /bin/sh

kill -9 $(pgrep webserver)

cd ~/starfish/
git pull https://github.com/anfeirr/starfish.git
./webserver &
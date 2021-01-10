#! /bin/bash
export GOPATH="${HOME}/go_projects/podcaster"
nohup /usr/share/code/code --no-sandbox --unity-launch ${GOPATH}/src > /dev/null &
nohup godoc -http=:6060 > /dev/null &


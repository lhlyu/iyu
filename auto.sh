#!/usr/bin/env bash

name="iyu"

case $1 in
        [gG]*)
                echo "git push start"
                git pull
                git add .
                git commit -m "update"
                git push
                echo "git push end"
                exit
                ;;
        [bB]*)
                case $2 in
                    linux)
                        echo "build for linux"
                        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${name}" main.go
                        exit
                        ;;
                    mac)
                        echo "build for mac"
                        CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o "${name}-mac" main.go
                        exit
                        ;;
                    win)
                        echo "build for mac"
                        CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o "${name}.exe" main.go
                        exit
                        ;;
                    *)
                        echo "build for current system"
                        go build main.go
                        exit
                        ;;
                esac
                exit
                ;;
        [dD]*)
            echo "docker build and run"
            docker build -t $name .
            docker run -itd --name=$name -p 9876:8080 $name
            exit
            ;;
esac

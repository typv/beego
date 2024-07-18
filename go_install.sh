#!/bin/sh

echo "========= Install CLI tools ========="

check_tool() {
    if ! command -v $1 &> /dev/null
    then
        echo "$1 not found"
    else
        echo "$1 has been installed at $(command -v $1)"
    fi
}

echo "Install Bee tools..."
go install github.com/beego/bee/v2@latest
check_tool bee


echo "========= Finish CLI tools installation ========="
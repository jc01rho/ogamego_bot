#!/usr/bin/env bash
#
# gh-dl-release! It works!
# ./updateApp.sh latest ogamebot
# 
# This script downloads an asset from latest or specific Github release of a
# private repo. Feel free to extract more of the variables into command line
# parameters.
#
# PREREQUISITES
#
# curl, wget, jq
#
# USAGE
#
# Set all the variables inside the script, make sure you chmod +x it, then
# to download specific version to my_app.tar.gz:
#
#     gh-dl-release 2.1.1 my_app.tar.gz
#
# to download latest version:
#
#     gh-dl-release latest latest.tar.gz
#
# If your version/tag doesn't match, the script will exit with error.

TOKEN="cfa3e998b56c0551011e308f194091f799cd9027"
REPO="jc01rho/ogamego_bot"
FILE="ogame-bot-linux-amd64"      # the name of your release asset file, e.g. build.tar.gz
VERSION="latest"                       # tag name or the word "latest"
GITHUB="https://api.github.com"

alias errcho='>&2 echo'

wget https://github.com/jc01rho/ogamego_bot/releases/latest/download/ogame-bot-linux-amd64 -O $2
chmod +x $2

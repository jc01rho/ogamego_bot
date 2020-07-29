#!/usr/bin/env bash

echo \"$1\"> universe
echo $2 > username
echo $3 > password
echo $4 > language
#lang en
#port 27015
#--api-new-hostname=http://304.198.117.29:27015
echo $5 > hostnamevar




if pgrep -x "ogamebot" > /dev/null
then
    echo "test"
else
    git pull
    ./updateApp.sh latest ogamebotNew;
    mv ogamebotNew ogamebot;
    universe=$(<universe)
    username=$(<username)
    password=$(<password)
    language=$(<language)
    hostnamevar=$(<hostnamevar)
    echo nohup ./ogamebot --universe=${universe} --username=${username} --password=${password} --language=${language} --port=27015 --host=0.0.0.0 --api-new-hostname=${hostnamevar} &
    sleep 2
    nohup ./ogamebot --universe=${universe} --username=${username} --password=${password} --language=${language} --port=27015 --host=0.0.0.0 --api-new-hostname=${hostnamevar} &
fi


#    jc85rho@instance-1:~/git/ogamego_bot$ nohup ./ogamebot --universe="Universe 1" --username=pku0369@naver.com --password=wjdals1234! --language=en --port=27015 --host=0.0.0.0 --api-new-hostname=http://104.198.117.29:27015 &
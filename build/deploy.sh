#!/bin/bash

chmod +x ./build/env.sh
chmod +x ./build/build.sh

./build/env.sh

rm -rf ./bin/template
cp -r ./template ./bin

./build/build.sh


PROCESSNO=`ps -ef | grep -v grep | grep ./app | awk '{print $2}'`
kill -9 ${PROCESSNO}

cd ./bin
nohup ./app > nohup.out 2>&1 &

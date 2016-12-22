#!/bin/bash

mysql -u workshopadm -h 172.16.150.199 -P 13306 -e 'drop database if exists zeroworkshop;create database zeroworkshop;' --password=workshopadm
mysql -u workshopadm -h 172.16.150.199 -P 13306 zeroworkshop < backup_zeroworkshop.sql --password=workshopadm

export PATH="$PATH:$PWD/bin"
cd robot
robot -v URL:http://localhost:8888/somkiatbank/account/ .

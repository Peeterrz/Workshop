#!/bin/bash

export PATH="$PATH:$PWD/bin"
cd robot
robot -v URL:http://localhost:8888/somkiatbank/account/ .

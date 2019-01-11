#!/bin/bash

curl --Location -o /tmp/scenario.txt https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.txt 
curl --Location -o /tmp/timekeeper.mac https://github.com/tkhr-sait/timekeeper/raw/master/bin/timekeeper.mac 
chmod +x /tmp/timekeeper.mac
/tmp/timekeeper.mac /tmp/scenario.txt

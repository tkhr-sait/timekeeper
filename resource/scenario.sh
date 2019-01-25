#!/bin/bash

curl --Location -o /tmp/scenario.txt https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.txt 
curl --Location -o /tmp/recog.swift https://github.com/tkhr-sait/timekeeper/raw/master/resource/recog.swift 
curl --Location -o /tmp/timekeeper.mac https://github.com/tkhr-sait/timekeeper/raw/master/bin/timekeeper.mac 
chmod +x /tmp/timekeeper.mac
chmod +x /tmp/recog.swift
/tmp/timekeeper.mac /tmp/scenario.txt

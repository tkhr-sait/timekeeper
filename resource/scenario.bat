@echo off
rem win10 or later only?
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.txt %TEMP%\scenario.txt
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/resource/recog.wsf %TEMP%\recog.wsf
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/bin/timekeeper.exe %TEMP%\timekeeper.exe

%TEMP%\timekeeper.exe %TEMP%\scenario.txt

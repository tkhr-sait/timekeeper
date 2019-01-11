@echo off
rem たぶんwindows10じゃないとだめ
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.txt %TEMP%\scenario.txt
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/bin/timekeeper.exe %TEMP%\timekeeper.exe

%TEMP%\timekeeper.exe %TEMP%\scenario.txt

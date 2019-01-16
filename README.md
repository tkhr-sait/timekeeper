# timekeeper

機械的に司会進行するツール

* mac

```
curl --Location -o /tmp/scenario.sh https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.sh && bash /tmp/scenario.sh
```

* windows command line

```
bitsadmin /TRANSFER htmlget https://github.com/tkhr-sait/timekeeper/raw/master/resource/scenario.bat %TEMP%\scenario.bat && %TEMP%\scenario.bat
```

* 入力テキストファイル

|command|動作|
|-------|--------|
|say,[message]|指定されたmessageを喋ります。|
|timer,[message],[minutes]|タイマー（分）です。's(Enter)','S(Enter)','skip(Enter)','Skip(Enter)'でタイマースキップ。'数値(Enter)'で時間延長します。|
|wait,[message]|指定されたmessageを喋ったあと、'(Enter)'が押されるまで待ちます。|
|open,[url]|urlを開きます。|

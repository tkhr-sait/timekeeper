# timekeeper

機械的に司会進行するツール

osの音声認識を利用するので、事前設定が必要となります。

|os|設定|
|-|-|
|mac|拡張音声入力を使用　https://support.apple.com/ja-jp/HT202584|
|windows|特になし？|

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
|timer,[message],[minutes]|タイマー（分）です。's(Enter)','S(Enter)','skip(Enter)','Skip(Enter)'もしくは"がいどさんつぎ"でタイマースキップ。'数値(Enter)'もしくは"がいどさんまって"で時間延長します。|
|wait,[message]|指定されたmessageを喋ったあと、'(Enter)'が押されるまで待ちます。"がいどさんつぎ"でも進みます。|
|open,[url]|urlを開きます。|

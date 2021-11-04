# golang-logging-library

golang-logging-library は Go ランタイム の マイクロサービス の ログ を出力する際に、ログの json フォーマットを統一するための Go ライブラリです。

## 動作環境

動作には以下の環境であることを前提とします。

・OS: Linux OS  
・CPU: ARM/AMD/Intel  
・Kubernetes  
・AION のリソース

## 利用方法

本リポジトリをインストールしてください。

```sh
go get -v github.com/latonaio/golang-logging-library/logger
```

各マイクロサービスのソース内に以下を配置してください。

```go
import "github.com/latonaio/golang-logging-library/logger"
```

インスタンスの作成は下記のように実行します。

```go
l := logger.NewLogger()
```

出力の形式

- log.Fatal(msg, tag) 
- log.Error(msg, tag)
- log.Info (msg, tag)
- log.Debug(msg, tag)
 
パラメーター

- msg: 文字列型、出力したい内容を指定する
- tag: Interface型、文字列を渡した場合のみ、出力タグをつけることが出来ます。指定しない場合はnilを渡すことを推奨します。

ログ出力例は以下の通りです。

```go
// tagを付けない場合
log.Debug("test", nil)
{"message":"test","level":"DEBUG","cursor":"/xxxxxx/.go#L11","time":"2021-11-02T19:39:19.404655+09:00","tag":null}

// tagをつける場合
log.Debug("test", "add tags")
{"message":"test","level":"DEBUG","cursor":"/xxxxxx/.go#L11","time":"2021-11-02T19:39:19.404655+09:00","tag":add tags}
```

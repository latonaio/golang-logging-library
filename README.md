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

インスタンスの作成は下記のように実行し、出力ログレベルの指定を引数に渡します。

```go
l := logger.NewLogger(4)
```

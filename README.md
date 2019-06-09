[![Go Report Card](https://goreportcard.com/badge/github.com/ebiebievidence/crowns)](https://goreportcard.com/report/github.com/ebiebievidence/crowns)

# なにこれ
研修用

# 構成

なんちゃってClean Architecture + なんちゃってCQRS。

```plain
├── app
│   ├── domain          - Domain層
│   │   ├── command     - CQRSにおけるCommand
│   │   ├── query       - CQRSにおけるQuery
│   │   ├── request     - HTTPリクエスト
│   │   └── response    - HTTPレスポンス
│   ├── handler         - HTTPのInterfaceであり、Controllerと呼ばれることもある
│   ├── repository      - RDBのInterface
│   ├── server
│   │   ├── router.go   - httprouterと組み合わせて適切なHandlerを設定する
│   │   └── server.go   - SimpleServer structが入っている
│   └── usecase         - Handlerからの呼び出しに応じて、Repositoryを呼び出す
├── config              - 各種の設定を書く
├── main.go             - SimpleServerを起動する
└── migration           - マイグレーション
    ├── down.sql        - テーブルを消す
    └── up.sql          - テーブルを作る
```

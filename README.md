# Clean Architecture Go Sample
Go言語でClean Architectureを実現するための勉強用サンプルコードです。  
Clean Architectureを実現するためにどういったことが必要となるのか自分なりにコメントを付与しながら作成した勉强用リポジトリです。　 

(自分が見返してもわかるように、作成しているので見づらい点はご容赦ください。)

## 使用方法
1. ソースコードをCloneする.  
   Local環境へソースコードをCloneしてください。
   ```sh
   git clone git@github.com/atEaE/clean-architecture-go-sample
   ```

## テスト実行方法
### 全テストを実行する場合
全テストケースの実行を行いたい場合は、下記のコマンドを実行する。  
テストが正常に終了した場合は、左記に「ok」と表示される。  
「.go」ファイルが存在するが、「xx_test.go」テストファイルが存在しない場合、左記に「?」が表示される。
```sh
$ go test ./...
?       github.com/atEaE/clean-architecture-go-sample   [no test files]
ok      github.com/atEaE/clean-architecture-go-sample/domain    0.462s
```

## 開発環境Tips
### goenvを使用したMacで開発
タイトルの組み合わせて開発を行う場合、言語サーバーなどの指定が必要なる場合がある。  
(自身の環境でしか検証していないため、他の環境でも発生するかは真偽不明。)
`.vscode` フォルダに `settings.json` ファイルを配置し、下記の設定を追加する。  
```json
{
    "go.useLanguageServer": true,
    "[go]": {
        "editor.snippetSuggestions": "none",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        }
    },
    "go.toolsEnvVars": {
        "GO111MODULE": "on",
     },
    "go.goroot": "/Users/[username]/.goenv/versions/1.13.0"
}

```


## 参考・引用
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)  
  こちらの記事を参考にさせていただきました。ありがとうございます。

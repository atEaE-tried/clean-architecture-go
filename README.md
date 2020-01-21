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
### 簡易的に全テストを実行する場合
簡易的に全テストケースの実行を行いたい場合は、下記のコマンドを実行する。  
テストが正常に終了した場合は、左記に「ok」と表示される。  
「.go」ファイルが存在するが、「xx_test.go」テストファイルが存在しない場合、左記に「?」が表示される。
```sh
$ go test ./...
?       github.com/atEaE/clean-architecture-go-sample   [no test files]
ok      github.com/atEaE/clean-architecture-go-sample/domain    0.462s
```

## 参考・引用
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)  
  こちらの記事を参考にさせていただきました。ありがとうございます。

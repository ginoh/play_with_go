## 開発環境構築

### 関連リンク

[go 公式](https://golang.org/)

### インストール
#### Windows
##### 本体
* 公式サイト提供のMSIインストーラを利用
* 公式サイト提供のzipファイルを利用

##### 環境変数
* path・・・環境変数にインストールしたGoフォルダ下のbinフォルダを追加する
* GOPATH・・・外部ライブラリが格納されているフォルダ。空のフォルダを作成して設定する。最初は `$HOME/go` あたりがよい

#### OSX
##### 本体
* 公式サイト提供のパッケージを利用してインストール
* homebrewを利用してインストール
```
$ brew install go
```

##### 環境変数
* PATH・・・インストールしたgoのbinフォルダをshellの設定ファイルなどで設定
* 追加する
* GOPATH・・・外部ライブラリが格納されているフォルダ。空のフォルダを作成して設定する。最初は `~/go` あたりがよい

#### Linux
##### 本体
* 公式サイト提供のtarボールを解凍、配置

##### 環境変数
* PATH・・・インストールしたgoのbinフォルダをshellの設定ファイルなどで設定
* GOPATH・・・外部ライブラリが格納されているフォルダ。空のフォルダを作成して設定する。最初は `~/go` あたりがよい

### 注意
GOPATHの設定については、複数したりとやり方があるようなので、
後ほど、検討する。

http://blog.wacul.co.jp/blog/2014/08/22/go/
http://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4
http://text.baldanders.info/golang/gopath-pollution/
http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/

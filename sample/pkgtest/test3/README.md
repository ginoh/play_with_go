### mockgenを利用したテストについて

#### パッケージ構成
```
├── app
│   └── app.go
└── main.go
```
### モックの作成

app.goのテストをモックを作成・利用してテストする

### mockgenパッケージのインストール
参考：https://github.com/golang/mock

```
go get github.com/golang/mock/mockgen
```

#### mockgen

モックは interface を利用して作成される

app.goの内容
```
package app

type Application interface {
        Name() string
}

type MyApp struct {
        name string
}

func NewMyApp(name string) *MyApp {
        return &MyApp{
                name: name,
        }
}

func (a *MyApp) Name() string {
        return a.name
}
```

mockの生成

mockgenの実行には`source`と`reflect`の2つのモードがある  
ここでは`source`モードを利用する
```
$ mockgen -source=app/app.go -destination app/mock_app.go
$ tree .
├── app
│   ├── app.go
│   └── mock_app.go
└── main.go
```

mock は以下の様な感じで利用できる
```
func TestXXXX(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := mock_app.NewMockApplicaion(ctrl)
	app.EXPECT().Name().Return("MockApp")
    ・
    ・
    ・
    // このMockを注入する形などで利用する
}
```
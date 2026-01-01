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

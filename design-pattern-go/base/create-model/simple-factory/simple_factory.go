package simple_factory

import "fmt"

// 简单工厂模式
func NewAPI() *API {
	return &API{}
}

type API struct {
}

func (a *API) Say() string {
	return fmt.Sprintf("say hello!")
}

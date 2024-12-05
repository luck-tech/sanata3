package container

import (
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway"
	"github.com/murasame29/go-httpserver-template/internal/driver"
	"github.com/murasame29/go-httpserver-template/internal/framework/jwts"
	"github.com/murasame29/go-httpserver-template/internal/router"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/murasame29/go-httpserver-template/internal/usecase/interactor"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
	"github.com/uptrace/bun"
	"go.uber.org/dig"
)

var container *dig.Container

type provideArg struct {
	constructor any
	opts        []dig.ProvideOption
}

// NewContainer は、digを用いて依存注入を行う
func NewContainer() error {
	var noOpts []dig.ProvideOption
	container = dig.New()

	args := []provideArg{
		{constructor: router.NewEcho, opts: noOpts},
		{constructor: jwts.NewJWTMaker, opts: noOpts},

		{constructor: service.NewGitHub, opts: noOpts},
		{constructor: service.NewSession, opts: noOpts},
		{constructor: service.NewSkill, opts: noOpts},
		{constructor: service.NewUsedSkill, opts: noOpts},
		{constructor: service.NewUser, opts: noOpts},

		{constructor: interactor.NewLogin, opts: noOpts},
		{constructor: interactor.NewUser, opts: noOpts},

		{constructor: router.NewDI, opts: noOpts},
		{constructor: driver.NewDB, opts: noOpts},
		{constructor: driver.NewBun, opts: as[bun.IDB]()},
		{constructor: gateway.NewRepository, opts: as[dai.DataAccessInterface]()},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func as[T any]() []dig.ProvideOption {
	return []dig.ProvideOption{dig.As(new(T))}
}

// Invoke は、 *dig.ContainerのInvokeをwrapしてる関数
func Invoke[T any](opts ...dig.InvokeOption) (T, error) {
	var r T
	if err := container.Invoke(func(t T) error {
		r = t
		return nil
	}, opts...); err != nil {
		return r, err
	}
	return r, nil
}

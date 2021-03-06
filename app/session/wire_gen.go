// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package session

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/go-session/session"
)

// Injectors from injector.go:

func Build() (*session.Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	sessionConfig, cleanup3, err := Cfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup4, err := Provider(context, sessionConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return manager, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (*session.Manager, func(), error) {
	manager, cleanup, err := ProviderTest()
	if err != nil {
		return nil, nil, err
	}
	return manager, func() {
		cleanup()
	}, nil
}

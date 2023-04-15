// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/lunabrain-ai/lunabrain/pkg/store/cache"
	"github.com/lunabrain-ai/lunapipe/internal/config"
	"github.com/urfave/cli/v2"
)

// Injectors from wire.go:

func Wire(cacheConfig cache.Config) (*cli.App, error) {
	localCache, err := cache.NewLocalCache(cacheConfig)
	if err != nil {
		return nil, err
	}
	provider, err := config.NewConfigProvider(localCache)
	if err != nil {
		return nil, err
	}
	openAIConfig, err := config.NewOpenAIConfig(provider)
	if err != nil {
		return nil, err
	}
	openAIQAClient := NewOpenAIQAClient(openAIConfig)
	logConfig, err := config.NewLogConfig(provider)
	if err != nil {
		return nil, err
	}
	localConfigurator := config.NewConfigurator(localCache)
	app := NewCLI(openAIQAClient, logConfig, localConfigurator)
	return app, nil
}

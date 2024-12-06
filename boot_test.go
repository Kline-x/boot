package boot

import (
	"fmt"
	"testing"
)

type Config struct {
	AppKey    string
	AppSecret string
	Debug     bool
}

type LogStarter struct {
	BaseStarter[*Config]
}

func (l *LogStarter) Init(cfg *Config) {
	fmt.Println(cfg.AppKey, cfg.AppSecret, cfg.Debug)
}

func (l *LogStarter) Start(cfg *Config) {
	fmt.Println(132123123)
}

func (l *LogStarter) Stop(cfg *Config) {}

func (l *LogStarter) Name() string {
	return "LogStarter"
}

func TestApplication(t *testing.T) {
	config := &Config{
		AppKey:    "123",
		AppSecret: "456",
		Debug:     true,
	}

	app := NewApplication(config, WithDebugOption(true))
	app.Register(&LogStarter{})
	app.Start()

	defer func() {
		app.Stop()
	}()
}

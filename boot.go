package boot

import (
	"fmt"
	"log"
)

type Application[T any] struct {
	starters []Starter[T]
	cfg      T
	appCfg   *appCfg
}

type appCfg struct {
	debug bool
}

func NewApplication[T any](cfg T, opts ...Option) *Application[T] {
	ac := &appCfg{}

	for _, opt := range opts {
		opt(ac)
	}
	return &Application[T]{cfg: cfg, starters: make([]Starter[T], 0), appCfg: ac}
}

func (a *Application[T]) Register(starters ...Starter[T]) {
	a.starters = append(a.starters, starters...)
}

func (a *Application[T]) Start() {
	a.init()
	a.start()
}

func (a *Application[T]) Stop() {
	a.stop()
}

func (a *Application[T]) init() {
	for _, starter := range a.starters {
		a.log(fmt.Sprintf("init %v\n", starter.Name()))
		starter.Init(a.cfg)
	}
}

func (a *Application[T]) start() {
	for i, starter := range a.starters {
		a.log(fmt.Sprintf("start %v\n", starter.Name()))
		// 如果是最后一个，则直接启动
		if i+1 == len(a.starters) {
			starter.Start(a.cfg)
		} else {
			if starter.Blocking() {
				// 否则的话判断一下是不是会阻塞，如果会阻塞，就开协程跑
				go starter.Start(a.cfg)
			} else {
				starter.Start(a.cfg)
			}
		}
	}
}

func (a *Application[T]) stop() {
	for _, starter := range a.starters {
		starter.Stop(a.cfg)
		a.log(fmt.Sprintf("stop %v\n", starter.Name()))
	}
}

func (a *Application[T]) log(msg string) {
	if a.appCfg.debug {
		log.Print(msg)
	}
}

package boot

import "log"

type Application[T any] struct {
	starters []Starter[T]
	cfg      T
}

func NewApplication[T any](cfg T) *Application[T] {
	return &Application[T]{cfg: cfg, starters: make([]Starter[T], 0)}
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
		log.Printf("init %v\n", starter.Name())
		starter.Init(a.cfg)
	}
}

func (a *Application[T]) start() {
	for i, starter := range a.starters {
		log.Printf("start %v\n", starter.Name())
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
		log.Printf("stop %v\n", starter.Name())
	}
}

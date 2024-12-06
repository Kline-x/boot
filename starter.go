package boot

type Starter[T any] interface {
	Init(cfg T)
	Start(cfg T)
	Stop(cfg T)
	Blocking() bool // 是否会阻塞，starter会按顺序启动，阻塞的starter会开协程处理
	Name() string
}

type BaseStarter[T any] struct{}

func (b BaseStarter[T]) Init(cfg T) {}

func (b BaseStarter[T]) Start(cfg T) {}

func (b BaseStarter[T]) Stop(cfg T) {}

func (b BaseStarter[T]) Blocking() bool { return false }

func (b BaseStarter[T]) Name() string { return "" }

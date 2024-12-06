package boot

type Option func(o *appCfg)

func WithDebugOption(b bool) Option {
	return func(o *appCfg) {
		o.debug = b
	}
}

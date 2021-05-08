package mockgen

type GetPageOptions struct {
	Getter Getter
}

type Option func(*GetPageOptions)

func WithGetter(g Getter) Option {
	return func(options *GetPageOptions) {
		options.Getter = g
	}
}

package sefl_aes

type selfAes struct {
	key  string
	iv   string
	head int
}

type Option interface {
	apply(aes *selfAes)
}

type funcOption struct {
	f func(*selfAes)
}

func (fdo *funcOption) apply(do *selfAes) {
	fdo.f(do)
}

func newFuncOption(f func(*selfAes)) *funcOption {
	return &funcOption{
		f: f,
	}
}
func WithKey(k string) Option {
	return newFuncOption(func(aes *selfAes) {
		aes.key = k
	})
}
func WithHeader(h int) Option {
	return newFuncOption(func(aes *selfAes) {
		aes.head = h
	})
}

func NewAes(opts ...Option) *selfAes {

	obj := &selfAes{}
	for _, opt := range opts {
		opt.apply(obj)
	}
	return obj
}

func Encrypt(str string) string {

	return ""
}

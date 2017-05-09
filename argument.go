package guerrillamail

type Argument map[string]string

func NewArgument() Argument {
	args := make(Argument)

	return args
}

func (a Argument) Set(k, v string) {
	if k != "" && v != "" {
		a[k] = v
	}
}

func (a Argument) Get(k string) string {
	return a[k]
}

func (a Argument) Del(k string) {
	delete(a, k)
}

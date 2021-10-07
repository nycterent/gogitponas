package registry

type Callback interface {
	Send()
	Set(interface{})
}

type Registry struct {
	target []Callback
	names  []string
}

func (r *Registry) Register(name string, target Callback) {
	if r.inArray(name) {
		r.target = append(r.target, target)
	}
}

func (r Registry) Send(i interface{}) {
	for _, target := range r.target {
		target.Set(i)
		target.Send()
	}
}

func (r Registry) inArray(name string) bool {
	for _, v := range r.names {
		if v == name {
			return true
		}
	}
	return false
}

func New(names []string) *Registry {
	return &Registry{names: names}
}

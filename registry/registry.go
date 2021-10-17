package registry

// Callback defines an interface for target notifications
type Callback interface {
	Send(interface{})
}

// Registry holds all targets
type Registry struct {
	target []Callback
	names  []string
}

// Register registers a callback
func (r *Registry) Register(name string, target Callback) {
	if r.inArray(name) {
		r.target = append(r.target, target)
	}
}

// Send defines an interface for target Send
func (r Registry) Send(i interface{}) {
	for _, target := range r.target {
		target.Send(i)
	}
}

// inArray internal function to check if callback is in config
func (r Registry) inArray(name string) bool {
	for _, v := range r.names {
		if v == name {
			return true
		}
	}
	return false
}

// New constructor sets notification packages from the string array
func New(names []string) *Registry {
	return &Registry{names: names}
}

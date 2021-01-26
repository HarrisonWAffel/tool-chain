package types

type Command struct {
	Name    string
	Handler func()
}

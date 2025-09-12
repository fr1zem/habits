package CLI

import (
	"flag"
	"log"
	"sync"
)

type Router struct {
	mu       sync.RWMutex
	commands map[string]command
	args     []string
}

func NewRouter(args []string) *Router {
	return &Router{
		mu:       sync.RWMutex{},
		commands: make(map[string]command),
		args:     args,
	}
}

func (r *Router) Register(pattern string, handlerFunc HandlerFunc, usage string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.commands[pattern]; ok {
		log.Fatal("Такая команда уже зарегистрирована!")
	}

	r.commands[pattern] = command{
		name:    pattern,
		handler: handlerFunc,
		usage:   usage,
	}
}

type flags map[string]string

func (r *Router) Serve() {

	var f = make(flags, len(r.commands))

	for c, h := range r.commands {
		f[c] = *flag.String(c, "", h.usage)
	}

	flag.Parse()

	for c, arg := range f {

		if _, ok := r.commands[c]; ok {
			r.commands[c].handler(arg)
			break
		}
	}
}

type HandlerFunc func(arg string)

type command struct {
	name    string
	handler HandlerFunc
	usage   string
}

package CLIRouter

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type HandlerFunc func(args []string)

type command struct {
	name    string
	handler HandlerFunc
	usage   string
}

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

func (r *Router) Register(pattern string, handlerFunc func(args []string), usage string) {
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

func (r *Router) Run() {

	cmd := r.args[1]
	args := r.args[2:]

	if h, ok := r.commands[cmd]; ok {
		h.handler(args)
	} else {
		fmt.Print(r.available())
	}

}

func (r *Router) available() string {
	firstMSG := fmt.Sprintf("Правильное использование данного приложения:\n")
	MSGs := make([]string, len(r.commands)+1)
	MSGs = append(MSGs, firstMSG)
	var msg string

	for c, h := range r.commands {
		msg = fmt.Sprintf("\t%s: %s\n", c, h.usage)
		MSGs = append(MSGs, msg)
	}

	helpMSG := strings.Join(MSGs, "")
	return helpMSG
}

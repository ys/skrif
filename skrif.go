package skrif

import (
	"log"
	"os"
)

type Skrif struct {
	l       *log.Logger
	context []interface{}
}

func New() *Skrif {
	s := &Skrif{l: log.New(os.Stdout, "", 0)}
	return s
}

func (s *Skrif) SetGlobalContext(context string) {
	s.l.SetPrefix(context + " ")
}

func (s *Skrif) Clone() *Skrif {
	newS := New()
	newS.l = s.l
	newS.context = s.context
	return newS
}

func (s *Skrif) Context(context string, fn func(*Skrif)) {
	newS := s.Clone()
	newS.context = append(s.context, context)
	fn(newS)
}

func (s *Skrif) Print(v ...interface{}) {
	s.l.Print(append(s.context, v...))
}

func (s *Skrif) Printf(format string, v ...interface{}) {
	s.l.Printf(format, v...)
}

func (s *Skrif) Println(v ...interface{}) {
	s.l.Println(append(s.context, v...)...)
}

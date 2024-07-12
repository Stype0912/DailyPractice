package main

import (
	"fmt"
)

func main() {
	m := &Manager{
		done: make(chan struct{}),
	}
	m.Add()
	m.Stop()
	<-m.done
}

type Manager struct {
	done chan struct{}
}

func (m *Manager) Stop() {
	close(m.done)
}

func (m *Manager) Wait() {
	<-m.done
	fmt.Println("wait done")
}

func (m *Manager) Add() {
	m.done <- struct{}{}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State struct {
	word   string
	listen chan string
}

func NewState(init string) State {
	return State{
		word:   init,
		listen: make(chan string),
	}
}

func (s State) GetState() string {
	return s.word
}

func (s State) StateString() {
	fmt.Println(s.word)
}

func (s *State) SendRandomWold(list []string) {
	s.listen <- list[rand.Intn(len(list))]
}

func main() {
	st := NewState("world")
	go st.Listener()
	go st.Render()

	myList := []string{"cat", "car", "man", "house", "go", "sport", "world"}
	for {
		time.Sleep(time.Duration(rand.Intn(10_000)) * time.Millisecond)
		st.SendRandomWold(myList)
	}
}

func (s *State) Render() {
	for {
		time.Sleep(500 * time.Millisecond)
		s.StateString()
	}
}

func (s *State) Listener() {
	for {
		s.word = <-s.listen
		fmt.Println("recive")
	}
}

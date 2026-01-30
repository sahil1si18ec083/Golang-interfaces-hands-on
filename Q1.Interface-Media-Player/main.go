package main

import "fmt"

type Player interface {
	Play()
}
type Music struct{}
type Video struct{}

func (m Music) Play() {
	fmt.Println("Playing music 🎵")

}
func (v Video) Play() {
	fmt.Println("Playing video 🎬")

}
func Start(p Player) {
	p.Play()
}

func main() {
	m := Music{}
	v := Video{}
	Start(m)
	Start(v)

}

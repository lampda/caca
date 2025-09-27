package main

import (
	"fmt"
	"reflect"
)

type Song struct {
	Lyrics string `json:"lyrics"`
}

type Silk struct {
	Name  string `json:"name"`
	Songs []Song `json:"songs"`
}

func uwu(s interface{}) {
	song, ok := s.(Song)
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(song.Lyrics)
}

func playground() {
	tmpSong := Song{Lyrics: "hot to go"}
	use(tmpSong)
	songs := []Song{Song{Lyrics: "uwu"}}
	silk := Silk{Name: "hornet", Songs: songs}
	t := reflect.TypeOf(silk.Songs).Elem()
	t2 := reflect.TypeOf(tmpSong)
	reflectedSong := reflect.New(t)
	reflectedSong2 := reflect.New(t2)
	fmt.Println(reflectedSong)
	fmt.Println(reflectedSong2)
}

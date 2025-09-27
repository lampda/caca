package unused

import (
	"fmt"
	"strconv"
)

type Fuzz interface {
	fizz() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type Silk struct {
	Hornet bool `json:"hornet"`
}

type Song struct {
	Bababui string `json:"bababui"`
}

func (s *Song) fizz() string {
	return "fizz"
}

func (s *Song) MarshalJSON() ([]byte, error) {
	fmt.Println("marshal")
	return []byte(s.Bababui), nil
}

func (s *Song) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshal")
	s.Bababui = string(data)
	return nil
}

func (s *Silk) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(s.Hornet)), nil
}

func (s *Silk) UnmarshalJSON(data []byte) error {
	hornet, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	s.Hornet = hornet
	return nil
}

func (s *Silk) fizz() string {
	return "fizz"
}

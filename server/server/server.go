package server

import (
	"encoding/json"
)

type MathService struct {
}
type Args struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (m *MathService) Add(args string, reply *int) error {
	var arg Args
	json.Unmarshal([]byte(args), &arg)
	*reply = arg.A + arg.B
	return nil
}
func (m *MathService) Sub(args string, reply *int) error {
	var arg Args
	json.Unmarshal([]byte(args), &arg)
	*reply = arg.A - arg.B
	return nil
}

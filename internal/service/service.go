package service

import (
	"L0/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	//"github.com/nats-io/stan.go"
)

var Count_id int

type repository interface {
	Create(id int, msg []byte) error
}
type Service struct {
	repository
}

func New(repository repository) *Service {
	return &Service{repository}
}

func (s *Service) ConsumeMessage() error {
	new_id := func() int {
		Count_id++
		return Count_id
	}
	nc, err := nats.Connect(nats.DefaultURL)
	//defer nc.Close()
	if err != nil {
		fmt.Println("Cann't connect to nats: ", err)
	}
	// Simple Publisher
	//nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		//fmt.Println(123)
		err = Validate(m.Data)
		if err == nil {
			id := new_id()
			fmt.Println(id)
			e := s.repository.Create(id, m.Data)
			if e != nil {
				fmt.Println("Cann't insert to table: ", err)
			} else {
				fmt.Printf("%d Received and insert a message: %s\n", id, string(m.Data))
			}
		} else {
			fmt.Printf("Received but not validate a message: %s\n", string(m.Data))
		}
	})
	return err
}
func Validate(data []byte) error {
	fmt.Println("Val")
	m := &model.Model{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	err := dec.Decode(m)
	if err != nil {
		return err
	}
	return nil
}

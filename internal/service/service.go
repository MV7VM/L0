package service

import (
	"L0/internal/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
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
	sc, _ := stan.Connect("test-cluster", "client-123", stan.NatsURL(stan.DefaultNatsURL))
	defer sc.Close()
	// Subscribe with durable name
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		err := Validate(m.Data)
		if err == nil {
			s.repository.Create(new_id(), m.Data)
		}
	}, stan.DurableName("my-durable"))
	defer sub.Unsubscribe()
	if err != nil {
		return err
	}
	return err
}
func Validate(data []byte) error {
	m := &model.Model{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	err := dec.Decode(m)
	if err != nil {
		return err
	}
	return nil
}

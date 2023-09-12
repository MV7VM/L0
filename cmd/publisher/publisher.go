package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"os"
	"os/exec"
	"time"
)

// this script allows published models on stan
func main() {
	//nc, err := nats.Connect(stan.DefaultNatsURL)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("!!!", err)
	}
	// Simple Publisher
	buff, err_read := os.ReadFile("/home/max/GolandProjects/L0/cmd/publisher/hole" + "/" + "model.json")
	if err_read != nil {
		fmt.Println("Fail to read file")
	}
	err = nc.Publish("foo", buff)
	if err != nil {
		fmt.Println("Fail to pub message")
	} else {
		fmt.Printf("Published [%s] : '%s'\n\n", "foo", string(buff))
	}

	defer nc.Close()

	//sc, err := stan.Connect("test-nats", "nats", stan.NatsURL(stan.DefaultNatsURL))
	//if err != nil {
	//	fmt.Printf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, nats.DefaultURL)
	//}
	//defer sc.Close()
	//
	//fmt.Println("Hole id active!")
	////hole("./hole", sc)
	//hole("cmd/publisher/hole", sc)
}

func hole(path string, sc stan.Conn) {
	for {
		dir, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)

		for _, fileInfo := range dir {
			if fileInfo.IsDir() {
				err = exec.Command("rm", "-rf", fileInfo.Name()).Run()
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			buff, err := os.ReadFile(path + "/" + fileInfo.Name())
			if err != nil {
				fmt.Println(err)
			}
			err = sc.Publish("updates", buff)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Published [%s] : '%s'\n\n", "updates", string(buff))

			err = exec.Command("rm", path+"/"+fileInfo.Name()).Run()
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}

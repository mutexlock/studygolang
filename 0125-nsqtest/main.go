package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type LogReportHandler struct {
}

var ll LogReportHandler

func (self *LogReportHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("%+v\n", message)
	return nil
}

func main() {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "channel1", cfg)
	if err != nil {
		fmt.Println(err)
	}
	consumer.AddConcurrentHandlers(&ll, 9)

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		fmt.Println(err)
	}

	select {}
}

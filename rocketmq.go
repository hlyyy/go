package main

import (
	"github.com/apache/rocketmq-client-go/core"
	"fmt"
)

func SendMessagge(){
    producer := rocketmq.NewProducer(config)
    producer.Start()
    defer producer.Shutdown()
    fmt.Printf("Producer: %s started... \n", producer)
    for i := 0; i < 100; i++ {
	    msg := fmt.Sprintf("%s-*d", *body, i)
        result, err := producer.SendMessageSync(&rocketmq.Message{Topic: "test", Body: msg})
        if err != nil {
            fmt.Println("Error:", err)
        }
	    fmt.Printf("send message: %s result: %s\n", msg, result)
    }
}

func main() {
	pConfig := &rocketmq.ProducerConfig{
		ClientConfig: rocketmq.ClientConfig{
			GroupID:    "GID_01",
			NameServer: "http://127.0.0.1:8080",
			Credentials: &rocketmq.SessionCredentials{
				AccessKey: "Your Access Key",
				SecretKey: "Your Secret Key",
				Channel:   "ALIYUN/OtherChannel",
			},
		},
		ProducerModel: rocketmq.CommonProducer,
	}
	sendMessage(pConfig)
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 1)
	conn.SetReadDeadline(time.Now().Add(time.Second * 0))
	// Used to read only one line of message
	// message,_ := conn.ReadMessage(1e6)
	// fmt.Println(string(message.Value))
	batch := conn.ReadBatch(1e3, 1e3) // 1e3 means 1000
	bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}

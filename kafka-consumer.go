package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"10.124.64.66:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to connect to Kafka server:", err)
		return
	}
	fmt.Println("Connect success...")
	defer consumer.Close()
	partitions, err := consumer.Partitions("springcloud-app")
	if err != nil {
		fmt.Println("Get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("springcloud-app", p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("PartitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

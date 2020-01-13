/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/swift9/rocketmq-client-go"
	"github.com/swift9/rocketmq-client-go/consumer"
	"github.com/swift9/rocketmq-client-go/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "TagA || TagC",
	}
	err := c.Subscribe("TopicTest", selector, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

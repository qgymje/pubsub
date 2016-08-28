package main

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

// 通过RabbitMQ的fanout类型的路由器, 实现房间弹幕收发机制
// 步骤
// 1. 主播创建一个房间, 连接MQ, 生成一个名为房间号, 类型为fanout的exchange, 作为生产者
// 2. 同时创建一个消费者, 用于接收消息
// 3. 玩家进入房间后, 生成同一个类型的exchange作为生产者, 并且同时创建一个消费者
// 4. 生产者通过websocket接收消息写入rabbitmq, 消费都通过rabbitmq接收消息写入websocket

// Client 表示一个keep alive的websocket连接以及一个连接RabbitMQ作为生产者的Channel和作为消费者的Channel
// 当客户端向websocket写消息时, websocket接收消息后, 通过producer向消息队列发送消息, 而每一个Client的consumer用于接收消息, 并且写到websocket里
type Client struct {
	conn     *websocket.Conn // 用于读取数据
	producer *amqp.Channel
	consumer *amqp.Channel
	name     string
	content  string
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

}

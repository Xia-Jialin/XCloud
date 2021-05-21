package locate

import (
	"XCloud/lib/rabbitmq"
	"os"
	"strconv"
	"time"
)

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}

package main

func main() {
	var rc RabbitClient
	rc.Consume("test-queue", funcName)
	rc.Publish("test-queue", mBody)

}

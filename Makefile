all:
	compile, build

compile:
	protoc -I order/ order/order.proto --go_out=plugins=grpc:order

build:
	go build -o nats_event/publisher nats_event/publisher
	go build -o nats_event/subscriber nats_event/subscriber
	cp -r publisher/config nats_event/


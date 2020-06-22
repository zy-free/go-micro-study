
proto:
	cd proto && protoc --proto_path=./ --micro_out=. --go_out=. greeter.proto

goose-create:
	cd migrations/ && goose mysql "root:123456@(localhost:3306)/bugu-activity?parseTime=true" create $(name) sql

goose-up:
	cd migrations/ && goose mysql "root:123456@(localhost:3306)/bugu-activity?parseTime=true" up

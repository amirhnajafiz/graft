proto:
	protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer

serve:
	go run server/main.go

use:
	go run client/main.go
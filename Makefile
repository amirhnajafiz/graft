proto:
	protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer

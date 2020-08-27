all:
	protoc -I/usr/local/include -I. --go_out=plugins=grpc,paths=source_relative:./ example/example.proto        
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:./ example/example.proto       
	protoc -I. --swagger_out=logtostderr=true:./ example/example.proto   
all:
	protoc --go_out=. --go_opt=paths=source_relative public.proto
	protoc --go-grpc_out=paths=source_relative:. public.proto
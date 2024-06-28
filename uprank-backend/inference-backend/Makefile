proto:
	python -m grpc_tools.protoc -Iproto --python_out=src/generated --pyi_out=src/generated --grpc_python_out=src/generated proto/inference.proto

proto_client:
	protoc --go_out=../queue-handler --go_opt=paths=source_relative \
    --go-grpc_out=../queue-handler --go-grpc_opt=paths=source_relative \
    proto/inference.proto

.PHONY: proto
.PHONY: proto_client
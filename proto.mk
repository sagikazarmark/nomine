PROTO_PATH = vendor/github.com/deshboard/boilerplate-proto/proto
GAPIS_PROTO_PATH = vendor/github.com/googleapis/googleapis
PROTOBUF_PROTO_PATH = vendor/github.com/google/protobuf/src

.PHONY: proto

proto: ## Generate code from protocol buffer
	@mkdir -p api
	protoc -I. -I ${GAPIS_PROTO_PATH} -I ${PROTOBUF_PROTO_PATH} nomine.proto --go_out=plugins=grpc:api --grpc-gateway_out=logtostderr=true:api

envcheck::
	$(call executable_check,protoc,protoc)
	$(call executable_check,protoc-gen-go,protoc-gen-go)
	$(call executable_check,protoc-gen-grpc-gateway,protoc-gen-grpc-gateway)

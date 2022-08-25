# SERVICE_RELATIVE_PATH=$(shell a=`basename $$PWD` && echo $$a)
PROTO_FILES=$(shell find . -name *.proto)

.PHONY: grpc
# generate grpc code
grpc:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
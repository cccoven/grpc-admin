# SERVICE_RELATIVE_PATH=$(shell a=`basename $$PWD` && echo $$a)
PROTO_FILES=$(shell find . -name *.proto)
# GOOGLE_PROTO_FILES=$(shell cd .. && cd .. && cd common && cd google && echo $$PWD)

.PHONY: grpc
# generate grpc code
grpc:
	protoc --go_out=. \
		--proto_path=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
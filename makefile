.PHONY: proto

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	protoc --proto_path=proto proto/src/*.proto --go_out=. --go-grpc_out=.

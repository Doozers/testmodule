test: generate
	go mod tidy
	go test -v .
	@echo "Done."

generate: testmodule.pb.go
.PHONY: generate

testmodule.pb.go: testmodule.proto
	protoc --go_out=./ --go-grpc_out=./ $<

clean:
	rm -f testmodule.pb.goa

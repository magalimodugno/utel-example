all: cover

cover:
	@go test -v -race -coverprofile --reportFile=./junit.xml -coverpkg=./... -coverprofile=coverage.out -outputdir=./test ./...

report:cover
	@go tool cover -html=./test/coverage.out -o ./test/coverage.html

lint:
	@go vet ./...

clean:
	@rm -fr **/*.{out,xml,html}

clean-cache:
	@go clean -cache
	@go clean -testcache
	@go clean -modcache


.PHONY: cover

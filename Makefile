
.PHONY: test
test:
	go test ./... --cover $(TAGS)

.PHONY: check-style
check-style:
	golangci-lint run -E errcheck -E gofmt -E revive

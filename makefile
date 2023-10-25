update-deps:
	go get -u ./...
	go mod tidy

test-ci:
	@echo "testing..."
	@go test -v -coverpkg=./... -coverprofile=tmp.cov ./...
	@EXIT_CODE=$$?
ifneq ($(COVERAGE), codecov)
	-@rm *.cov
endif
	@exit $$EXIT_CODE

generate-coverage-report:
	@cat tmp.cov | grep -v '.pb.go' > nongenerated_code.cov
	@go tool cover -func nongenerated_code.cov
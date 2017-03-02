FN_PKGS = $$(go list ./... | grep -v /vendor/)

LINTER_ARGS = \
	--vendor --disable=gas\
	--deadline=15m --tests

TEST_FLAGS = -v -race $(TEST_EXTRAFLAGS)

metalint:
	go get -u github.com/alecthomas/gometalinter; \
	gometalinter --install; \
	for pkg in $$(go list ./...); do go install $$pkg; done; \
	gometalinter $(LINTER_ARGS) ./...

test:
	go clean $(GO_EXTRAFLAGS) $(FN_PKGS)
	go test $(TEST_FLAGS) $(GO_EXTRAFLAGS) $(FN_PKGS)

coverage:
	sh test.sh

race:
	go test $(GO_EXTRAFLAGS) -race -i $(FN_PKGS)
	go test $(GO_EXTRAFLAGS) -race $(FN_PKGS)

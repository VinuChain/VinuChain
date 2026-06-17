.PHONY: all
all: opera

GOPROXY ?= "https://proxy.golang.org,direct"
.PHONY: opera
opera:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	GOPROXY=$(GOPROXY) \
	go build \
	    -ldflags "-s -w -X github.com/Fantom-foundation/go-opera/cmd/opera/launcher.gitCommit=$${GIT_COMMIT} -X github.com/Fantom-foundation/go-opera/cmd/opera/launcher.gitDate=$${GIT_DATE}" \
	    -o build/opera \
	    ./cmd/opera


TAG ?= "latest"
.PHONY: opera-image
opera-image:
	docker build \
    	    --network=host \
    	    -f ./docker/Dockerfile.opera -t "opera:$(TAG)" .

.PHONY: test
test:
	go test ./...

.PHONY: vulncheck
vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@v1.3.0 ./...

.PHONY: coverage
coverage:
	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock')
	go tool cover -func cover.prof | tail -n 1

# Native Go fuzzing (go test -fuzz) of the untrusted P2P message-decode path.
.PHONY: fuzz-native
fuzz-native:
	go test -run=^$$ -fuzz=FuzzHandleMsg -fuzztime=60s ./gossip/


.PHONY: clean
clean:
	rm -fr ./build/*

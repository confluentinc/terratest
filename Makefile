deps:
	@which dep 2>/dev/null || go get -u github.com/golang/dep/cmd/dep
	@which golint 2>/dev/null || go get -u golang.org/x/lint/golint
	@dep ensure

golint:
	@golint ./modules/...

installtf:
	wget https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip
	@which unzip 2>/dev/null || apt-get install -y unzip && rm -rf /var/lib/apt/lists/*
	unzip terraform_0.11.7_linux_amd64.zip
	mv ./terraform /usr/local/bin/

test: installtf golint
	@go test -v ./modules/...

.PHONY: deps golint test

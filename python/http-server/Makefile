SHELL	:= bash
.ONESHELL:
.SHELLFLAGS	:= -eu -o pipefail -c

REGISTRY	?= docker.io/pras/
IMAGE_NAME	?= simple-python-server
VERSION	?= 0.1.0

build-image:
	docker image build -t $(REGISTRY)$(IMAGE_NAME):$(VERSION) .

.PHONY: build-image

SHELL	:= bash
.ONESHELL:
.SHELLFLAGS	:= -eu -o pipefail -c

REGISTRY	?= docker.io/pras/
IMAGE_NAME	?= simple-go-server
TAG_NAME	?= $(shell git --no-pager tag -l --contains HEAD)

image:
	docker image build -t $(REGISTRY)$(IMAGE_NAME):$(TAG_NAME) .

.PHONY: image

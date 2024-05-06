TAG = tzuhsuanhuang/y-ws:latest
EXTERNAL_PORT = 13701
DOCKER_BUILD = docker build -t $(TAG)
DOCKER_PUSH = docker push $(TAG)
DOCKER_RUN = docker run -p 127.0.0.1:$(EXTERNAL_PORT):80 -it $(TAG)

build: Dockerfile package.json
	@echo "Building y-ws docker image..."
	$(DOCKER_BUILD) .
	@echo "Build complete."

push: build
	$(DOCKER_PUSH)

run: build
	$(DOCKER_RUN)

include .bingo/Variables.mk

metadata: $(MDATAGEN)
	$(MDATAGEN) ./pkg/receiver/nanoleafreceiver/metadata.yaml

build: $(BINGO) $(BUILDER)
	$(BUILDER) --config config/builder-config.yml

run: 
	./collector/home-otel-collector --config config.yaml

dev: metadata build run

docker-build:
	@echo "building docker container home-otel-collector"
	docker build -t home-otel-collector .

docker-run:
	@echo "running docker container"
	docker run -it -v $$PWD:/tmp -p 3333:3333  \
 		home-otel-collector:latest --config /tmp/config.yaml

docker: docker-build docker-run

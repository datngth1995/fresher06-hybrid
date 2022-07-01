all:
	docker-compose up -d
	
build:
	docker build -t my-golang -f Dockerfile .

.PHONY: all build
IMAGE_NAME=my-golang

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p 8080:80 -h ${IMAGE_NAME} --network my-net $(IMAGE_NAME)

IMAGE_NAME=my-nginx

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p 8000:80 -p 8001:81 -h ${IMAGE_NAME} --network my-net $(IMAGE_NAME)

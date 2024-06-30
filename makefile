build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p 80:80 $(IMAGE_NAME)

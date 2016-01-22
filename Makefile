compile:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o dummy-web-service

build:
	docker build -t quay.io/vietpham/dummy-web-service .

publish:
	docker push quay.io/vietpham/dummy-web-service

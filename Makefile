build:
	export GOPROXY=https://goproxy.cn && go build -o bin/qrcode
start:
	./bin/qrcode -d
stop:
	killall qrcode
docker-image:
	DOCKER_BUILDKIT=1 docker build -t hetao29/qrcode .
docker-image-nocache:
	DOCKER_BUILDKIT=1 docker build --no-cache -t hetao29/qrcode .
docker-push:
	docker push hetao29/qrcode:latest

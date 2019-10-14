build:
	cd src/qrcode / && go build -o ../../bin/qrcode .
start:
	./bin/qrcode -d
stop:
	killall qrcode

build:
	dep ensure

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/sensors/temperature/show/main.go
	mkdir -p bin/sensors/temperature
	zip bin/sensors/temperature/show.zip main
	mv main bin/sensors/temperature/show

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/sensors/humidity/show/main.go
	mkdir -p bin/sensors/humidity
	zip bin/sensors/humidity/show.zip main
	mv main bin/sensors/humidity/show

	env GOOS=linux go build -ldflags="-s -w" -o main endpoints/panel/index/main.go
	mkdir -p bin/panel
	zip bin/panel/index.zip main
	mv main bin/panel/index

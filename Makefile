src_files = cmd/adventcode22/day_one.go cmd/adventcode22/main.go

build:
	go build -o bin/adventcode22 ${src_files}

run:
	go run ${src_files}
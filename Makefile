src_files = cmd/adventcode22/day_three.go cmd/adventcode22/day_one.go cmd/adventcode22/main.go
test_files = cmd/adventcode22/day_three_test.go

build:
	go build -o bin/adventcode22 ${src_files}

run:
	go run ${src_files}

test:
	go test ${src_files} ${test_files}
src_files = cmd/adventcode22/main.go\
	cmd/adventcode22/day_one.go\
	cmd/adventcode22/day_two.go\
	cmd/adventcode22/day_three.go\
	cmd/adventcode22/day_four.go\
	cmd/adventcode22/day_five.go\
	cmd/adventcode22/day_six.go\
	cmd/adventcode22/day_seven.go\
	cmd/adventcode22/day_eight.go\
	cmd/adventcode22/day_nine.go\
	cmd/adventcode22/day_ten.go\
	cmd/adventcode22/day_eleven.go\
	cmd/adventcode22/day_fifteen.go\
	cmd/adventcode22/day_sixteen.go\
	cmd/adventcode22/day_twenty.go\
	cmd/adventcode22/day_twentyone.go

test_files = cmd/adventcode22/day_three_test.go\
	cmd/adventcode22/day_four_test.go\
	cmd/adventcode22/day_five_test.go\
	cmd/adventcode22/day_seven_test.go\
	cmd/adventcode22/day_fifteen_test.go

build:
	go build -o bin/adventcode22 ${src_files}

run:
	go run ${src_files}

test:
	go test ${src_files} ${test_files}

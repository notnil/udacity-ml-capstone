csv:
	sqlite3 -header -csv labels.db "select * from labels where label >= 0;" > labels.csv
run:
	glide install
	go run main.go -port=9090
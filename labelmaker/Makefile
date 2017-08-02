csv:
	sqlite3 -header -csv labels.db "select * from labels where label >= 0;" > labels.csv
run:
	go get github.com/mattn/go-sqlite3
	go run main.go -port=9090
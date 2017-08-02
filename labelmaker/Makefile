csv:
	sqlite3 -header -csv labels.db "select * from labels where label >= 0;" > labels.csv
run:
	glide install
	./labelmaker -port=9090
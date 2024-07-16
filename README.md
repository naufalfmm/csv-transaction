How to run:
1. Copy .env.example to .env
2. Fill the configurations
```
CSV_FILENAME=CSV_FILENAME # The path of csv with the filename

DB_HOST=DB_HOST # The database host 
DB_PORT=DB_PORT # The database port
DB_USERNAME=DB_USERNAME # The database username
DB_PASSWORD=DB_PASSWORD # The database password
DB_NAME=DB_NAME # The database name

DB_LOG_MODE=DB_LOG_MODE # The db log mode. Set true, if the query log want to be printed
DB_LOG_SLOW_THRESHOLD=DB_LOG_SLOW_THRESHOLD # The log slow threshold. The query runs more than the thresold will be printed as slow query

DB_RETRY=DB_RETRY # The db retry connection
DB_WAIT_SLEEP=DB_WAIT_SLEEP

LOG_MODE=LOG_MODE # Log mode. Set true if you want to log
```
3. Run go run main.go

NB:
1. The data2.csv filled with the 100000 dummy data
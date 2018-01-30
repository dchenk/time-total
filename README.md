# time-total
Easily count up hours worked in a date range.

A timesheet lists one work day per line, each line of the form:
```
MONTHNUM/DAYNUM HH:MM-HH:MM[, HH:MM-HH:MM]...
```

For example, give the program this argument:
```
1/9 9:20-9:50 10:05-10:45 11:25-2:55 3:30-4:30 5:30-6:00 6:15-8:05
1/11 9:10-9:25 10:15-12:15 12:30-1:30 2:40-3:40 3:55-5:05 5:45-6:00 8:35-9:00
1/12 10:50-11:20 11:40-2:30 2:50-3:20
```
like so:
```
cd $GOPATH/src/github.com/dchenk/time-total
go run main.go "1/9 9:20-9:50 10:05-10:45 11:25-2:55 3:30-4:30 5:30-6:00 6:15-8:05
1/11 9:10-9:25 10:15-12:15 12:30-1:30 2:40-3:40 3:55-5:05 5:45-6:00 8:35-9:00
1/12 10:50-11:20 11:40-2:30 2:50-3:20"
```

And you'll get this result printed to the terminal:
```
(1/9)  8.00 hrs total [9:20-9:50 10:05-10:45 11:25-2:55 3:30-4:30 5:30-6:00 6:15-8:05]
(1/11) 6.08 hrs total [9:10-9:25 10:15-12:15 12:30-1:30 2:40-3:40 3:55-5:05 5:45-6:00 8:35-9:00]
(1/12) 3.83 hrs total [10:50-11:20 11:40-2:30 2:50-3:20]
```

And a CSV file in the same directory as main.go looking like this:

|month/day |total hrs |
| -------- |----------|
|1/9       |8         |
|1/11      |6.08      |
|1/12      |3.83      |


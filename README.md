# Taxi-app-api

Taxi API + Load testing

-------------

## Usage

### Clone
Clone the Project

```sh
$ git clone https://github.com/Saloroler/taxi-app-api.git
$ cd taxi-app-api
```

### Running Taxi API
```sh
$ cd cmd/orderapi
$ go run main.go
```

This will run application on server at http://localhost:8085 <br />
You can change configuration PORT,COUNT_OF_ORDERS in .env config file

Also this will trigger cron job, which updating random order list every 200 milliseconds <br />

Module already included - no need to install dependencies

### API testing

1) **GET** request `/order` <br />
*OutPut*: <br />
Header: StatusOK <br />
JSON: 
```json
{"order": "{value}"} // randomly generated order
```
2) **GET** request `/admin/orders` <br />
*OutPut*: <br />
Header: StatusOK or StatusNotFound (if no previous order requests)<br />
JSON: 
```json
{
"{order_name}": {count_of_requests},
...
} 
```

### Apache/AB LOAD testing 

Load testing was perfomed on *local machine - MacbookPro 2017*

With max limit of user processes 

```sh
$ ulimit -u
2048
```

With max limit of open files

```sh
$ ulimit -n
524288
```

Optimal testing was perfomed without keepalive feauture and without changing default parameters <br />
With concurency 100, number of requests - 15000(because mac os out of [ephemeral ports](https://www.ncftp.com/ncftpd/doc/misc/ephemeral_ports.html) limit)
```sh
ab -n 15000 -c 100 http://localhost:8085/order
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1500 requests
Completed 3000 requests
Completed 4500 requests
Completed 6000 requests
Completed 7500 requests
Completed 9000 requests
Completed 10500 requests
Completed 12000 requests
Completed 13500 requests
Completed 15000 requests
Finished 15000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8085

Document Path:          /order
Document Length:        14 bytes

Concurrency Level:      100
Time taken for tests:   1.801 seconds
Complete requests:      15000
Failed requests:        0
Total transferred:      1965000 bytes
HTML transferred:       210000 bytes
Requests per second:    8329.17 [#/sec] (mean)
Time per request:       12.006 [ms] (mean)
Time per request:       0.120 [ms] (mean, across all concurrent requests)
Transfer rate:          1065.55 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  32.2      2     680
Processing:     1    7  35.2      5     587
Waiting:        0    6  34.9      4     587
Total:          1   12  47.6      7     688

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      9
  75%     10
  80%     10
  90%     11
  95%     12
  98%     15
  99%     26
 100%    688 (longest request)
```

Also was perfomed second test witth keep alive feature <br />

```sh
 ab -n 20000 -c 1000 -k http://localhost:8085/order
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8085

Document Path:          /order
Document Length:        14 bytes

Concurrency Level:      1000
Time taken for tests:   0.802 seconds
Complete requests:      20000
Failed requests:        0
Keep-Alive requests:    20000
Total transferred:      3100000 bytes
HTML transferred:       280000 bytes
Requests per second:    24928.73 [#/sec] (mean)
Time per request:       40.114 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          3773.39 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   8.7      0      60
Processing:     0   36  13.7     35     117
Waiting:        0   36  13.7     35     117
Total:          0   38  15.6     35     136

Percentage of the requests served within a certain time (ms)
  50%     35
  66%     38
  75%     40
  80%     41
  90%     56
  95%     69
  98%     90
  99%     96
 100%    136 (longest request)
```

### Race support
To prevent race conditions and panics was implemented mutex.Sync on writing in main order ticket list and on request history list.

Race conditions was found on library call, which I used - github.com/rakyll/ticktock.
```sh
go run -race main.go
2020/04/07 12:14:36 Server is starting on port 8085
==================
WARNING: DATA RACE
Write at 0x00c0000a8170 by goroutine 8:
  github.com/rakyll/ticktock.(*jobC).schedule()
      /Users/olegmaster/go/pkg/mod/github.com/rakyll/ticktock@v0.0.0-20140205200441-dd30f2fe99ad/ticktock.go:151 +0x30e
  github.com/rakyll/ticktock.(*jobC).schedule.func1()
      /Users/olegmaster/go/pkg/mod/github.com/rakyll/ticktock@v0.0.0-20140205200441-dd30f2fe99ad/ticktock.go:155 +0xf0

Previous write at 0x00c0000a8170 by goroutine 6:

  github.com/rakyll/ticktock.(*jobC).schedule()
      /Users/olegmaster/go/pkg/mod/github.com/rakyll/ticktock@v0.0.0-20140205200441-dd30f2fe99ad/ticktock.go:151 +0x30e
  github.com/rakyll/ticktock.(*Scheduler).Start()
      /Users/olegmaster/go/pkg/mod/github.com/rakyll/ticktock@v0.0.0-20140205200441-dd30f2fe99ad/ticktock.go:124 +0x114
  github.com/rakyll/ticktock.Start()
      /Users/olegmaster/go/pkg/mod/github.com/rakyll/ticktock@v0.0.0-20140205200441-dd30f2fe99ad/ticktock.go:67 +0x4a

Goroutine 8 (running) created at:
  time.goFunc()
      /usr/local/Cellar/go@1.12/1.12.9/libexec/src/time/sleep.go:169 +0x51

Goroutine 6 (running) created at:
  main.main()
      /Users/olegmaster/OlegItems/CodeActivity/IdeaSoftTest/taxi-app-api/cmd/orderapi/main.go:54 +0x6bc
==================
```
Working on 200 millisecond routine.

Rest part without races.

### Example 
1) Run application;
2) Open terminal in new window and run ab -n 15000 -c 100 http://localhost:8085/order to make 15000 request(you can skip it and just test it manually GET http://localhost:8085/order)
3) Open HTTPClient to provide HTTP API testing;
4) Make GET request http://localhost:8085/admin/orders <br />
Will receive response:
```json
{"ac":280,"af":168,"bh":327,"cr":306,"cx":18,"di":98,"dr":299,"du":149,"el":286,"ey":326,"fa":308,"fs":311,"gd":307,"gh":298,"gk":300,"im":116,"in":300,"ki":342,"ko":313,"kw":288,"lf":287,"ll":310,"lr":321,"ml":316,"mn":326,"ne":310,"ni":315,"nn":313,"nu":69,"oe":303,"oh":294,"os":299,"pe":307,"ph":300,"pk":300,"ps":18,"qf":329,"qh":301,"ro":305,"sc":275,"se":290,"tj":303,"uc":139,"ur":258,"vc":286,"vm":303,"vr":330,"vv":295,"wr":144,"xf":315,"xl":238,"xn":283,"xt":307,"yb":176,"yi":178,"yo":296,"yt":121}
```
You will see how many times was requested specific order.
Orders list changes every 200 milliseconds.

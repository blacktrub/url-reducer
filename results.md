GOLANG 
ab -r -n 1000000 -c 1000 http://82.146.54.161:80/api/read\?hash\=GE

Concurrency Level:      1000
Time taken for tests:   39.588 seconds
Complete requests:      1000000
Failed requests:        1870902
   (Connect: 0, Receive: 870903, Length: 129096, Exceptions: 870903)
Total transferred:      29436624 bytes
HTML transferred:       13427232 bytes
Requests per second:    25260.10 [#/sec] (mean)
Time per request:       39.588 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          726.14 [Kbytes/sec] received

PYTHON WORKERS 4
ab -r -n 1000000 -c 1000 http://82.146.54.161:8000/FTbhk7ku

Concurrency Level:      1000
Time taken for tests:   83.057 seconds
Complete requests:      1000000
Failed requests:        1827992
   (Connect: 0, Receive: 913996, Length: 0, Exceptions: 913996)
Non-2xx responses:      86047
Total transferred:      13509379 bytes
HTML transferred:       0 bytes
Requests per second:    12039.96 [#/sec] (mean)
Time per request:       83.057 [ms] (mean)
Time per request:       0.083 [ms] (mean, across all concurrent requests)
Transfer rate:          158.84 [Kbytes/sec] received

PYTHON WITHOUT DB WORKERS 4
ab -r -n 1000000 -c 1000 http://82.146.54.161:8000/FTbhk7ku

Concurrency Level:      1000
Time taken for tests:   48.087 seconds
Complete requests:      1000000
Failed requests:        2720043
   (Connect: 0, Receive: 906681, Length: 906681, Exceptions: 906681)
Total transferred:      14656030 bytes
HTML transferred:       1213342 bytes
Requests per second:    20795.43 [#/sec] (mean)
Time per request:       48.087 [ms] (mean)
Time per request:       0.048 [ms] (mean, across all concurrent requests)
Transfer rate:          297.64 [Kbytes/sec] received


PYTHON WITHOUT PYDANTIC WORKERS 4
Concurrency Level:      1000
Time taken for tests:   40.245 seconds
Complete requests:      1000000
Failed requests:        2595093
   (Connect: 0, Receive: 865031, Length: 865031, Exceptions: 865031)
Total transferred:      21191717 bytes
HTML transferred:       1754597 bytes
Requests per second:    24847.81 [#/sec] (mean)
Time per request:       40.245 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          514.23 [Kbytes/sec] received

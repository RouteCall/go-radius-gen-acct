```sh
go-radius-gen-acct.go --help
NAME:
   go-radius-gen-acct - A Go (golang) RADIUS client accounting (RFC 2866) implementation for perfomance testing

USAGE:
   go-radius-gen-acct - A Go (golang) RADIUS client accounting (RFC 2866) implementation for perfomance testing with generated data according dictionary (./dictionary.routecall.opensips) and RFC2866 (./rfc2866).

VERSION:
   0.11.8

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --pps value, -p value        packets per second (default: 10)
   --server value, -s value     server to send accts
   --port value, -P value       port to send accts (default: "1813")
   --nas-ip value               NAS-IP-Address on radius packet (default: "127.0.0.1")
   --nas-port value             NAS-Port on radius packet (default: 5666)
   --key value, -k value        key for acct
   --max-req value, -m value    stop the test and exit when max-req are reached (default: 9223372036854775807)
   --retry-int value, -r value  interval in second, on which to resend packet (zero or negative value means no retry) (default: 3)
   --stats, -c                  show count of requests
   --daemon, -d                 daemon (background) proccess
   --log-file value             the destination file of the log (default: "./go-radius-gen-acct.log")
   --pid-file value             file to save the pid of daemon (default: "./go-radius-gen-acct.pid")
   --help, -h                   show help
   --version, -v                print the version
```
---
```sh
# EXAMPLES

# run daemon (background)
go-radius-gen-acct --pps 100 --max-req 10000 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -d -c --log-file /tmp/go-radius-gen-acct.log --pid-file /tmp/go-radius-gen-acct.pid

# run on the shell, get counts and the max number of requests is 100
go-radius-gen-acct --pps 100 --max-req 10000 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -c 
```
---
```sh
# stats per second
tail /tmp/go-radius-gen-acct.log
2018/12/08 03:19:53
2018/12/08 03:19:53 Stats [refresh 1s]:
2018/12/08 03:19:53 estimated accounting-request per second:  101
2018/12/08 03:19:53 total count accounting-request:           9686
2018/12/08 03:19:54
2018/12/08 03:19:54 Stats [refresh 1s]:
2018/12/08 03:19:54 estimated accounting-request per second:  100
2018/12/08 03:19:54 total count accounting-request:           9786
2018/12/08 03:19:55
2018/12/08 03:19:55 Stats [refresh 1s]:
2018/12/08 03:19:55 estimated accounting-request per second:  99
2018/12/08 03:19:55 total count accounting-request:           9885
2018/12/08 03:19:56
2018/12/08 03:19:56 Stats [refresh 1s]:
2018/12/08 03:19:56 estimated accounting-request per second:  101
2018/12/08 03:19:56 total count accounting-request:           9986
2018/12/08 03:19:57
2018/12/08 03:19:57 Stats [refresh 1s]:
2018/12/08 03:19:57 estimated accounting-request per second:  14
2018/12/08 03:19:57 total count accounting-request:           10000
```


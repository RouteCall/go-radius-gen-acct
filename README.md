```sh
~# go-radius-gen-acct -h
NAME:
   go-radius-gen-acct - A Go (golang) RADIUS client accounting (RFC 2866) implementation for perfomance testing

USAGE:
   go-radius-gen-acct - A Go (golang) RADIUS client accounting (RFC 2866) implementation for perfomance testing with generated data according dictionary (./dictionary.routecall.opensips) and RFC2866 (./rfc2866).

VERSION:
   0.10.3

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --pps value, -p value      packets per second (default: 10)
   --server value, -s value   server to send accts
   --port value, -P value     port to send accts (default: "1813")
   --nas-ip value             NAS-IP-Address on radius packet (default: "127.0.0.1")
   --nas-port value           NAS-Port on radius packet (default: 5666)
   --key value, -k value      key for acct
   --max-req value, -m value  stop the test and exit when max-req are reached (default: 9223372036854775807)
   -c                         show count of requests
   -d                         daemon (background) proccess
   --log-file value           the destination file of the log (default: "./go-radius-gen-acct.log")
   --pid-file value           file to save the pid of daemon (default: "./go-radius-gen-acct.pid")
   --help, -h                 show help
   --version, -v              print the version
```

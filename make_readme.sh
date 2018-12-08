#!/usr/bin/env bash

DST_README="./README.md"
NULL="/dev/null"

source ~/.bash{rc,_profile} 2> ${NULL}

echo -e "\`\`\`sh
go-radius-gen-acct.go --help
$(go run go-radius-gen-acct.go --help 2> ${NULL})
\`\`\`" > "${DST_README}"

echo -e "\
---
\`\`\`sh
# EXAMPLES

# run daemon (background)
go-radius-gen-acct --pps 100 --max-req 10000 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -d -c --log-file /tmp/go-radius-gen-acct.log --pid-file /tmp/go-radius-gen-acct.pid

# run on the shell, get counts and the max number of requests is 100
go-radius-gen-acct --pps 100 --max-req 10000 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -c 
\`\`\`
---
\`\`\`sh
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
\`\`\`
" >> "${DST_README}"

git add README.md && git commit -m '[make_readme.sh] auto generated' && git push origin master

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
go-radius-gen-acct --pps 10 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -d -c --log-file /tmp/go-radius-gen-acct.log --pid-file /tmp/go-radius-gen-acct.pid

# run on the shell, get counts and the max number of requests is 100
go-radius-gen-acct --pps 10 --server 242.95.79.224 --port 1813 --nas-ip 199.143.213.200 --nas-port 5666 --key 4fW5xa28ba38e4e60d0wk187sdfi17454ef524x0z -c --max-req 100
\`\`\`
---
\`\`\`sh
# stats per second
tail /tmp/go-radius-gen-acct.log
2018/12/05 22:44:10 Stats [refresh 1s]:
2018/12/05 22:44:10 estimated accounting-request per second:  195
2018/12/05 22:44:10 total count accounting-request:           67187
2018/12/05 22:44:11
2018/12/05 22:44:11 Stats [refresh 1s]:
2018/12/05 22:44:11 estimated accounting-request per second:  196
2018/12/05 22:44:11 total count accounting-request:           67383
\`\`\`
" >> "${DST_README}"

git add README.md && git commit -m '[make_readme.sh] auto generated' && git push origin master

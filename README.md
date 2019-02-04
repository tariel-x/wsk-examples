# README

At host: `wsk property get --auth | cut -c13-` returns key.

At local: `wsk -i property set --apihost 0.0.0.0 --auth KEY --namespace guest`, where `KEY` is the result from the host.

Test at local: `wsk -i namespace get`.

## Content

1. [Hello world](hello/README.md).
2. [Hello world standalone binary](hello-standalone/README.md).

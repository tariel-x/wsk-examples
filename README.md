# README

At host: `wsk property get --auth | cut -c13-` returns key.

At local: `wsk -i property set --apihost 0.0.0.0 --auth KEY --namespace guest`, where `KEY` is the result from the host.

Test at local: `wsk -i namespace get`.

## Content

1. [Hello world](hello).
2. [Hello world standalone binary](hello-standalone).
3. [Deploy sequences with wskdeploy](wskdeploy)
4. [Implement own action loop](actionloop)
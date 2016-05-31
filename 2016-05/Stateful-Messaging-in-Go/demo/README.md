This is example code to illustrate the use of a `REQ/REP` (request/reply) socket using [go-mangos](https://godoc.org/bitbucket.org/gdamore/mangos).

Compile it using `make`. 

The server/client each have _help_ options (`-h`) and allow the various timeouts to be set from the commandline.

For example, start the client before the server:
```bash
$ ./client -recv 60s -send 30s -retry 1s
```

And then start the server:
```bash
$ ./server
```

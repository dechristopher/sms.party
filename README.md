# sms.party

- A simplistic, free, and dev-centric SMS API made with love by cSMS.
- Containerized, scalable, and dead simple.
- The people's SMS API - **free forever**.

First check your system requirements with *tools/requires.sh*

**THEN**

Build with *tools/docker-build.sh* and run with:
```
docker run -d --name sms-p -h sms-p -e PORT=<port> -p <port>:<port> sms-p
```

**OR**

Build and run all at once with our lifecycle scripts *tools/sms-up.sh* and *tools/sms-down.sh*

***

To simply build binaries, run *tools/build-(darwin/linux).sh* and start like so:

```
./build/main-(darwin/linux) <port>

EXAMPLE: ./build/main-darwin 8080
```

http://sms.party

**TODO**

- [X] Flesh out API
- [X] Send handler
- [X] Auth & IP middlewares
- [X] Data Models
- [X] Helper functions
- [X] Error & Info strings
- [X] Global config
- [X] config.json -> global config
- [X] API key genertion endpoint
- [ ] API key verification in auth mw
- [ ] API key storage & expiry
- [ ] Redis credential acquisition
- [ ] Cast and Batch endpoint handlers
- [ ] Rate limiting per API key (2s)
- [ ] Rate limiting per number (10s)
- [ ] Do not message list (API & Site)
- [ ] API key statistics
- [ ] Number statistics
- [ ] Total app statistics
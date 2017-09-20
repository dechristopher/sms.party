# sms.party

- A simplistic, free, and dev-centric SMS API made with love by cSMS.
- Containerized, scalable, and dead simple.
- The people's SMS API - **free forever**.

First check your system requirements with *tools/requires.sh*

**THEN**

Build with *tools/docker-build.sh* and run with:
```
*docker run -d --name sms-p -h sms-p -e PORT=<port> -p <port>:<port> sms-p*
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
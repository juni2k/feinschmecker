# Feinschmecker: FreeBSD Installation Guide

## Locally, on a Go-enabled machine
```
# Build for FreeBSD
$ go get -v -u github.com/go-bindata/go-bindata/...
$ go get -v -u github.com/PuerkitoBio/goquery
$ go get -v -u github.com/go-telegram-bot-api/telegram-bot-api
$ GOOS=freebsd make

# Upload build
$ rsync -czvaP build 'me@server:/usr/home/me/'
```

## On the server
```
# Create a group
$ sudo pw group add feinschmecker

# Create a new system account
$ sudo pw user add feinschmecker -c 'Feinschmecker Telegram Bot' -s /usr/sbin/nologin -g feinschmecker -d /var/db/feinschmecker

# Create a home directory
$ sudo mkdir /var/db/feinschmecker
$ sudo chown -R feinschmecker:feinschmecker /var/db/feinschmecker

# Move build
$ sudo mv build/ /var/db/feinschmecker/
$ sudo chown -R feinschmecker:feinschmecker /var/db/feinschmecker

# Find out if it's working (after creating config.json)
$ sudo -u feinschmecker /var/db/feinschmecker/build/feinschmecker -c /var/db/feinschmecker/config.json

# Install an rc.d script
$ cat > /usr/local/etc/rc.d/feinschmecker
#!/bin/sh

# PROVIDE: feinschmecker
# REQUIRE: LOGIN
# KEYWORD: shutdown

# Add the following lines to /etc/rc.conf to enable feinschmecker:
# feinschmecker_enable="YES"
# feinschmecker_flags="<set as needed>"

. /etc/rc.subr

name="feinschmecker"
rcvar=feinschmecker_enable

load_rc_config $name

: ${feinschmecker_enable="NO"}
: ${feinschmecker_config="/var/db/feinschmecker/config.json"}

pidfile="/var/run/feinschmecker.pid"
procname="/var/db/feinschmecker/build/feinschmecker"
command="/usr/sbin/daemon"
command_args="-f -T feinschmecker -p ${pidfile} -u feinschmecker /var/db/feinschmecker/build/feinschmecker -c ${feinschmecker_config}"

run_rc_command "$1"
^D

# Enable auto-startup
$ sudo sysrc feinschmecker_enable=YES
```

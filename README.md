# Feinschmecker

Find something inexpensive and unhealthy to eat at TU Hamburg.

## Motivation

Feinschmecker is a semi-ambitious rewrite of [tuhh-mensa-bot][tmb].  It
is more robust and more performant than its predecessor while also being
easier to maintain and deploy.

Most of it is written in Go. Minor parts dealing with text munging and
annoying input are written in Perl.

[tmb]: https://github.com/nanont/tuhh-mensa-bot

## Building, Deployment, etc.

[I have written a guide to build and install things on FreeBSD, and I
think that's sufficient due to the fairly easy build process.][doc]

[doc]: https://github.com/nanont/feinschmecker/blob/master/contrib/INSTALL-FreeBSD.md

### Runtime requirements

- The actual binary
- Perl 5.18 or newer (already present on macOS and modern *nix systems)

## License

This software is Copyright (c) 2019 by Martin Borchert.

This is free software, licensed under:

  The GNU General Public License, Version 3, June 2007

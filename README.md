# Nomine

[![Build Status](https://img.shields.io/travis/sagikazarmark/nomine.svg?style=flat-square)](https://travis-ci.org/sagikazarmark/nomine)
[![Go Report Card](https://goreportcard.com/badge/github.com/sagikazarmark/nomine?style=flat-square)](https://goreportcard.com/report/github.com/sagikazarmark/nomine)

**Check your desired name in various services.**

When choosing a name for a new project or even just a nickname, it's often annoying
that the desired name is already taken in certain online services. It's even worse
when you find that out *after* you started registering with that name.

Nomine is here to help. It checks your desired name in different services,
so that you can find a globally unique name.

(Nomine is *name* in Latin)


## Supported services

- [ ] Github
- [ ] Twitter
- [ ] Docker
- [ ] Domain names (using Namecheap API)


## How does it work?

When you try to register to a service, it often provides a way to check if the chosen name
is available or not. It's often done asynchronously. When possible, Nomine uses these name
checkers, because sometimes registered names also cannot be certain words (like in case of Github).

If the above trick does not work, Nomine tries to check the name using some kind of API.
This is probably less accurate, but should be fine as well.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

# Gomematic: API server

[![Build Status](http://drone.gomematic.tech/api/badges/gomematic/gomematic-api/status.svg)](http://drone.gomematic.tech/gomematic/gomematic-api)
[![Stories in Ready](https://badge.waffle.io/gomematic/gomematic-api.svg?label=ready&title=Ready)](http://waffle.io/gomematic/gomematic-api)
[![Join the Matrix chat at https://matrix.to/#/#gomematic:matrix.org](https://img.shields.io/badge/matrix-%23gomematic%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gomematic:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ca2aacc664fb4d118b920fd7068edf37)](https://www.codacy.com/app/gomematic/gomematic-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomematic/gomematic-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gomematic/gomematic-api?status.svg)](http://godoc.org/github.com/gomematic/gomematic-api)
[![Go Report](https://goreportcard.com/badge/github.com/gomematic/gomematic-api)](https://goreportcard.com/report/github.com/gomematic/gomematic-api)
[![](https://images.microbadger.com/badges/image/gomematic/gomematic-api.svg)](http://microbadger.com/images/gomematic/gomematic-api "Get your own image badge on microbadger.com")


**This project is under heavy development, it's not in a working state yet!**

TBD


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.gomematic.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/gomematic/homebrew-gomematic).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/gomematic/gomematic-api/cmd/gomematic-api` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/gomematic/gomematic-api
cd $GOPATH/src/github.com/gomematic/gomematic-api
make clean build

./gomematic-api -h
```


## Security

If you find a security issue please contact gomematic@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

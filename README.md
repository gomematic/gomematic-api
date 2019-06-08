# Gomematic: API server

[![Build Status](https://cloud.drone.io/api/badges/gomematic/gomematic-api/status.svg)](https://cloud.drone.io/gomematic/gomematic-api)
[![Join the Matrix chat at https://matrix.to/#/#gomematic:matrix.org](https://img.shields.io/badge/matrix-%23gomematic-7bc9a4.svg)](https://matrix.to/#/#gomematic:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ca2aacc664fb4d118b920fd7068edf37)](https://www.codacy.com/app/gomematic/gomematic-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomematic/gomematic-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gomematic/gomematic-api?status.svg)](http://godoc.org/github.com/gomematic/gomematic-api)
[![Go Report](https://goreportcard.com/badge/github.com/gomematic/gomematic-api)](https://goreportcard.com/report/github.com/gomematic/gomematic-api)
[![](https://images.microbadger.com/badges/image/gomematic/gomematic-api.svg)](http://microbadger.com/images/gomematic/gomematic-api "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Gomematic will be a simple web interface to manage my own home automation based on Homematicc because I don't really like the currently available interfaces for that. I thought it's time to implement a shiny application with Go for the API and with VueJS for the UI.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.gomematic.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/gomematic/homebrew-gomematic).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/gomematic/gomematic-api.git
cd gomematic-api

make generate build

./bin/gomematic-api -h
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

# d2n
[![Go Report Card](https://goreportcard.com/badge/github.com/muehlburger/d2n)](https://goreportcard.com/report/github.com/muehlburger/d2n)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/muehlburger/d2n/blob/master/LICENSE)

d2n adds ISO time-stamps to files.

[Features](#features) | [Installation](#installation) | [Usage](#usage) | [Examples](#examples) | [Command-line options](#options) | [Configuration](#configuration)

## Features

## Installation

```bash
go get -u github.com/muehlburger/d2n
```

## Usage

### Examples

Rename files:
1. Add ISO Timestamps
2. Add Filenames

```bash
d2n rename -s ./folder

2019/10/13 21:09:49 rename /tmp/IMG_20191011_145813.jpg -> /tmp/2019-10-11T14.58.15 IMG_20191011_145813.jpg
```

## Authors

[Herbert Mühlburger](https://github.com/muehlburger) and [contributors](https://github.com/muehlburger/d2n/graphs/contributors).

## License

[Apache 2.0 License](LICENSE)

[report-card-image]: https://goreportcard.com/badge/github.com/muehlburger/d2n
[report-card-url]: https://goreportcard.com/report/github.com/muehlburger/d2n

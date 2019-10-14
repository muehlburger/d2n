# d2n
[![Go Report Card](https://goreportcard.com/badge/github.com/muehlburger/d2n)](https://goreportcard.com/report/github.com/muehlburger/d2n)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/muehlburger/d2n/blob/master/LICENSE)

d2n adds timestamps in ISO 8601+ format
YYYY-MM-DD (http://datestamps.org/index.shtml) at the beginning of
the filenames to files. d2n is inspired by Karl Voit's article on "[Managing Digital Files]" and his [date2name] tool but not 100% compatible.

If the timestamp exists at the beginning of the filename, it will be ignored. Executed with an examplefilename of "image.jpg", it results in "2006-01-02T15.04.05 image.jpg".

Note: Other that defined in ISO 8601+ the delimiter between hours,
minutes, and seconds is not a colon but a dot. Colons are causing
several problems on different file systems and are therefore replaced
with the (older) DIN 5008 version with dots. (see [date2name] for more information).

[Features](#features) | [Installation](#installation) | [Usage](#usage) | [Examples](#examples) | [Command-line options](#options) | [Configuration](#configuration)

## Features

- Adds timestamps in ISO 8601+ format to filenames
- Delimits hours, minutes, seconds using dots as in DIN 5008

## Installation

```bash
go get -u github.com/muehlburger/d2n
```

## Usage

```bash
d2n rename -s ./image.jpg

2019/10/14 10:43:09 rename /image.jpg -> /2019-10-14T10.43.05 test.jpg
```

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

[Managing Digital Files]: https://karl-voit.at/managing-digital-photographs/
[date2name]: https://github.com/novoid/date2name
[![License][bsd-batch]][bsd] [![Go Report Card][goreport-batch]][goreport] [![Build Status][travis-batch]][travis] [![GoDoc][cover-batch]][cover] [![Covarage][doc-batch]][doc]
# Answers Your dev Questions via The Command Line
`wie` means "how" in german and helps you to answer your programmer questions in the command line.

## Usage
```bash
wie golang set $GOPATH # export GOPATH=$HOME/go
wie docker link container # --link
```

## Build
```bash
go get -u github.com/sbani/wie
```

## Roadmap
- [x] Add display method `all` (short `a`): Show complete answer
- [ ] Add number parameter to get more than one answer
- [ ] Cache answers
- [x] Create something like an *attach/detach mode* with dynamic reaction
- [ ] Add more search engines
- [ ] Add user configuration files where one can change default behaviour
- [ ] Add unit tests
- [ ] Pretty print results

## Contributors / Special Thanks
This little helper is a port of the tool [howdoi][howdoi] written in Python. Special thanks to [Benjamin Gleitzman][gleitz]!

## Author
[Sufijen Bani][sbani] ([@sbani_ger][twit])

## License
The [BSD 3-Clause license][bsd], the same as the [Go language][golic].
goquery's license is [here][qrylic].

[howdoi]: https://github.com/gleitz/howdoi
[bsd]: http://opensource.org/licenses/BSD-3-Clause
[bsd-batch]: https://img.shields.io/badge/license-BSD--3--Clause-blue.svg
[golic]: https://golang.org/LICENSE
[qrylic]: https://github.com/PuerkitoBio/goquery/blob/master/LICENSE
[gleitz]: https://twitter.com/gleitz
[sbani]: http://sbani.net
[twit]: https://twitter.com/sbani_ger
[goreport-batch]: https://goreportcard.com/badge/github.com/sbani/wie
[goreport]: https://goreportcard.com/report/github.com/sbani/wie
[travis]: https://travis-ci.org/sbani/wie
[travis-batch]: https://travis-ci.org/sbani/wie.svg?branch=master
[doc]: https://godoc.org/github.com/sbani/wie
[doc-batch]: https://godoc.org/github.com/sbani/wie?status.svg
[cover]: https://gocover.io/github.com/sbani/wie
[cover-batch]: http://gocover.io/_badge/github.com/sbani/wie

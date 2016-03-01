[![Go Report Card](https://goreportcard.com/badge/github.com/sbani/wie)](https://goreportcard.com/report/github.com/sbani/wie) [![License](https://img.shields.io/badge/license-BSD--3--Clause-blue.svg)][bsd]

# Answers Your DEV Questions Via The Command Line
`wie` means "how" in german and helps you to answer your programmer questions in the command line.

## Usage
### Simple/Short
```bash
wie "golang set $GOPATH" # export GOPATH=$HOME/go
wie "docker link container" # --link
```
### Complete answer
```bash
wie -a "golang set $GOPATH" # You will see the hole answer text
```

## Build
```bash
go get github.com/sbani/wie
```

## Roadmap
- [x] Add display method `all` (short `a`): Show complete answer
- [ ] Add number parameter to get more than one answer
- [ ] Cache answers
- [ ] Create something like an *attach/detach mode* with dynamic reaction
- [ ] Add more search engines
- [ ] Add user configuration files where one can change default behaviour
- [ ] Add unit tests
- [ ] Pretty print results

## Contributors / Special Thanks
This little helper is a port of the tool [howdoi](https://github.com/gleitz/howdoi) written in Python. Special thanks to [Benjamin Gleitzman][gleitz]!

## Author
[Sufijen Bani][sbani] ([@sbani_ger][twit])

## License
The [BSD 3-Clause license][bsd], the same as the [Go language][golic].
Kingpin's license is [here][kinglic]. goquery's license is [here][qrylic]. color's license is [here][colorlic]

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[golic]: https://golang.org/LICENSE
[kinglic]: https://github.com/alecthomas/kingpin/blob/master/COPYING
[qrylic]: https://github.com/PuerkitoBio/goquery/blob/master/LICENSE
[colorlic]: https://github.com/fatih/color/blob/master/LICENSE.md
[gleitz]: https://twitter.com/gleitz
[sbani]: http://sbani.net
[twit]: https://twitter.com/sbani_ger
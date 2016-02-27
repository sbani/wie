# Answers Your dev Questions via The Command Line
`wie` means "how" in german and helps you to answer your programmer questions in the command line.

## Usage
### Simple
```bash
wie "golang set $GOPATH" # export GOPATH=$HOME/go
wie docker link container # --link
```

## Build
```bash
go get github.com/sbani/wie
```

## Roadmap
- [ ] Add display method `all` (short `a`): Show complete answer
- [ ] Add number parameter to get more than one answer
- [ ] Cache answers
- [ ] Create something like an *attach/detach mode* with dynamic reaction
- [ ] Add more search engines
- [ ] Add user configuration files where one can change default behaviour

## Contributors / Special Thanks
This little helper is a port of the tool [howdoi](https://github.com/gleitz/howdoi) written in Python. Special thanks to [Benjamin Gleitzman][gleitz]!

## Author
[Sufijen Bani][sbani] ([@sbani_ger][twit])

## License
The [BSD 3-Clause license][bsd], the same as the [Go language][golic]. Kingpin's license is [here][kinglic]. goquery's license is [here][qrylic]

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[golic]: https://golang.org/LICENSE
[kinglic]: https://github.com/alecthomas/kingpin/blob/master/COPYING
[qrylic]: https://github.com/PuerkitoBio/goquery/blob/master/LICENSE
[gleitz]: https://twitter.com/gleitz
[sbani]: http://sbani.net
[twit]: https://twitter.com/sbani_ger
docc [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/tcnksm/docc/blob/master/LICENCE)
==== 

`docc` open your project GitHub page or README.

## Usage

`docc` try to open your github page by Browser:

```bash
$ docc
$ docc /path/to/project
```

If you want to open project `README.md` by `$EDITOR`:

```
$ docc -e
$ docc -e /path/to/project
```

## Configuration

You can set command to open `README.md` in `~/.gitconfig`.

For example, if you want to use `more`:

```
[docc]
    cmd = more
```

Then execute with `-c` option:

```
$ docc -c
$ docc -c /path/to/project
```
## Install

```bash
$ brew tap tcnksm/docc
$ brew install docc
```

## VS.

- [motemen/git-browse-remote](https://github.com/motemen/git-browse-remote) - Extension of git command. It support opening branch/tag page ("ref" mode) and commit page ("rev" mode). docc is not git extension. docc aim to refer documentation(Github page or README). 

## Contribution

1. Fork ([https://github.com/tcnksm/docc/fork](https://github.com/tcnksm/docc/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

You can get source with `go get`:

```bash
$ go get -d github.com/dotcloud/docker
$ go get github.com/tcnksm/docc
```

## Author

[tcnksm](https://github.com/tcnksm)

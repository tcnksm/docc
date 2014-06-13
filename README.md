# docc

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

## Development

```bash
$ go get -d github.com/dotcloud/docker
$ go get github.com/tcnksm/docc
```


## Author

[tcnksm](https://github.com/tcnksm)

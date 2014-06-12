# docc

`docc` open your project GitHub page.

## Usage

`docc` try to open your github page by Browser.

```bash
$ docc
$ docc /path/to/project
```

You can open project `README.md` by `$EDITOR`.

```
$ docc -e
$ docc -e /path/to/project
```

## Configuration

You can set command to open `README.md` in `~/.gitconfig`.

For example, if you want to use `more`,

```
[docc]
    cmd = more
```

Then execute `docc` it with `-c` option.

```
$ docc -c
$ docc -c /path/to/project
```
## Install

```bash
$ brew tap tcnksm/docc
$ brew install docc
```

## Development

```bash
$ go get -d github.com/dotcloud/docker
$ go get github.com/tcnksm/docc
```

## Author

[tcnksm](https://github.com/tcnksm)

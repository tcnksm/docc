# docc

`docc` open your project document.

## Usage

By default, `docc` try to open your github page by Browser.

```bash
$ docc
$ docc /path/to/project
```

## Configuration

You can open project `README.md` by your favorite command. Set it in `~/.gitconfig`. 
```
[docc]
    cmd = more
```

## Install

```bash
$ go get -d github.com/dotcloud/docker
$ go get github.com/tcnksm/docc
```

## Author

[tcnksm](https://github.com/tcnksm)

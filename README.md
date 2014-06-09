# docc

`docc` open your project document.

## Usage

By default, `docc` try to open your github page by Browser

```bash
$ docc
$ docc /path/to/project
```

## Configuration

You can set default opening command in `~/.gitconfig` and open `README.md` by it

```
[docc]
    cmd = "more"
```

## Install

```bash
$ go get -d github.com/dotcloud/docker
$ go get github.com/tcnksm/docc
```

**ToDo**

```bash
$ brew tap tcnksm/docc
$ brew install docc
```

## Author

[tcnksm](https://github.com/tcnksm)

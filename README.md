# docc

`docc` open your project document.

## Usage

By default, `docc` try to open your github page by Browser

```bash
$ docc
$ docc /path/to/project
```

You can also open `README.md` by your `$EDITOR`.

```bash
$ docc -e 
$ docc -e /path/to/project
```
You can use your favorite command for opening `README.md`

```bash
$ docc -c 'more'
$ docc -c 'more' /path/to/project
```

## Configuration

You can set default opening command in `~/.gitconfig`.

```
[docc]
    cmd = "less"
```

This configutation is defaultly used, so just type

```bash
$ docc
$ docc /path/to/project
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

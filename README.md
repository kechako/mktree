# mktree

## Install

``` sh
$ go get github.com/kechako/mktree
```

## Usage

``` sh
$ cat input.txt
AAAA
  BBBB
    CCCC
      DDDD
      EEEE
    FFFF
  GGGG
    HHHH
$ mktree input.txt
# or
$ cat input.txt | mktree
AAAA
├── BBBB
│   ├── CCCC
│   │   ├── DDDD
│   │   └── EEEE
│   └── FFFF
└── GGGG
    └── HHHH
```

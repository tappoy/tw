# Package
`github.com/tappoy/tw`

# About
This golang package provides a simple way to write memo to a file, and output it to the stdout.

# Features
- Append memo to a file `~/.tw.txt` with a timestamp
  - ex) `2024-04-24 14:32:00 tappoy@thinkpad1: Hello, world!`
- Output memo to the stdout

# Installation
```bash
go install github.com/tappoy/tw@latest
```

# Usage
```
$ tw Hello, world! # append 'Hello, world!' to ~/.tw/tw.txt
$ tw -a              # print all memo
$ tw -t              # print today's memo
$ tw -g 'Hello'      # search 'Hello' using regex in all memo and print it
$ tw -v              # show version
```

# License
[GPL-3.0](LICENSE)

# Author
[tappoy](https://github.com/tappoy)

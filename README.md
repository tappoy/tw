# Package
`github.com/tappoy/tw`

# About
This golang package provides a simple way to write memo to a file, and output it to the stdout.

# Features
- Append memo to a file `~/.tw.txt` with a timestamp
  - ex) `2024-04-24 12:06 Hello, world!`
- Output memo to the stdout

# Usage
```
$ tw 'Hello, world!' # append 'Hello, world!' to ~/.tw.txt
$ tw -a              # print all memo
$ tw -t              # print today's memo
$ tw -g 'Hello'      # search 'Hello' using regex in all memo and print it
```

# License
[GPL-3.0](LICENSE)

# Author
[tappoy](https://github.com/tappoy)

# files-cli-tool-go
A Golang based CLI tool for reading files.

This is a cheap copy of the Unix Command line tool wc. 

## How to build
Execute `make build` to create the executable. It will be saved in the `/bin` directory.

## How to run
Execute the binary with a file as an argument to get the result.

The following are the possible arguments:
```
# Outputs the number of bytes

> ./bin/ccwc -c test.text
341836 test.txt

# Outputs the number of line breaks

> ./bin/ccwc -l test.text
7137 test.txt

# Outputs the number of characters

> ./bin/ccwc -m test.txt
339120 test.txt

# Outputs the number of words

> ./bin/ccwc -w test.text
58159 test.txt

# Outputs with -c, -l and -w flags
> ./bin/ccwc test.text
7137 58159 341836 test.txt
```
You can also read from standard input
```
> cat test.txt | ./bin/ccwc -l
7137
```

# gisty

Gisty is a command line tool to upload your files as gists onto your github account.

## Installation

This software requires `go` to be installed to your machine. You can get `go` [here](http://golang.org/dl/).

After installig `go`, run these commands in your shell.

```Shell
~ git clone git@github.com:akoskaaa/gisty.git
~ cd gisty
~ ./scripts/setup.sh
```

## Usage
Gisty gets your github information from your global git config. The experience is the best when you have your suername and password there as well. If any of those are missing, gisty will prompt you for the required information.

```Shell
~ gisty filename.txt
All done! Find your uploaded files @ https://gist.github.com/akoskaaa/
```

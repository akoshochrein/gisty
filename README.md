# gisty

Gisty is a command line tool to upload your files as gists onto your GitHub account.

## Installation

This software requires `Go` to be installed to your machine. Follow `Go`'s [installation instructions](https://golang.org/doc/install) if needed.

After installig `Go`, run these commands in your shell.

```Shell
git clone git@github.com:akoskaaa/gisty.git
cd gisty
. ./scripts/setup.sh
```

## Usage

Gisty gets your GitHub information from your global Git config. The experience is best when you also have your username and password configured. If any of those are missing, gisty will prompt you for the required information.

```Shell
gisty filename.txt
# All done! Find your uploaded files @ https://gist.github.com/akoskaaa/
```

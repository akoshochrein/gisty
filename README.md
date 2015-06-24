# gisty

Gisty is a command line tool to upload your files as gists onto your github account.

## Setup

This software requires `go` to be installed to your machine. You can get `go` [here](http://golang.org/dl/).

After installig `go`, run these commands in your shell.

```Shell
git clone git@github.com:akoskaaa/gisty.git
cd gisty
. ./scripts/setup.sh
```

Gisty gets your github information from your global git config. The experience is the best when you have your username and an access token there as well, like this:

```
[user]
        name = <your github username>
        token = <hexadecimal characters of an access token>
```

Access tokens with fine-grained permission settings can be created on [your github account's settings page] (https://github.com/settings/tokens).

## Usage
```Shell
gisty filename.txt
# All done! Find your uploaded files @ https://gist.github.com/akoskaaa/
```

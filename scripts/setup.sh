#!/bin/bash

# ensure go
command -v go >/dev/null 2>&1 || {
    echo >&2 "I cannot find go. Please install it and run the command again!";
    exit 1;
}

OLD_GOPATH=$GOPATH
export GOPATH=$GOPATH:`pwd`

# intall dependencies
echo "Installing dependencies"
go get github.com/howeyc/gopass
echo "Installed dependencies"

echo "Installing gisty"
go install gisty
echo "Installed gisty"

export PATH=$PATH:`pwd`/bin
if [ -f ~/.zshrc ]; then
    echo "export PATH=\$PATH:`pwd`/bin" >> ~/.zshrc     
fi
if [ -f ~/.bashrc ]; then
    echo "export PATH=\$PATH:`pwd`/bin" >> ~/.bashrc 
fi

echo "Added gisty to your path!"
echo "All done!"

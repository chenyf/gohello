#!/bin/bash
set -x

srcdir=$1
dstdir=$2
appname=$3
cluster=$4

mkdir -p /letv/gopath
export HOME=/root
export GOPATH=/letv/gopath
export GO15VENDOREXPERIMENT=1

eval "$(gimme 1.5.4)"
[ $? -ne 0 ] && { exit 1; }

case $appname in
    pushapi)
        cd $dstdir
        go build -o gohello github.com/chenyf/gohello
        [ $? -ne 0 ] && { exit 2; }
        cp -f $srcdir/run.sh $dstdir/
        ;;
    *)
        exit 1
        ;;
esac	

exit 0


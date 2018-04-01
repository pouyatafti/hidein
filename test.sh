#!/bin/sh

tmpdir=$(mktemp -d)
srcim=$tmpdir/src.png
destim=$tmpdir/dest.png
msg="hello, world!"

wget -O $srcim "http://via.placeholder.com/350x150" >/dev/null 2>&1

echo "$msg" | bin/encode $srcim - >$destim
[ "X$(bin/decode $destim ${#msg})" == "X$msg" ] && echo "ok" || echo "failed!"

#!/bin/sh
# Please install upx first, https://github.com/upx/upx/releases
find ./build -xdev -maxdepth 1 -type f -iname 'check-in*'  -exec upx --best --brute --ultra-brute {} \;
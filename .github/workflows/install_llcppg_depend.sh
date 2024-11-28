#!/bin/bash
set -e

llgo install ./_xtool/llcppsymg
git clone -b llcppsigfetch/include https://github.com/luoliwoshang/llcppg.git llcppg_temp
cd llcppg_temp
llgo install ./_xtool/llcppsigfetch
cd ..
rm -rf llcppg_temp
go install ./cmd/gogensig

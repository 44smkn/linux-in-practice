#!/usr/bin/env bash

# taskset -c <CPU番号> <プログラム> <同時に動かすプロセス数> <プログラムを動作させる時間> <統計情報の採取間隔>
# 時間はミリ秒単位
taskset -c 0 /go/bin/app 1 100 1
#taskset -c 0 /go/bin/app 2 100 1
#taskset -c 0 /go/bin/app 4 100 1
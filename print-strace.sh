#!/usr/bin/env bash

echo "-------------------------PROGRAM START-------------------------"
strace -o trace.log /go/bin/app
echo "--------------------------PROGRAM END--------------------------"
echo "-------------------------SYSCALL START-------------------------"
cat trace.log
echo "--------------------------SYSCALL END--------------------------"
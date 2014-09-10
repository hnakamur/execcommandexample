#!/bin/sh
echo out1
sleep 1
echo err1 1>&2
sleep 2
echo out2
echo out3
sleep 1
echo err2 1>&2
exit 1

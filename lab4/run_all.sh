#!/bin/bash

args=`find dataset -type f | xargs`

time bash go/serial/run.sh $args
time bash go/concurrent-0/run.sh $args
time bash go/serial-partial/run.sh $args
time bash go/concurrent-partial/run.sh $args

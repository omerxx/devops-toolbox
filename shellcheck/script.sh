#!/bin/bash


if [ -z "$X" ]; then
	echo "HI"
fi	

if [[ 1 -lt $X ]]; then
	echo "yes 1 is smaller than 2"
else
	echo "something is wrong"
fi

#!/bin/bash

if [[ -f /env ]]; then
  source /env
fi

/usr/sbin/kesshad

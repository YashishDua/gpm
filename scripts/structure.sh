#!/usr/bin/env bash

echo "Setting up project..."

add_touch_file()
{
  if [ `ls | wc -l` -gt 1 ]
  then
      : #directory already have content
  else
      touch .keep
  fi
}

# cmd Directory
mkdir -p cmd && cd cmd
add_touch_file
cd ..

# internal Directory
mkdir -p internal && cd internal
add_touch_file
cd ..

# pkg Directory
mkdir -p pkg && cd pkg
add_touch_file
cd ..

# scripts Directory
mkdir -p scripts && cd scripts
add_touch_file
cd ..

# api Directory
mkdir -p api && cd api
add_touch_file
cd ..

# test Directory
mkdir -p test && cd test
add_touch_file
cd ..

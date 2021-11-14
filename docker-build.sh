#!/bin/bash

function build_image (){
  if [ -z "${BUILD_ENV}" ]; then
      docker build  \
        -t zcharyma/zedisdb \
        .
  else
      echo "BUILD_ENV is set to ${BUILD_ENV}"
      docker build  \
        -t zcharyma/zedisdb \
        --build-arg "HTTP_PROXY=http://192.169.0.101:7890" \
        --build-arg "HTTPS_PROXY=http://192.169.0.101:7890" \
        --build-arg "NO_PROXY=localhost,127.0.0.1,.example.com" \
        .
  fi
}

build_image


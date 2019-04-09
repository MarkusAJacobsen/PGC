#!/usr/bin/env bash

 if [[ ! -d "./PGL" ]]; then
    git clone https://github.com/MarkusAJacobsen/PGL.git
  else
    cd ./PGL
    git pull origin master
    cd ..
 fi

cd ./deployment

docker-compose up --build

exit

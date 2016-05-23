#!/bin/sh
set -ex
#wget  https://github.com/google/protobuf/releases/download/v3.0.0-beta-2/protobuf-cpp-3.0.0-beta-2.tar.gz
#tar -xzvf *.tar.gz
#cd *protobuf* && ./configure --prefix=/usr && make && sudo make install
wget https://github.com/google/protobuf/releases/download/v3.0.0-beta-3/protoc-3.0.0-beta-3-linux-x86_64.zip
unzip protoc-3.0.0-beta-3-linux-x86_64.zip
sudo cp protoc /usr/bin

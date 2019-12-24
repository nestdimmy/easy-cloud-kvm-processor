#!/bin/bash

BUILD=iso
IMAGE=ubuntu-10.04-alternate-i386.iso

# Распаковываем образ в директорию
rm -rf $BUILD/
mkdir $BUILD/
echo "** Mounting image..."
sudo mount -o loop $IMAGE /mnt/
echo "** Syncing..."
rsync -av /mnt/ $BUILD/
chmod -R u+w $BUILD/
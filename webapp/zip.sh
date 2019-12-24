#!/bin/bash

IMAGE=ubuntu-custom.iso
BUILD=iso

# Запаковываем содержимое iso/ в образ ubuntu-custom.iso
echo ">>> Calculating MD5 sums..."
rm $BUILD/md5sum.txt
(cd $BUILD/ && find . -type f -print0 | xargs -0 md5sum | grep -v "boot.cat" | grep -v "md5sum.txt" > md5sum.txt)
echo ">>> Building iso image..."

mkisofs -r -V "Ubuntu OEM install" \
            -cache-inodes \
            -J -l -b isolinux/isolinux.bin \
            -c isolinux/boot.cat -no-emul-boot \
            -boot-load-size 4 -boot-info-table \
            -o $IMAGE $BUILD/
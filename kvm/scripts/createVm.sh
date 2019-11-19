sudo virt-install \
              --hvm \
              --name demo1 \
              --ram 500 \
              --nodisks \
              --livecd \
              --graphics vnc \
              --debug \
              --cdrom ./imgs/ubuntu.iso

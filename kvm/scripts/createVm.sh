sudo virt-install \
              --hvm \
              --name demo \
              --ram 500 \
              --nodisks \
              --livecd \
              --graphics vnc \
              --debug \
              --cdrom ubuntu.iso

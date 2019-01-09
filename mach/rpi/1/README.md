# Boot files and kernel
```
Copy files from github.com/raspberrypi/firmware/tree/stable/boot
to an sdcard (FAT32):
	/bcm2708-rpi-0-w.dtb         23K
	/bcm2708-rpi-b-plus.dtb      23K
	/bcm2708-rpi-b.dtb           23K
	/bcm2708-rpi-cm.dtb          23K
	/bcm2709-rpi-2-b.dtb         24K
	/bcm2710-rpi-3-b-plus.dtb    25K
	/bcm2710-rpi-3-b.dtb         25K
	/bcm2710-rpi-cm3.dtb         24K
	/bootcode.bin                51K
	/fixup_cd.dat               2.6K
	/kernel.img                 4.5M
	/kernel7.img                4.8M
	/start_cd.elf               663K
This includes files for all hardware models.
The _cd* files are the cut down versions for gpu_mem=16
Add
	/config.txt

$ cat config.txt
gpu_mem=16

Linux should now be able to boot into a nice Kernel panic in about 2-3s.
```

# Compile the userland
```
We need to cross compile:
$ GOOS=linux GOARCH=arm go build -o init

Put the result in an initramfs.
$ echo init | cpio -o -H newc | gzip > initramfs.gz
Requirements are:
- static binary
- owned by root
- executable bit must be set
This does not work when working on Windows to my knowledge.
Alternative: 
$ go run github.com/ktye/usr/mkinit

$ cat config.txt
gpu_mem=16
initramfs initramfs.gz 0x00800000

The init program will be in /init instead of the default /sbin/init.
Tell the kernel about that with a parameter:
$ cat cmdline.txt
init=/init

	
Ref: www.raspberrypi.org/documentation/configuration/boot_folder.md
Ref: www.raspberrypi.org/documentation/hardware/raspberrypi/bootmodes/
Ref: landley.net/writing/rootfs-howto.html
```
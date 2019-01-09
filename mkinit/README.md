# mkinit creates an cpio newc archive of a single file

```
// Mkinit creates an cpio newc archive of a single file
//
// Usage:
//	mkinit
//
// This is equivalent to do
//	echo init | cpio -o -H newc | gzip > initfs.gz
//
// The input file must be called'init' and the archive is
// written to initfs.gz
//
// Mkinit ignores the file attributes and sets mode to 555
// and owner to root.
//
// Note: the result is an initramfs with the file 'init' in the root
// directory and not the linux kernel default /sbin/init
```

## Reference
This is based on: `github.com/u-root/u-root/pkg/cpio/newc.go`

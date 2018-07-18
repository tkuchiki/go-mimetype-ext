# go-mimetype-ext

Mimetype utilities for Golang

## Usage

```golang
// []string{"png"}, nil
ext, err := mimetype_ext.GetExtension("image/png")

// []string{"image/png", "image/x-citrix-png", "image/x-png"}, nil
mtype, err := mimetype_ext.GetMimetype("png")
mtype, err := mimetype_ext.GetMimetype(".png")
```

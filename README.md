# go-whosonfirst-mimetypes

There are many mime-type lookup tables. This one is ours.

## Important

This is not ready for you.

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.7 so let's just assume you need [Go 1.9](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## See also

* https://golang.org/pkg/mime/#ExtensionsByType
* https://golang.org/pkg/net/http/#DetectContentType
* https://svn.apache.org/viewvc/httpd/httpd/branches/2.2.x/docs/conf/mime.types?view=co
* https://www.iana.org/assignments/media-types/media-types.xhtml

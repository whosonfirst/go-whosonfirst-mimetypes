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

## Tools

### wof-mimetype-lookup

A simple command line tool to lookup mimetypes by extension or vice versa.

```
./bin/wof-mimetype-lookup -h
Usage of ./bin/wof-mimetype-lookup:
  -extension
    	Lookup mimetypes by extension
  -mimetype
    	Lookup extensions by mimetype
```

The output is a line-separated list containing a tab-separated list of input (extension or mimetype) followed by one or more matched. For example:

```
./bin/wof-mimetype-lookup -mimetype '.json' jpg
.json	application/json
jpg	image/jpeg

./bin/wof-mimetype-lookup -extension 'image/jpeg' 'image/gif'
image/jpeg	jpeg	jpg	jpe
image/gif	gif
```

## See also

* https://golang.org/pkg/mime/#ExtensionsByType
* https://golang.org/pkg/net/http/#DetectContentType
* https://svn.apache.org/viewvc/httpd/httpd/branches/2.2.x/docs/conf/mime.types?view=co
* https://www.iana.org/assignments/media-types/media-types.xhtml

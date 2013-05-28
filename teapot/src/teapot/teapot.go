/**
 * Teapot Library - used by annotated media collation (AMC) server (teapotd)
 * User: Nyk Cowham <nyk@demotix.com>
 * Date: 5/25/13 AD
 */
package teapot

import "net/http"

type Annotator interface {
	Annotate(string, string)
}

/**
 * Annotations are key/value metadata fields. In terms of HTML form semantics they correspond to form fields.
 *
 * Both MediaFile and Collation objects may contain Annotations.
 */
type Annotations map[string] string

/**
 * MediaFiles are collections (slices) of MediaFile objects
 */
type MediaFiles map[string] *MediaFile

/**
 * MediaFile is a representation of an uploaded file complete with it's annotations
 */
type MediaFile struct {
	file io.File
	annotations *Annotations
}

type Collation struct {
	media MediaFiles
	annotations *Annotations
}


func (c *Collation) GetMediaByKey(key string) *MediaFile {
	return c.media[key]
}

func (c *Collation) Annotate(key, value string) {
	c.annotations[key] = value
}

func (f *MediaFile) Annotate(key, value string) {
	f.annotations[key] = value
}


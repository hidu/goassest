/**
* generated by goassest
* https://github.com/hidu/goassest/
 */
package demo

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type StaticFile struct {
	Name    string
	Mtime   int64
	Content string
}

type StaticFiles map[string]*StaticFile

//for debug
var DebugAssestDir string = ""

func (statics StaticFiles) GetStaticFile(name string) (*StaticFile, error) {
	if DebugAssestDir != "" {
		return getStaticFile(DebugAssestDir, name)
	}
	if sf, has := statics[path.Clean(name)]; has {
		return sf, nil
	}
	return nil, fmt.Errorf("not exists")
}

func getStaticFile(baseDir string, name string) (*StaticFile, error) {
	fullPath := baseDir + string(filepath.Separator) + name
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if info.Mode().IsRegular() {
		content, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return &StaticFile{
			Content: string(content),
			Name:    name,
			Mtime:   info.ModTime().Unix(),
		}, nil
	}
	return nil, fmt.Errorf("not file")
}

/**
*res.DebugAssestDir="../";//for debug read it direct
*http.Handle("/res/",res.Files.HttpHandler("/res/"))
 */
func (statics *StaticFiles) HttpHandler(pdir string) http.Handler {
	return &fileServer{sf: statics, pdir: pdir}
}

func decode(data string) string {
	b, _ := base64.StdEncoding.DecodeString(data)
	gr, _ := gzip.NewReader(bytes.NewBuffer(b))
	bs, _ := ioutil.ReadAll(gr)
	return string(bs)
}

func base64decode(data string) string {
	b, _ := base64.StdEncoding.DecodeString(data)
	return string(b)
}

type fileServer struct {
	sf   *StaticFiles
	pdir string
}

func (f *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rname, _ := filepath.Rel(f.pdir, r.URL.Path)
	static, err := f.sf.GetStaticFile(rname)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	modtime := time.Unix(static.Mtime, 0)
	modifiedSince := r.Header.Get("If-Modified-Since")
	if modifiedSince != "" {
		t, err := time.Parse(http.TimeFormat, modifiedSince)
		if err == nil && modtime.Before(t.Add(1*time.Second)) {
			h := w.Header()
			delete(h, "Content-Type")
			delete(h, "Content-Length")
			w.Header().Set("Last-Modified", modtime.UTC().Format(http.TimeFormat))
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	mimeType := mime.TypeByExtension(filepath.Ext(static.Name))
	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
	w.Header().Set("Last-Modified", modtime.UTC().Format(http.TimeFormat))
	w.Write([]byte(static.Content))
}

var Files = StaticFiles{

	base64decode("UkVBRE1FLm1k"): &StaticFile{
		Name:    base64decode("UkVBRE1FLm1k"),
		Mtime:   1417276377,
		Content: decode("H4sIAAAJbogA/0rPTywuTi0uUUhJzc3n4oJw9NLzFTKLFdJT81KLEktSUxSSKhVgCrkSuJJTFHITM/MU1NSAogpFpXlgLkhTtG5KalJpeixXAgAAAP//AQAA//+9SujXWAAAAA=="),
	},

	base64decode("Yi5jc3M="): &StaticFile{
		Name:    base64decode("Yi5jc3M="),
		Mtime:   1417275726,
		Content: decode("H4sIAAAJbogA/8owrE7LzyvRLc6sSrVSMLEoqLBOzs/JL7JKL0pNzavl4sowroYIFKWm1AIAAAD//wEAAP//9lO6yS4AAAA="),
	},

	base64decode("aW5kZXguaHRtbA=="): &StaticFile{
		Name:    base64decode("aW5kZXguaHRtbA=="),
		Mtime:   1417275823,
		Content: decode("H4sIAAAJbogA/5RROVPzMBDt8ys26h19OYqvkNVwtFCEgorZ2EukRLY80g7gf8/6YMZQQbXjt/sOP5n17cPN8fnxDhw3wa7M1yCsZTTECJXDlIlL9XS8L/4rgdlzIOsohAiYM2U2esJWJvj2ColCqTL3gbIjYgXcd1Qqpg/WVc4KXKLXUunTZvgSVq6S73h5dsE3nFAFOVVy3KBvNW4uQjB6Wv2FuWCBWRcFHP4doCh+KdHhmYSwFPmhoufOTrHu7Uoq3M4NvccUallvZdlZg99/HhjTeWj35RSwvSo7wkajOHWiI0I7e45TzVBTE0VqNxrswdelGk2GXG5vx0RDRVDFlsUXPM/pplhyNTzwJwAAAP//AQAA//+NTqtV9wEAAA=="),
	},

	base64decode("bWFpbi9hLmpz"): &StaticFile{
		Name:    base64decode("bWFpbi9hLmpz"),
		Mtime:   1417275667,
		Content: decode("H4sIAAAJbogA/+IqTi0JycxNzS8t0UgrzUsuyczP09Cs5uJMyU8uzU3NK9FLTy1xzUkFMZ0qPVM0lDJSc3LylTT1MvPyUos8Qnx9bCFCOol6WcVK1ly1OoYGBgaa1gAAAAD//wEAAP//cUsdmloAAAA="),
	},

	base64decode("bWFpbi9tYWluLmdv"): &StaticFile{
		Name:    base64decode("bWFpbi9tYWluLmdv"),
		Mtime:   1417275564,
		Content: decode("H4sIAAAJbogA/yxPX2vCMBB/bj7FERikKqnuZSKM4ZDhYA+CnyAzVw1L0pJEGQy/++5S6cPR/P6P5vRjzgjBuCiEC+OQCijRSK07Saf35lxvKHwilu5SyihFK8TNJLD4fT3DKzBPvw+DV7I+yQX0xmdcgLxmBJMz5gLWJTyVN9lWbc16SD9jUZIfSLherpek4xzImG44UUkl+ms81a6qhT/RVOnBpIyqFY3rYTb1IaixGAa9499tDd+5RGGPXXfRsL3em2g9KtlRbBV8OI9Z7wmboMRYS+bG2gQbahuKPo7JxdIruXp+0Uv6VpsnSw4zrklc5hyY4iON8hluLrsCnLjpOglzqG5zkJ2LFn/1pQRP86ZOXy4XjNtoj7xdMXUB0flW3MU/AAAA//8BAAD//yWcFpOvAQAA"),
	},
}

package server

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/dyweb/gommon/errors"
	"gopkg.in/yaml.v2"

	"github.com/at15/goyourcassandra/pkg/types"
)

func (srv *Server) handleGetBookmarks(w http.ResponseWriter, r *http.Request) {
	home := srv.cfg.Home
	log.Infof("list bookmarks, home is %s", home)
	files, err := ioutil.ReadDir(home)
	if err != nil {
		write500(w, errors.Wrap(err, "error list bookmark files from folder"))
		return
	}
	bookmarks := make(map[string]types.Bookmark, len(files))
	for _, file := range files {
		log.Infof("found file %s", file.Name())
		switch filepath.Ext(file.Name()) {
		case ".yaml", ".yml":
			// do nothing
		default:
			log.Infof("skip file ext %s", filepath.Ext(file.Name()))
			continue // skip other files
		}
		p := filepath.Join(home, file.Name())
		bookMarkName := file.Name()
		log.Infof("read file %s", p)
		bk, err := readBookmark(p)
		if err != nil {
			write500(w, errors.Wrap(err, "error read bookmark"))
			return
		}
		log.Infof("bookmark host %s #templates %d", bk.Host, len(bk.Templates))
		bookmarks[bookMarkName] = bk
	}
	writeJSON(w, bookmarks)
}

func readBookmark(p string) (types.Bookmark, error) {
	var bk types.Bookmark
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return bk, errors.Wrap(err, "error read bookmark file")
	}
	if err := yaml.UnmarshalStrict(b, &bk); err != nil {
		return bk, errors.Wrap(err, "error decode bookmark as yaml")
	}
	return bk, nil
}

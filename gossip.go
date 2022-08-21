package gossip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Gossip struct {
	Comment string
	Count   int
	byName  map[string]*Entry
	byIndex []*Entry
}

type Entry struct {
	Comment string
	Name    string
}

func create() *Gossip {
	g := new(Gossip)
	g.Comment = ""
	g.Count = 0
	g.byName = map[string]*Entry{}
	g.byIndex = []*Entry{}
	return g
}

func endWith(s, suffix string) bool {
	idx := len(s) - len(suffix)
	if s[idx:] == suffix {
		return true
	} else {
		return false
	}
}

func Extract(zipPath, to string) error {
	var err error
	if _, err = os.Lstat(to); !os.IsNotExist(err) {
		err = os.RemoveAll(to)
		if err != nil {
			return err
		}
	}
	err = os.Mkdir(to, 0775)
	if err != nil {
		return err
	}
	z, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	for _, e := range z.File {
		path := filepath.Join(to, e.FileHeader.Name)
		if endWith(e.FileHeader.Name, "/") {
			if _, err = os.Lstat(path); os.IsNotExist(err) {
				err = os.MkdirAll(path, 0775)
				if err != nil {
					return err
				}
			}
			continue
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		rc, err := e.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func Open(path string) (*Gossip, error) {
	g := create()
	z, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer z.Close()
	g.Comment = z.Comment

	for _, rc := range z.File {
		e := new(Entry)
		e.Name = rc.FileHeader.Name
		e.Comment = rc.Comment
		g.byName[e.Name] = e
		g.byIndex = append(g.byIndex, e)
		g.Count++
	}

	return g, nil
}

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
)

var (
	reName = regexp.MustCompile(`([a-zA-Z0-9_]+).json`)
)

type FileInfo struct {
	Name  string
	Path  string
	Route string
}

func fileInfo(filepath string) (FileInfo, error) {

	var fi FileInfo

	sm := reName.FindStringSubmatch(filepath)
	if len(sm) < 2 {
		return fi, errors.New("can't extract name from file")
	}
	fi.Name = sm[1]
	fi.Path = filepath
	fi.Route = fmt.Sprintf("/%s", sm[1])

	return fi, nil
}

func main() {
	// cli args
	port := flag.String("port", "8001", "port to run server on")
	dbDir := flag.String("dir", "db", "folder containing json models")

	// all model file names
	ms, err := filepath.Glob(fmt.Sprintf("%s/*.json", *dbDir))
	if err != nil {
		log.Fatal(err)
	}
	var ps []string
	for _, m := range ms {
		var p json.RawMessage // json.RawMessage validates json a a side effect
		f, err := ioutil.ReadFile(m)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(f, &p); err != nil {
			log.Fatal(err)
		}
		fi, err := fileInfo(m)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(fi.Route)

		// generating route for each model
		http.HandleFunc(fi.Route, func(w http.ResponseWriter, r *http.Request) {

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(p)

		})
		ps = append(ps, fi.Name)
	}

	log.Println("running server with following paths")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

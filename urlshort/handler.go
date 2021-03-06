package urlshort

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		if val, ok := pathsToUrls[url]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

type path struct {
	Path string `yaml:"path"`
	Dest string `yaml:"url"`
}

func parseYaml(yml []byte) ([]path, error) {

	var data []path
	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildMap(list []path) map[string]string {
	m := map[string]string{}
	for item := range list {
		m[list[item].Path] = list[item].Dest
	}
	return m
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	data, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	m := buildMap(data)

	return MapHandler(m, fallback), nil
}

/*
Pass the function a database connection and then it
queries for the path

Table schema is expected to be the following

CREATE TABLE paths if not exists (
	source varchar(255),
	dest varchar(255)
)
*/
func DBHandler(conn *sql.DB, fallback http.Handler) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query, err := conn.Query(fmt.Sprintf("select * from paths where source = %s", r.URL.String()))

		if err != nil {
			return
		}

		result := path{}
		if err := query.Next(); !err {
			fallback.ServeHTTP(w, r)
			return
		}

		query.Scan(&result.Path, &result.Dest)
		http.Redirect(w, r, result.Dest, http.StatusFound)
		return
	}), nil
}

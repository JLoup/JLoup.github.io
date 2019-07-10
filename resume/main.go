package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	TEMPLATE_FILE = "template.mustache"
	RESUME_FILE   = "resume.json"
)

func MergeMaps(src, dst map[string]interface{}) map[string]interface{} {
	for key, value := range src {
		dst[key] = value
	}
	return dst
}

func GetResume(path string) (map[string]interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		// it is ok if resume file does not exist
		if os.IsNotExist(err) {
			return make(map[string]interface{}), nil
		} else {
			return nil, err
		}
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func ProcessDir(dirname string, globalJSON map[string]interface{}) (string, error) {
	localJSON, err := GetResume(filepath.Join(dirname, RESUME_FILE))
	if err != nil {
		return "", err
	}

	localJSON = MergeMaps(globalJSON, localJSON)

	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return "", err
	}

	for _, dir := range dirs {

		dir, _ = os.Stat(filepath.Join(dirname, dir.Name()))

		if dir.IsDir() {
			str, err := ProcessDir(filepath.Join(dirname, dir.Name()), localJSON)
			if err != nil {
				return "", err
			}
			localJSON[dir.Name()] = str
		}

	}

	return mustache.RenderFile(filepath.Join(dirname, TEMPLATE_FILE), localJSON), nil
}

func main() {
	dir := flag.String("r", ".", "resume data directory")
	baseUrl := flag.String("baseurl", ".", "root url")

	flag.Parse()

	*dir = strings.TrimSuffix(*dir, "/")

	baseJson := map[string]interface{}{"base": *baseUrl}
	str, err := ProcessDir(*dir, baseJson)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", str)
}

package phpcfg

import "io/ioutil"

func ParsePath(path string) (map[string]string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(body)
}

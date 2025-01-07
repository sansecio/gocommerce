package phpcfg

import "os"

func ParsePath(path string) (map[string]string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(body)
}

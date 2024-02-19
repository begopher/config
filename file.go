package vars

import (
	"bufio"
	"os"
	"strings"
)

// File is parsed with the following rules:
//  - empty lines are ignored.
//  - lines start with # are ignored (comments).
//  - lines start with = are ignored (no key).
//  - lines does not have equal symbol (=) are ignored.
//  - the first equal symbol (=) in a line acts as a delimiter for key and value pair.
//  - last identical key override previous value.
//  - key is associated with empty string if value is omitted.
func File(path string) (Source, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	vars := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "=") {
			continue
		}
		if !strings.Contains(line, "=") {
			continue
		}
		keyVal := strings.SplitN(line, "=", 2)
		if len(keyVal) == 1 {
			key := strings.TrimSpace(keyVal[0])
			vars[key] = ""
			continue
		}
		key := strings.TrimSpace(keyVal[0])
		val := strings.TrimSpace(keyVal[1])
		vars[key] = val
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return Map(vars), nil
}

package kindle

import (
	"os/exec"
	"strings"
)

// RawLIPC fetches a property using LIPC. You should not be using this.
func RawLIPC(service, property string) (string, error) {
	cmd := exec.Command("/usr/bin/lipc-get-prop", service, property)
	res, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	str := string(res)
	str = strings.TrimSpace(str)

	return str, nil
}

// RawLIPCHashArray fetches a HashArray property using LIPC. You should not be using this.
func RawLIPCHashArray(service, property string) (hash []map[string]string, err error) {
	cmd := exec.Command("/usr/bin/lipc-hash-prop", "-n", service, property)

	var res []byte
	res, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	str := string(res)
	lines := strings.Split(str, "\n")

	var currentHash = make(map[string]string)
	for i := 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Hash Index:") {
			// New entry
			hash = append(hash, currentHash)
			currentHash = make(map[string]string)
			continue
		}

		line := strings.TrimSpace(lines[i])
		parts := strings.SplitN(line, " = ", 2)

		key := parts[0]
		value := parts[1][1 : len(parts[1])-1] // remove encasing "" or ()

		currentHash[key] = value
	}
	hash = append(hash, currentHash)

	return
}

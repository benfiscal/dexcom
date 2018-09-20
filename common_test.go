package dexcom

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// Force timezone to match test data.
func init() {
	os.Setenv("TZ", "America/New_York")
}

func readBytes(r io.Reader) ([]byte, error) {
	var data []byte
	for {
		var b byte
		n, err := fmt.Fscanf(r, "%02x", &b)
		if n == 0 {
			break
		}
		if err != nil {
			return data, err
		}
		data = append(data, b)
	}
	return data, nil
}

func parseTime(s string) time.Time {
	t, err := time.ParseInLocation(UserTimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	return t
}

func compareJSON(data interface{}, jsonFile string) (bool, string) {
	// Write data in JSON format to temporary file.
	tmpfile, err := ioutil.TempFile("", "json")
	if err != nil {
		return false, err.Error()
	}
	defer func() { _ = os.Remove(tmpfile.Name()) }()
	e := json.NewEncoder(tmpfile)
	e.SetIndent("", "  ")
	err = e.Encode(data)
	_ = tmpfile.Close()
	if err != nil {
		return false, err.Error()
	}
	// Write JSON in canonical form for comparison.
	canon1 := canonicalJSON(jsonFile)
	canon2 := canonicalJSON(tmpfile.Name())
	// Find differences.
	cmd := exec.Command("diff", "-u", "--label", jsonFile, "--label", "decoded", canon1, canon2)
	diffs, err := cmd.Output()
	_ = os.Remove(canon1)
	_ = os.Remove(canon2)
	return err == nil, string(diffs)
}

// canonicalJSON reads the given file and creates a temporary file
// containing equivalent JSON in canonical form
// (using the "jq" command, which must be on the user's PATH).
// It returns the temporary file name; it is the caller's responsibility
// to remove it when done.
func canonicalJSON(file string) string {
	canon, err := exec.Command("jq", "-S", ".", file).Output()
	if err != nil {
		panic(err)
	}
	tmpfile, err := ioutil.TempFile("", "json")
	if err != nil {
		panic(err)
	}
	_, _ = tmpfile.Write(canon)
	_ = tmpfile.Close()
	return tmpfile.Name()
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
)

type processStat struct {
	UID, PID, PPID  int64
	UTIME, STIME    float64
	USER, EXEC_NAME string
	MEMORY          uint64
}

func readDirectories(path string) []os.FileInfo {

	data, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	return data

}

func getProcessStatus(dir string) *processStat {

	var userIs string
	uid, err := ioutil.ReadFile("/proc/" + dir + "/loginuid")
	if err != nil {
		panic(err)
	}

	// User look up using UID
	uidCheck, _ := user.LookupId(string(uid))
	if uidCheck == nil {
		// if loginuid is not set
		userIs = "root"
	} else {
		userIs = uidCheck.Username
	}

	dirData, err := ioutil.ReadFile("/proc/" + dir + "/stat")
	if err != nil {
		panic(err)
	}

	data := strings.Fields(string(dirData))

	var processData processStat

	processData.USER = userIs
	processData.UID, _ = strconv.ParseInt(string(uid), 10, 64)
	execName := data[1]
	processData.EXEC_NAME = execName[1 : len(execName)-1]

	processData.PID, _ = strconv.ParseInt(data[0], 10, 64)
	processData.PPID, _ = strconv.ParseInt(data[3], 10, 64)
	processData.MEMORY, _ = strconv.ParseUint(data[22], 10, 64)

	processData.UTIME, _ = strconv.ParseFloat(data[13], 64) // #14 -> utime
	processData.UTIME = processData.UTIME / 100.00          // to convert in seconds

	processData.STIME, _ = strconv.ParseFloat(data[14], 64) // #15 -> stime
	processData.STIME = processData.STIME / 100.00          // to convert in seconds

	return &processData
}

func main() {

	var dirNames []string

	// regexp for matching with numerical directory name
	re := regexp.MustCompile("^[0-9]+$")

	allDirectories := readDirectories("/proc")
	for _, dir := range allDirectories {
		name := dir.Name()

		// directories with matching regexp
		if re.MatchString(name) {
			dirNames = append(dirNames, name)
		}

	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "PID\tUSER\tEXECUTABLE NAME\tPPID\tMEMORY\tUTIME\tSTIME \n")

	for _, dirName := range dirNames {

		process := getProcessStatus(dirName)
		fmt.Fprintln(w, process.PID, "\t", process.USER, "\t", process.EXEC_NAME, "\t", process.PPID, "\t", process.MEMORY, "B", "\t", process.UTIME, "s", "\t", process.STIME, "s", "\t")

	}
	w.Flush()
}

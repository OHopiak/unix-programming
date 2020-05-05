package linux_programming

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

var (
	adList = []string{"a", "b", "c", "d"}
	cdList = []string{"c", "d"}
	jlList = []string{"j", "k", "l"}
)

//noinspection SpellCheckingInspection
var (
	wordCount = map[string]int64{
		"scripta":    13,
		"expetendis": 12,
		"affert":     11,
		"doming":     11,
		"natum":      11,
		"omittam":    11,
		"percipit":   11,
		"esse":       10,
		"fabellas":   10,
		"falli":      10,
	}
)

func runTestTask1(t *testing.T, mainCodeBlock func(dir string)) {
	dir, err := ioutil.TempDir(".", "test_data_")
	require.NoError(t, err, "failed to create a temporary directory")
	//t.Log("Created test dir")

	fillTestDirectoryTask1(t, dir)
	mainCodeBlock(dir)
	if !t.Failed() {
		checkExpectedDirectoryTask1(t, dir)
	}

	err = os.RemoveAll(dir)
	require.NoError(t, err, "failed to delete the temporary directory")
}

func fillTestDirectoryTask1(t *testing.T, dir string) {
	/*
		touch {a..d}.file.{0..10}
		mkdir -p {a..d}.folder
		touch {c..d}.folder/file.{0..3}
		touch {j..l}
	*/

	for _, i := range adList {
		// make files
		for j := 0; j <= 10; j++ {
			filename := fmt.Sprintf("%s/%s.file.%d", dir, i, j)
			f, err := os.Create(filename)
			_ = os.Chmod(filename, 0000)
			if err != nil {
				t.Error(err, "Failed to create a test file", filename)
				return
			}
			f.Close()
		}

		// make dirs
		err := os.Mkdir(fmt.Sprintf("%s/%s.folder", dir, i), 0755)
		if err != nil {
			t.Error(err, "Failed to create a test dir")
			return
		}
	}
	//t.Log("Made main test files")

	for _, i := range cdList {
		for j := 0; j <= 3; j++ {
			filename := fmt.Sprintf("%s/%s.folder/file.%d", dir, i, j)
			f, err := os.Create(filename)
			if err != nil {
				t.Error(err, "Failed to create a test file", filename)
				return
			}
			f.Close()
		}
	}

	for _, i := range jlList {
		// make files
		filename := fmt.Sprintf("%s/%s", dir, i)
		f, err := os.Create(filename)
		if err != nil {
			t.Error(err, "Failed to create a test file", filename)
			return
		}
		f.Close()
	}
	//t.Log("Made additional test files")
}

func checkExpectedDirectoryTask1(t *testing.T, dir string) {
	for _, i := range adList {
		// make files
		for j := 0; j <= 10; j++ {
			filename := fmt.Sprintf("%s/%s/%s.file.%d", dir, i, i, j)
			filenameBefore := fmt.Sprintf("%s/%s.file.%d", dir, i, j)

			assert.NoFileExists(t, filenameBefore, "file %s still exists", filenameBefore)
			assert.FileExists(t, filename, "file %s was not moved to %s", filenameBefore, filename)
		}

		dirName := fmt.Sprintf("%s/%s/%s.folder", dir, i, i)
		dirNameBefore := fmt.Sprintf("%s/%s.folder", dir, i)

		assert.NoDirExists(t, dirName, "directory %s was moved, but shouldn't be", dirNameBefore)
		assert.DirExists(t, dirNameBefore, "directory %s was moved, but shouldn't be", dirNameBefore)
	}
	for _, i := range cdList {
		for j := 0; j <= 3; j++ {
			filename := fmt.Sprintf("%s/%s.folder/file/file.%d", dir, i, j)
			filenameBefore := fmt.Sprintf("%s/%s.folder/file.%d", dir, i, j)
			assert.NoFileExists(t, filename, "file %s was moved, but shouldn't be", filenameBefore)
			assert.FileExists(t, filenameBefore, "file %s was moved, but shouldn't be", filenameBefore)
		}
	}
	for _, i := range jlList {
		// make files
		filename := fmt.Sprintf("%s/%s/%s", dir, i, i)
		filenameBefore := fmt.Sprintf("%s/%s", dir, i)
		assert.NoFileExists(t, filename, "file %s was moved, but shouldn't be", filenameBefore)
		assert.FileExists(t, filenameBefore, "file %s was moved, but shouldn't be", filenameBefore)
	}
}

func runTestTask2(t *testing.T, mainCodeBlock func(inputFile string) string) {
	inputFile := "test_file.txt"
	output := mainCodeBlock(inputFile)
	if !t.Failed() {
		checkOutputTask2(t, output)
	}
}

func checkOutputTask2(t *testing.T, output string) {
	rows := strings.Split(strings.TrimSpace(output), "\n")
	assert.Len(t, rows, 10, "invalid output format")
	for _, row := range rows {
		items := strings.Split(strings.TrimSpace(row), " ")
		require.Len(t, items, 2, "invalid row format, two columns expected")
		count, err := strconv.ParseInt(items[0], 10, 64)
		require.NoError(t, err, "invalid row format, the first column is count")
		expectedCount, ok := wordCount[items[1]]
		require.True(t, ok, "word '%s' not expected", items[1])
		assert.Equal(t, expectedCount, count, "incorrect count")
	}
}

func checkAppExecutionTask1(t *testing.T, dir string, command string, args ...string) {
	cmd := exec.Command(command, append(args, "--help")...)
	output, err := cmd.CombinedOutput()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, output, "'<cmd> --help' should return the help info")
	}
	err = exec.Command(command, append(args, "a")...).Run()
	assert.Error(t, err, "should fail on not existing directory")

	err = exec.Command(command, append(args, dir, "a")...).Run()
	assert.Error(t, err, "should fail on multiple params")

	cmd = exec.Command(command, append(args, dir)...)
	err = cmd.Run()
	//output, err = cmd.CombinedOutput()
	//t.Log(string(output))
	assert.NoError(t, err, "should run successfully for the existing directory")
}

func checkAppExecutionTask2(t *testing.T, inputFile string, command string, args ...string) string {
	cmd := exec.Command(command, append(args, "--help")...)
	output, err := cmd.CombinedOutput()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, output, "'<cmd> --help' should return the help info")
	}
	err = exec.Command(command, append(args, "a")...).Run()
	assert.Error(t, err, "should fail on not existing file")

	err = exec.Command(command, append(args, inputFile, "a")...).Run()
	assert.Error(t, err, "should fail on multiple params")

	c := exec.Command(command, append(args, inputFile)...)
	output, err = c.CombinedOutput()
	assert.NoError(t, err, "should run successfully for the existing file")
	//t.Log("Executed the app to process the folder")
	return string(output)
}

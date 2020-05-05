package linux_programming

import (
	"github.com/stretchr/testify/require"
	"os/exec"
	"testing"
)

/*
QUndoStack
Log
QConnect to database
QStyle
*/

func TestBash_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		command := "./bash/t1/restructure_folder.sh"
		checkAppExecutionTask1(t, dir, command)
	})
}

func TestBash_Task2(t *testing.T) {
	t.Parallel()
	runTestTask2(t, func(inputFile string) string {
		command := "./bash/t2/convert_file.sh"
		output := checkAppExecutionTask2(t, inputFile, command)
		return output
	})
}

func TestCpp_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		cmd := exec.Command("cmake", "..")
		cmd.Dir = "./gcc/cmake-build-debug"
		require.NoError(t, cmd.Run(), "failed to build cmake")

		cmd = exec.Command("make", "task1")
		cmd.Dir = "./gcc/cmake-build-debug"
		require.NoError(t, cmd.Run(), "failed to build task1")

		command := "./gcc/cmake-build-debug/task1"
		checkAppExecutionTask1(t, dir, command)
	})
}

func TestCpp_Task2(t *testing.T) {
	t.Parallel()
	runTestTask2(t, func(inputFile string) string {
		cmd := exec.Command("cmake", "..")
		cmd.Dir = "./gcc/cmake-build-debug"
		require.NoError(t, cmd.Run(), "failed to build cmake")

		cmd = exec.Command("make", "task2")
		cmd.Dir = "./gcc/cmake-build-debug"
		require.NoError(t, cmd.Run(), "failed to build task1")

		command := "./gcc/cmake-build-debug/task2"
		output := checkAppExecutionTask2(t, inputFile, command)
		return output
	})
}

func TestPython_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		command := "./python/task1.py"
		checkAppExecutionTask1(t, dir, command)
	})
}

func TestPython_Task2(t *testing.T) {
	t.Parallel()
	runTestTask2(t, func(inputFile string) string {
		command := "./python/task2.py"
		output := checkAppExecutionTask2(t, inputFile, command)
		return output
	})
}

func TestPerl_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		command := "./perl/task1.pl"
		checkAppExecutionTask1(t, dir, command)
	})
}

func TestPerl_Task2(t *testing.T) {
	t.Parallel()
	runTestTask2(t, func(inputFile string) string {
		command := "./perl/task2.pl"
		output := checkAppExecutionTask2(t, inputFile, command)
		return output
	})
}

func TestJava_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		cmd := exec.Command("./gradlew", "build")
		cmd.Dir = "./java/"
		output, err := cmd.CombinedOutput()
		require.NoError(t, err, "failed to build gradle:", string(output))

		command := "java"
		args := []string{
			"-cp", "./java/build/classes/java/main/",
			"tasks.task1.Main",
		}
		checkAppExecutionTask1(t, dir, command, args...)
	})
}

func TestJava_Task2(t *testing.T) {
	t.Parallel()
	runTestTask2(t, func(inputFile string) string {
		cmd := exec.Command("./gradlew", "build")
		cmd.Dir = "./java/"
		combinedOut, err := cmd.CombinedOutput()
		require.NoError(t, err, "failed to build gradle:", string(combinedOut))

		command := "java"
		args := []string{
			"-cp", "./java/build/classes/java/main/",
			"tasks.task2.Main",
		}

		output := checkAppExecutionTask2(t, inputFile, command, args...)
		return output
	})
}

func TestDotnet_Task1(t *testing.T) {
	t.Parallel()
	runTestTask1(t, func(dir string) {
		cmd := exec.Command("dotnet", "publish", "-c", "release", "-r", "linux-x64")
		cmd.Dir = "./dotnet/"
		output, err := cmd.CombinedOutput()
		require.NoError(t, err, "failed to build gradle:", string(output))

		command := "./dotnet/bin/release/netcoreapp3.1/linux-x64/publish/dotnet"
		checkAppExecutionTask1(t, dir, command)
	})
}
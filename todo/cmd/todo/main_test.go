package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s: %s\n", binName, err)
		os.Exit(1)
	}
	fmt.Println("running tests...")
	result := m.Run()

	fmt.Println("cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	require.NoError(t, err)

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		err := cmd.Run()
		require.NoError(t, err)
	})
	task2 := "test task number 2"
	task3 := "test task number 3"
	t.Run("AddNewTasksFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		require.NoError(t, err)
		io.WriteString(cmdStdIn, task2)
		io.WriteString(cmdStdIn, task3)
		io.WriteString(cmdStdIn, "")
		cmdStdIn.Close()
		err = cmd.Run()
		require.NoError(t, err)
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		require.NoError(t, err)

		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task, task2)
		require.Equal(t, expected, string(out), fmt.Errorf("expected %q, got %q instead\n", expected, string(out)))
	})
}

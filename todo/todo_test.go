package todo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/koneal2013/cli-tools/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	require.Equal(t, l[0].Task, taskName, fmt.Errorf("expected %q, got %q instead", taskName, l[0].Task))
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	require.Equal(t, taskName, l[0].Task, fmt.Errorf("expected %q, got %q instead", taskName, l[0].Task))
	require.False(t, l[0].Done, "new task should not be completed")
	l.Complete(1)
	require.True(t, l[0].Done, "new task should be completed")
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	require.Equal(t, tasks[0], l[0].Task, fmt.Errorf("expected %q, got %q instead", tasks[0], l[0].Task))
	l.Delete(2)
	require.Equal(t, 2, len(l), fmt.Errorf("expected list lenght %d, got %d instead", 2, len(l)))
	require.Equal(t, tasks[2], l[1].Task, fmt.Errorf("expected %q, got %q instead", tasks[2], l[1].Task))
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	taskName := "New Task"
	l1.Add(taskName)
	require.Equal(t, taskName, l1[0].Task, fmt.Errorf("expected %q, got %q instead", taskName, l1[0].Task))
	tf, err := os.CreateTemp("", "")
	require.NoError(t, err, fmt.Errorf("error creating temp file: %s", err))
	defer os.Remove(tf.Name())
	err = l1.Save(tf.Name())
	require.NoError(t, err, fmt.Errorf("error saving list to fiile: %s", err))
	err = l2.Get(tf.Name())
	require.NoError(t, err, fmt.Errorf("error getting list from fiile: %s", err))
	require.Equal(t, l1[0].Task, l2[0].Task, fmt.Errorf("task %q should match %q task", l1[0].Task, l2[0].Task))
}

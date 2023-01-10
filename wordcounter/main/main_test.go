package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	actual, err := count(b, false, false)
	require.NoError(t, err)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 \nline2\nline3")

	expected := 3
	actual, err := count(b, true, false)
	require.NoError(t, err)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 \nline2\nline3")

	expected := b.Len()
	actual, err := count(b, false, true)
	require.NoError(t, err)

	if expected != actual {
		t.Errorf("Expected %d, got %d instead", expected, actual)
	}
}

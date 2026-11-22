package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadIntegers(t *testing.T) {
	filename := writeTestCSV(t, "value\n5\n1\n3\n")

	got, err := readIntegers(filename)
	if err != nil {
		t.Fatalf("readIntegers returned an unexpected error: %v", err)
	}

	want := []int{5, 1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("readIntegers() = %v; want %v", got, want)
	}
}

func TestReadFloats(t *testing.T) {
	filename := writeTestCSV(t, "value\n5.25\n1.50\n3.75\n")

	got, err := readFloats(filename)
	if err != nil {
		t.Fatalf("readFloats returned an unexpected error: %v", err)
	}

	want := []float64{5.25, 1.50, 3.75}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("readFloats() = %v; want %v", got, want)
	}
}

func TestReadIntegersRejectsInvalidValue(t *testing.T) {
	filename := writeTestCSV(t, "value\n1\nnot-an-integer\n3\n")

	_, err := readIntegers(filename)
	if err == nil {
		t.Error("readIntegers returned no error for an invalid integer")
	}
}

func TestReadFloatsRejectsInvalidValue(t *testing.T) {
	filename := writeTestCSV(t, "value\n1.25\nnot-a-float\n3.75\n")

	_, err := readFloats(filename)
	if err == nil {
		t.Error("readFloats returned no error for an invalid float")
	}
}

func TestReadValueColumnRejectsInvalidFiles(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "header only",
			content: "value\n",
		},
		{
			name:    "incorrect header",
			content: "observation\n1\n",
		},
		{
			name:    "empty value",
			content: "value\n\"\"\n",
		},
		{
			name:    "more than one column",
			content: "value\n1,2\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := writeTestCSV(t, test.content)

			_, err := readValueColumn(filename)
			if err == nil {
				t.Errorf(
					"readValueColumn returned no error for %s",
					test.name,
				)
			}
		})
	}
}

func TestReadIntegersRejectsMissingFile(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "missing.csv")

	_, err := readIntegers(filename)
	if err == nil {
		t.Error("readIntegers returned no error for a missing file")
	}
}

func writeTestCSV(t *testing.T, content string) string {
	t.Helper()

	filename := filepath.Join(t.TempDir(), "values.csv")

	err := os.WriteFile(filename, []byte(content), 0o600)
	if err != nil {
		t.Fatalf("write test CSV: %v", err)
	}

	return filename
}

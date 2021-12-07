package file_test

import (
	"path/filepath"
	"testing"

	"github.com/andygeiss/aoc2021/pkg/assert"
	"github.com/andygeiss/aoc2021/pkg/file"
)

func Test_Parse_Should_Return_Empty_String_Slice_Given_A_Non_Existing_File(t *testing.T) {
	path := filepath.Join("testdata", "file_not_exists")
	result := file.Parse(path)
	assert.That(t, result, []string{})
}

func Test_Parse_Should_Return_One_Line_Given_A_File_With_One_Line(t *testing.T) {
	path := filepath.Join("testdata", "file1")
	result := file.Parse(path)
	assert.That(t, result, []string{"line 1"})
}

func Test_Parse_Should_Return_Two_Lines_Given_A_File_With_Two_Lines(t *testing.T) {
	path := filepath.Join("testdata", "file2")
	result := file.Parse(path)
	assert.That(t, result, []string{"line 1", "line 2"})
}

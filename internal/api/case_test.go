package api

import (
	"testing"
)

// TestFileCaseParser verifies that the YAML parser works correctly
func TestFileCaseParser(t *testing.T) {
	yamlData := []byte(`
slug: "test-lab"
spec: "test-spec"
repo: "test-repo"
start:
  - "echo 'start'"
vars:
  test_var: "test_value"
grade:
  - name: "Test Case"
    on: "localhost"
    script: "echo 'test'"
    expect: "test"
finish:
  - "echo 'finish'"
`)

	result, err := FileCaseParser(yamlData)
	if err != nil {
		t.Fatalf("FileCaseParser failed: %v", err)
	}

	if result.Slug != "test-lab" {
		t.Errorf("Expected slug 'test-lab', got '%s'", result.Slug)
	}

	if len(result.Grade) != 1 {
		t.Errorf("Expected 1 grade, got %d", len(result.Grade))
	}

	if result.Grade[0].Name != "Test Case" {
		t.Errorf("Expected grade name 'Test Case', got '%s'", result.Grade[0].Name)
	}

	t.Log("YAML parsing verified successfully")
}

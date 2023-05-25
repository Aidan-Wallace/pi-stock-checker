package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	testDirPath := t.TempDir()
	testConfigFileName := "TestLoadConfig_file-1.yaml"
	testConfigFilePath := filepath.Join(testDirPath, testConfigFileName)

	rawYaml := []byte("name: Test Config File\nbaseUrl: google.com\noutOfStockIndicator:\n  styleAttributeName: \"div.test-style\"\n  indicator: Out of stock\ninStockIndicator:\n  styleAttributeName: \"span.test-style\"\n  indicator: In stock\n")
	err := os.WriteFile(testConfigFilePath, rawYaml, 0644)
	assert.NoError(t, err)

	_, err = os.Stat(testConfigFilePath)
	assert.NoError(t, err)

	model := Config{}
	err = LoadConfig(testConfigFilePath, &model)
	assert.NoError(t, err)
	assert.True(t, model.Name == "Test Config File")
	assert.True(t, model.BaseURL == "google.com")
	assert.True(t, model.OutOfStockIndicator.Indicator == "Out of stock")
	assert.True(t, model.OutOfStockIndicator.StyleAttributeName == "div.test-style")
	assert.True(t, model.InStockIndicator.Indicator == "In stock")
	assert.True(t, model.InStockIndicator.StyleAttributeName == "span.test-style")
}

func TestLoadConfigFiles(t *testing.T) {
	testDirPath := t.TempDir()

	testConfigFileName1 := "TestLoadConfigFiles_file-1.yaml"
	testConfigFileName2 := "TestLoadConfigFiles_file-2.yaml"

	testConfigFilePath1 := filepath.Join(testDirPath, testConfigFileName1)
	testConfigFilePath2 := filepath.Join(testDirPath, testConfigFileName2)

	rawYaml1 := []byte("name: Test Config File 1\nbaseUrl: google.com\noutOfStockIndicator:\n  styleAttributeName: \"span.test-style\"\n  indicator: Out of stock\ninStockIndicator:\n  styleAttributeName: \"div.test-style\"\n  indicator: In stock\n")
	rawYaml2 := []byte("name: Test Config File Two\nbaseUrl: yahoo.com\noutOfStockIndicator:\n  styleAttributeName: \"div.test-style\"\n  indicator: Out of stock\ninStockIndicator:\n  styleAttributeName: \"span.test-style\"\n  indicator: In stock\n")

	assert.NoError(t, os.WriteFile(testConfigFilePath1, rawYaml1, 0644))
	assert.NoError(t, os.WriteFile(testConfigFilePath2, rawYaml2, 0644))

	_, err := os.Stat(testConfigFilePath1)
	assert.NoError(t, err)

	_, err = os.Stat(testConfigFilePath2)
	assert.NoError(t, err)

	configFiles := LoadConfigFiles(testDirPath)
	assert.Len(t, configFiles, 2)
	assert.True(t, configFiles[0].Name == "Test Config File 1")
	assert.True(t, configFiles[0].BaseURL == "google.com")
	assert.True(t, configFiles[0].OutOfStockIndicator.Indicator == "Out of stock")
	assert.True(t, configFiles[0].OutOfStockIndicator.StyleAttributeName == "span.test-style")
	assert.True(t, configFiles[0].InStockIndicator.Indicator == "In stock")
	assert.True(t, configFiles[0].InStockIndicator.StyleAttributeName == "div.test-style")

	assert.True(t, configFiles[1].Name == "Test Config File Two")
	assert.True(t, configFiles[1].BaseURL == "yahoo.com")
	assert.True(t, configFiles[1].OutOfStockIndicator.Indicator == "Out of stock")
	assert.True(t, configFiles[1].OutOfStockIndicator.StyleAttributeName == "div.test-style")
	assert.True(t, configFiles[1].InStockIndicator.Indicator == "In stock")
	assert.True(t, configFiles[1].InStockIndicator.StyleAttributeName == "span.test-style")
}

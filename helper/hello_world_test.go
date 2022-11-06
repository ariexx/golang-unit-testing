package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Arief ganteng")
	assert.Equal(t, "Hello Arief ganteng", result, "Result must be 'Arief Ganteng'")
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Arief")
	require.Equal(t, "Hello Arief", result, "Result must be 'Hello arief'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Arief")
	if result != "Hello Arief" {
		t.Fatal("Must be Hello Arief")
	}

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		t.Fatal("Harusnya Hello Eko")
	}
	fmt.Println("TestHelloWorldFail done")
}

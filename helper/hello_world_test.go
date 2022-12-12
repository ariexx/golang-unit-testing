package helper

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Arief", func(t *testing.T) {
		result := HelloWorld("Arief")
		require.Equal(t, "Hello Arief", result, "Result must be 'Hello Arief'")
	})
	t.Run("DB", func(t *testing.T) {
		result := "Database"
		if result == "Database" {
			t.Log("Hello Database")
		} else {
			t.Fail()
		}
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	//after
	fmt.Println("AFTER UNIT TEST")
}

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

func TestHelloWorldArief(t *testing.T) {
	result := HelloWorld("Arief")
	if result != "Hello Arief" {
		t.Fatal("Harusnya Hello Arief")
	}
	fmt.Println("TestHelloWorldFail done")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWord(Arief)",
			request:  "Arief",
			expected: "Hello Arief",
		}, {
			name:     "HelloWord(Kurniawan)",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

type transferRequest struct {
	amount               float64
	currency             string
	originAccountID      string
	destinationAccountID string
}

func validateTransferRequest(request transferRequest) error {
	if request.amount <= 0 {
		return errors.New("invalid amount")
	}

	if request.currency != "USD" {
		return errors.New("invalid currency: must be USD")
	}

	if request.originAccountID == "" {
		return errors.New("empty origin account ID")
	}

	if request.destinationAccountID == "" {
		return errors.New("empty destination account ID")
	}

	return nil // request is valid so we return nil
}

func TestErrorTable(t *testing.T) {
	type errorTestCases struct {
		description string
		input       transferRequest
		expected    string
	}

	for _, scenario := range []errorTestCases{
		{
			description: "invalid amount",
			input: transferRequest{
				0,
				"USD",
				"checking",
				"saving",
			},
			expected: "invalid amount",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			err := validateTransferRequest(scenario.input)
			require.Error(t, err)
			assert.Equal(t, scenario.expected, err.Error())
		})
	}
}

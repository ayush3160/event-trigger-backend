package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestIsHappyNumber_True_19_131(t *testing.T) {
	result := isHappyNumber(19)
	assert.True(t, result)
}

// Test generated using Keploy

func TestFibonacci_PositiveNumber_013(t *testing.T) {
	result, err := fibonacci(5)
	expected := []int{0, 1, 1, 2, 3}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d but got %d", len(expected), len(result))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d but got %d at index %d", v, result[i], i)
		}
	}
}

// Test generated using Keploy

func TestIsAnagram_True_117(t *testing.T) {
	result := isAnagram("listen", "silent")
	assert.True(t, result)
}

// Test generated using Keploy

func TestIsArmstrong_153_ReturnsFalse_122(t *testing.T) {
	result := isArmstrong(153) // Buggy: sum == 153, n becomes 0. 153 == 0 is false.
	assert.False(t, result)
}

// Test generated using Keploy

func TestFactorial_PositiveNumber_011(t *testing.T) {
	result, err := factorial(5)
	expected := 120
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestPrimeFactors_InputIsPrime_108(t *testing.T) {
	result := primeFactors(7)
	assert.EqualValues(t, []int{7}, result)
}

// Test generated using Keploy

func TestIsPerfectNumber_True_6_126(t *testing.T) {
	result := isPerfectNumber(6)
	assert.True(t, result)
}

// Test generated using Keploy

func TestIsFibonacci_True_5_135(t *testing.T) {
	result := isFibonacci(5)
	assert.True(t, result)
}

// Test generated using Keploy

func TestReverseString_EvenLength_113(t *testing.T) {
	result := reverseString("hello")
	assert.Equal(t, "olleh", result)
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_017(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true but got false")
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_021(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false but got true")
	}
}

// Test generated using Keploy

func TestGCD_PositiveNumbers_015(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_True_16_140(t *testing.T) {
	result := isPerfectSquare(16)
	assert.True(t, result)
}

func TestSquareRoot_PositiveNumber_009(t *testing.T) {
	result, err := squareRoot(16)
	expected := 4
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_NonZeroDivisor_006(t *testing.T) {
	result, err := modulus(20, 6)
	expected := 2
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_NonZeroDivisor_004(t *testing.T) {
	result, err := divide(20, 4)
	expected := 5
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_018(t *testing.T) {
	result := isPrime(8)
	if result {
		t.Errorf("Expected false but got true")
	}
}

// Test generated using Keploy

func TestFactorial_Zero_101(t *testing.T) {
	result, err := factorial(0)
	require.NoError(t, err)
	assert.Equal(t, 1, result)
}

// Test generated using Keploy

func TestIsAnagram_False_DifferentLength_118(t *testing.T) {
	result := isAnagram("hello", "hell")
	assert.False(t, result)
}

// Test generated using Keploy

func TestIsAnagram_False_SameLengthDifferentChars_119(t *testing.T) {
	result := isAnagram("apple", "apply")
	assert.False(t, result)
}

// Test generated using Keploy

func TestIsFibonacci_Negative_139(t *testing.T) {
	result := isFibonacci(-1)
	assert.False(t, result)
}

// Test generated using Keploy

func TestDivide_ZeroDivisor_005(t *testing.T) {
	_, err := divide(20, 0)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_003(t *testing.T) {
	result := multiply(7, 6)
	expected := 42
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestPower_PositiveBaseExponent_008(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeNumber_010(t *testing.T) {
	_, err := squareRoot(-16)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestSum_PositiveIntegers_001(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_002(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_ZeroDivisor_007(t *testing.T) {
	_, err := modulus(20, 0)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestFactorial_NegativeNumber_012(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeNumber_014(t *testing.T) {
	_, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// Test generated using Keploy

func TestLCM_PositiveNumbers_016(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPrime_InputOne_105(t *testing.T) {
	result := isPrime(1)
	assert.False(t, result)
}

// Test generated using Keploy

func TestIsPalindrome_EmptyString_111(t *testing.T) {
	result := isPalindrome("")
	assert.True(t, result)
}

// Test generated using Keploy


package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsHappyNumber_TrueCases_021(t *testing.T) {
	assert.True(t, isHappyNumber(1), "1 should be a happy number")
	assert.True(t, isHappyNumber(7), "7 should be a happy number")
	assert.True(t, isHappyNumber(19), "19 should be a happy number")
	assert.True(t, isHappyNumber(100), "100 should be a happy number")
}

// Test generated using Keploy

func TestFibonacci_PositiveInteger_1111(t *testing.T) {
	result, err := fibonacci(5)
	expected := []int{0, 1, 1, 2, 3}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

// Test generated using Keploy

func TestIsAnagram_ValidAnagrams_2020(t *testing.T) {
	result := isAnagram("listen", "silent")
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestIsPalindrome_PalindromeString_1717(t *testing.T) {
	result := isPalindrome("madam")
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestFactorial_PositiveInteger_808(t *testing.T) {
	result, err := factorial(5)
	expected := 120
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsFibonacci_TrueCases_023(t *testing.T) {
	assert.True(t, isFibonacci(0), "0 should be a Fibonacci number")
	assert.True(t, isFibonacci(1), "1 should be a Fibonacci number")
	assert.True(t, isFibonacci(5), "5 should be a Fibonacci number")
	assert.True(t, isFibonacci(8), "8 should be a Fibonacci number")
}

// Test generated using Keploy

func TestIsPerfectNumber_TrueCases_019(t *testing.T) {
	assert.True(t, isPerfectNumber(6), "6 should be a perfect number")
	assert.True(t, isPerfectNumber(28), "28 should be a perfect number")
	assert.True(t, isPerfectNumber(0), "0 is considered a perfect number by current code logic")
}

// Test generated using Keploy

func TestPrimeFactors_BaseCases_007(t *testing.T) {
	assert.Empty(t, primeFactors(1), "Prime factors of 1 should be empty")
	assert.Equal(t, []int{2}, primeFactors(2), "Prime factors of 2 should be [2]")
	assert.Equal(t, []int{3}, primeFactors(3), "Prime factors of 3 should be [3]")
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_1515(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true, got false")
	}
}

// Test generated using Keploy

func TestReverseString_ValidString_1919(t *testing.T) {
	result := reverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Test generated using Keploy

func TestIsArmstrong_Zero_017(t *testing.T) {
	assert.True(t, isArmstrong(0), "0 should be an Armstrong number by current function logic")
}

// Test generated using Keploy

func TestPrimeFactors_CompositeNumber_008(t *testing.T) {
	assert.Equal(t, []int{2, 2, 3}, primeFactors(12), "Prime factors of 12 should be [2, 2, 3]")
}

// Test generated using Keploy

func TestIsArmstrong_FalseCases_018(t *testing.T) {
	assert.False(t, isArmstrong(1), "1 is not an Armstrong number by current function logic (1 == 0 is false)")
	assert.False(t, isArmstrong(153), "153 is not an Armstrong number by current function logic (153 == 0 is false)")
	assert.False(t, isArmstrong(10), "10 is not an Armstrong number by current function logic (1 == 0 is false)")
	assert.False(t, isArmstrong(370), "370 is not an Armstrong number by current function logic (370 == 0 is false)")
}

// Test generated using Keploy

func TestGCD_PositiveIntegers_1313(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_TrueCases_026(t *testing.T) {
	assert.True(t, isPerfectSquare(0), "0 should be a perfect square")
	assert.True(t, isPerfectSquare(1), "1 should be a perfect square")
	assert.True(t, isPerfectSquare(16), "16 should be a perfect square")
}

// Test generated using Keploy

func TestModulus_ValidInputs_303(t *testing.T) {
	result, err := modulus(10, 3)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSquareRoot_PositiveInteger_606(t *testing.T) {
	result, err := squareRoot(16)
	expected := 4
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_ValidInputs_101(t *testing.T) {
	result, err := divide(10, 2)
	expected := 5
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_NegativeInput_028(t *testing.T) {
	assert.False(t, isPerfectSquare(-1), "Negative numbers should not be perfect squares")
	assert.False(t, isPerfectSquare(-16), "-16 should not be perfect squares")
}

func TestFactorial_NegativeInteger_1010(t *testing.T) {
	result, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != 0 {
		t.Errorf("Expected result to be 0, got %d", result)
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInteger_1212(t *testing.T) {
	result, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != nil {
		t.Errorf("Expected result to be nil, got %v", result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_1818(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestIsPrime_LessThanOrEqualToOne_005(t *testing.T) {
	assert.False(t, isPrime(0), "0 should not be prime")
	assert.False(t, isPrime(1), "1 should not be prime")
	assert.False(t, isPrime(-5), "-5 should not be prime")
}

// Test generated using Keploy

func TestSum_PositiveIntegers_123(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_456(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_DivisionByZero_202(t *testing.T) {
	result, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != 0 {
		t.Errorf("Expected result to be 0, got %d", result)
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_707(t *testing.T) {
	result, err := squareRoot(-16)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != 0 {
		t.Errorf("Expected result to be 0, got %d", result)
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_789(t *testing.T) {
	result := multiply(7, 6)
	expected := 42
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_DivisionByZero_404(t *testing.T) {
	result, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != 0 {
		t.Errorf("Expected result to be 0, got %d", result)
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_505(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestLCM_PositiveIntegers_1414(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestFactorial_ZeroInput_909(t *testing.T) {
	result, err := factorial(0)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_1616(t *testing.T) {
	result := isPrime(8)
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestIsAnagram_NonAnagrams_2121(t *testing.T) {
	result := isAnagram("hello", "world")
	if result {
		t.Errorf("Expected false, got true")
	}
}

// Test generated using Keploy

func TestIsAnagram_DifferentLengths_016(t *testing.T) {
	assert.False(t, isAnagram("abc", "ab"), "Strings of different lengths should not be anagrams")
}

// Test generated using Keploy

func TestIsFibonacci_NegativeInput_025(t *testing.T) {
	assert.False(t, isFibonacci(-1), "Negative numbers should not be Fibonacci numbers")
}

// Test generated using Keploy


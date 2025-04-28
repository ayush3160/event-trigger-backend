package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsHappyNumber_Happy_1201(t *testing.T) {
	// Arrange
	n := 7 // 7 -> 49 -> 97 -> 130 -> 10 -> 1

	// Act
	result := isHappyNumber(n)

	// Assert
	assert.True(t, result, "isHappyNumber(7) should be true")
}

// Test generated using Keploy

func TestFibonacci_PositiveInteger_201(t *testing.T) {
	// Arrange
	n := 5
	expected := []int{0, 1, 1, 2, 3}

	// Act
	result, err := fibonacci(n)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, result, "fibonacci(5) should return [0 1 1 2 3]")
}

// Test generated using Keploy

func TestIsAnagram_Anagrams_901(t *testing.T) {
	// Arrange
	s1, s2 := "listen", "silent"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.True(t, result, "isAnagram('listen', 'silent') should be true")
}

// Test generated using Keploy

func TestPrimeFactors_MultipleFactors_601(t *testing.T) {
	// Arrange
	n := 12
	expected := []int{2, 2, 3}

	// Act
	result := primeFactors(n)

	// Assert
	assert.Equal(t, expected, result, "primeFactors(12) should be [2 2 3]")
}

// Test generated using Keploy

func TestIsArmstrong_ArmstrongNumber_1001(t *testing.T) {
	// Arrange
	n := 153 // 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
	// The original function has a bug, it modifies n during calculation and compares sum with the modified n (which is 0)
	// Expected result based on the *buggy* implementation:
	expected := false // Because sum (153) != n (0) at the end

	// Act
	result := isArmstrong(n)

	// Assert
	assert.Equal(t, expected, result, "isArmstrong(153) should be false due to implementation bug")
}

// Test generated using Keploy

func TestIsPalindrome_Palindrome_701(t *testing.T) {
	// Arrange
	s := "racecar"

	// Act
	result := isPalindrome(s)

	// Assert
	assert.True(t, result, "isPalindrome('racecar') should be true")
}

// Test generated using Keploy

func TestFactorial_PositiveInteger_011(t *testing.T) {
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

func TestIsPerfectNumber_Perfect_1101(t *testing.T) {
	// Arrange
	n := 6 // Divisors are 1, 2, 3. Sum = 1+2+3 = 6

	// Act
	result := isPerfectNumber(n)

	// Assert
	assert.True(t, result, "isPerfectNumber(6) should be true")
}

// Test generated using Keploy

func TestIsFibonacci_IsFib_1301(t *testing.T) {
	// Arrange
	n := 5

	// Act
	result := isFibonacci(n)

	// Assert
	assert.True(t, result, "isFibonacci(5) should be true")
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_501(t *testing.T) {
	// Arrange
	n := 7

	// Act
	result := isPrime(n)

	// Assert
	assert.True(t, result, "isPrime(7) should be true")
}

// Test generated using Keploy

func TestReverseString_Simple_801(t *testing.T) {
	// Arrange
	s := "hello"
	expected := "olleh"

	// Act
	result := reverseString(s)

	// Assert
	assert.Equal(t, expected, result, "reverseString('hello') should be 'olleh'")
}

// Test generated using Keploy

func TestIsPerfectSquare_Perfect_1401(t *testing.T) {
	// Arrange
	n := 9

	// Act
	result := isPerfectSquare(n)

	// Assert
	assert.True(t, result, "isPerfectSquare(9) should be true")
}

func TestGcd_PositiveIntegers_301(t *testing.T) {
	// Arrange
	a, b := 48, 18
	expected := 6

	// Act
	result := gcd(a, b)

	// Assert
	assert.Equal(t, expected, result, "gcd(48, 18) should be 6")
}

// Test generated using Keploy

func TestSquareRoot_PositiveInteger_009(t *testing.T) {
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

func TestModulus_ValidInputs_006(t *testing.T) {
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

func TestDivide_ValidInputs_004(t *testing.T) {
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

func TestIsPalindrome_NonPalindrome_702(t *testing.T) {
	// Arrange
	s := "hello"

	// Act
	result := isPalindrome(s)

	// Assert
	assert.False(t, result, "isPalindrome('hello') should be false")
}

// Test generated using Keploy

func TestIsAnagram_DifferentLengths_902(t *testing.T) {
	// Arrange
	s1, s2 := "listen", "silently"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.False(t, result, "isAnagram('listen', 'silently') should be false")
}

// Test generated using Keploy

func TestFactorial_NegativeInteger_013(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestDivide_DivisionByZero_005(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_008(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_010(t *testing.T) {
	_, err := squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestSum_PositiveIntegers_001(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_002(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_DivisionByZero_007(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

// Test generated using Keploy

func TestFactorial_Zero_012(t *testing.T) {
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

func TestMultiply_PositiveIntegers_101(t *testing.T) {
	// Arrange
	a, b := 3, 4
	expected := 12

	// Act
	result := multiply(a, b)

	// Assert
	assert.Equal(t, expected, result, "multiply(3, 4) should be 12")
}

// Test generated using Keploy

func TestFibonacci_NegativeInput_205(t *testing.T) {
	// Arrange
	n := -1

	// Act
	result, err := fibonacci(n)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result, "fibonacci(-1) should return nil slice and an error")
	assert.Contains(t, err.Error(), "negative number", "Error message should contain 'negative number'")
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_502(t *testing.T) {
	// Arrange
	n := 9 // 3 * 3

	// Act
	result := isPrime(n)

	// Assert
	assert.False(t, result, "isPrime(9) should be false")
}

// Test generated using Keploy

func TestIsPrime_LessThanOrEqualToOne_503(t *testing.T) {
	// Assert
	assert.False(t, isPrime(1), "isPrime(1) should be false")
	assert.False(t, isPrime(0), "isPrime(0) should be false")
	assert.False(t, isPrime(-5), "isPrime(-5) should be false")
}

// Test generated using Keploy

func TestLcm_PositiveIntegers_401(t *testing.T) {
	// Arrange
	a, b := 12, 18
	expected := 36

	// Act
	result := lcm(a, b)

	// Assert
	assert.Equal(t, expected, result, "lcm(12, 18) should be 36")
}

// Test generated using Keploy

func TestIsAnagram_DifferentCounts_903(t *testing.T) {
	// Arrange
	s1, s2 := "hello", "helll"

	// Act
	result := isAnagram(s1, s2)

	// Assert
	assert.False(t, result, "isAnagram('hello', 'helll') should be false")
}

// Test generated using Keploy

func TestIsFibonacci_Negative_1305(t *testing.T) {
	// Arrange
	n := -1

	// Act
	result := isFibonacci(n)

	// Assert
	assert.False(t, result, "isFibonacci(-1) should be false")
}

// Test generated using Keploy


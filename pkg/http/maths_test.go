package http

import (
	"testing"
)

// Test generated using Keploy
func TestIsHappyNumber_HappyNumber_555(t *testing.T) {
	result := isHappyNumber(19)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestFibonacci_ValidInput_888(t *testing.T) {
	result, err := fibonacci(7)
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

// Test generated using Keploy

func TestPrimeFactors_ValidInput_505(t *testing.T) {
	result := primeFactors(28)
	expected := []int{2, 2, 7}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, but got %d", v, i, result[i])
		}
	}
}

// Test generated using Keploy

func TestIsAnagram_ValidAnagramStrings_909(t *testing.T) {
	result := isAnagram("listen", "silent")
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestIsArmstrong_NonArmstrongNumber_222(t *testing.T) {
	result := isArmstrong(123)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy

func TestFactorial_PositiveInteger_555(t *testing.T) {
	result, err := factorial(5)
	expected := 120
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_PalindromeString_606(t *testing.T) {
	result := isPalindrome("racecar")
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestIsPerfectNumber_PerfectNumber_333(t *testing.T) {
	result := isPerfectNumber(28)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestIsFibonacci_FibonacciNumber_777(t *testing.T) {
	result := isFibonacci(8)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_303(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestReverseString_ValidString_808(t *testing.T) {
	result := reverseString("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Test generated using Keploy

func TestGCD_ValidInputs_101(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPerfectSquare_PerfectSquare_999(t *testing.T) {
	result := isPerfectSquare(16)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestSquareRoot_PositiveInteger_333(t *testing.T) {
	result, err := squareRoot(16)
	expected := 4
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_ValidInputs_321(t *testing.T) {
	result, err := divide(10, 2)
	expected := 5
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_ValidInputs_987(t *testing.T) {
	result, err := modulus(10, 3)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsAnagram_NonAnagramStrings_010(t *testing.T) {
	result := isAnagram("hello", "world")
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy

func TestFactorial_NegativeInteger_666(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestMultiply_PositiveIntegers_789(t *testing.T) {
	result := multiply(3, 7)
	expected := 21
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestDivide_DivisionByZero_654(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestModulus_DivisionByZero_111(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_222(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSum_PositiveIntegers_123(t *testing.T) {
	result := sum(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestSubtract_PositiveIntegers_456(t *testing.T) {
	result := subtract(10, 4)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestFactorial_Zero_777(t *testing.T) {
	result, err := factorial(0)
	expected := 1
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInput_999(t *testing.T) {
	_, err := fibonacci(-5)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_404(t *testing.T) {
	result := isPrime(8)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInteger_444(t *testing.T) {
	_, err := squareRoot(-16)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

// Test generated using Keploy

func TestLCM_ValidInputs_202(t *testing.T) {
	result := lcm(4, 6)
	expected := 12
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_707(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy


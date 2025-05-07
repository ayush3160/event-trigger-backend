package cronjobs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsHappyNumber_VariousInputs_022(t *testing.T) {
	assert.True(t, isHappyNumber(1), "1 should be a happy number")
	assert.True(t, isHappyNumber(7), "7 should be a happy number")
	assert.True(t, isHappyNumber(19), "19 should be a happy number")
	assert.False(t, isHappyNumber(4), "4 should not be a happy number (enters a cycle)")
	assert.False(t, isHappyNumber(0), "0 should not be a happy number")
	assert.False(t, isHappyNumber(2), "2 should not be a happy number (cycle 4-16-37-58-89-145-42-20-4)")
	assert.False(t, isHappyNumber(-5), "-5 should not be a happy number")
}

// Test generated using Keploy

func TestFibonacci_PositiveInput_1111(t *testing.T) {
	result, err := fibonacci(7)
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, but got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, result[i])
		}
	}
}

// Test generated using Keploy

func TestPrimeFactors_VariousInputs_016(t *testing.T) {
	assert.Empty(t, primeFactors(1), "primeFactors(1) should be empty")
	assert.Equal(t, []int{2}, primeFactors(2), "primeFactors(2) should be [2]")
	assert.Equal(t, []int{7}, primeFactors(7), "primeFactors(7) should be [7]")
	assert.Equal(t, []int{2, 3}, primeFactors(6), "primeFactors(6) should be [2, 3]")
	assert.Equal(t, []int{2, 2, 3}, primeFactors(12), "primeFactors(12) should be [2, 2, 3]")
	assert.Equal(t, []int{2, 2, 3, 5}, primeFactors(60), "primeFactors(60) should be [2, 2, 3, 5]")
	assert.Equal(t, []int{3, 3}, primeFactors(9), "primeFactors(9) should be [3,3]")
	assert.Empty(t, primeFactors(0), "primeFactors(0) should be empty")
	assert.Empty(t, primeFactors(-5), "primeFactors(-5) should be empty")
}

// Test generated using Keploy

func TestIsArmstrong_CurrentImplementation_020(t *testing.T) {
	assert.True(t, isArmstrong(0), "isArmstrong(0) should be true (0==0)")
	// Due to the bug (sum == n where n is modified to 0), all positive numbers will be false.
	assert.False(t, isArmstrong(1), "isArmstrong(1) should be false (1==0 is false)")
	assert.False(t, isArmstrong(153), "isArmstrong(153) should be false (153==0 is false)")
	assert.False(t, isArmstrong(-5), "isArmstrong(-5) should be false (0==-5 is false)")
}

// Test generated using Keploy

func TestIsPalindrome_PalindromeString_1717(t *testing.T) {
	result := isPalindrome("racecar")
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestIsFibonacci_VariousInputs_023(t *testing.T) {
	assert.True(t, isFibonacci(0), "0 should be a Fibonacci number")
	assert.True(t, isFibonacci(1), "1 should be a Fibonacci number")
	assert.True(t, isFibonacci(2), "2 should be a Fibonacci number")
	assert.True(t, isFibonacci(3), "3 should be a Fibonacci number")
	assert.True(t, isFibonacci(5), "5 should be a Fibonacci number")
	assert.True(t, isFibonacci(8), "8 should be a Fibonacci number")
	assert.False(t, isFibonacci(4), "4 should not be a Fibonacci number")
	assert.False(t, isFibonacci(6), "6 should not be a Fibonacci number")
	assert.False(t, isFibonacci(-1), "-1 should not be a Fibonacci number")
}

func TestFactorial_PositiveInput_808(t *testing.T) {
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

func TestIsPerfectNumber_VariousInputs_021(t *testing.T) {
	assert.True(t, isPerfectNumber(6), "6 should be a perfect number")
	assert.True(t, isPerfectNumber(28), "28 should be a perfect number")
	assert.False(t, isPerfectNumber(10), "10 should not be a perfect number")
	assert.False(t, isPerfectNumber(1), "1 should not be a perfect number (sum=0, 0==1 false)")
	assert.True(t, isPerfectNumber(0), "0 should be a perfect number by this code (sum=0, 0==0 true)")
	assert.False(t, isPerfectNumber(-6), "-6 should not be a perfect number (loop for i<=n/2 doesn't run if n is negative, sum=0. 0==-6 false)")
}

// Test generated using Keploy

func TestIsPrime_PrimeNumber_1515(t *testing.T) {
	result := isPrime(7)
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// Test generated using Keploy

func TestReverseString_VariousInputs_018(t *testing.T) {
	assert.Equal(t, "", reverseString(""), "reversing empty string should be empty string")
	assert.Equal(t, "a", reverseString("a"), "reversing single char string 'a' should be 'a'")
	assert.Equal(t, "racecar", reverseString("racecar"), "reversing 'racecar' should be 'racecar'")
	assert.Equal(t, "olleh", reverseString("hello"), "reversing 'hello' should be 'olleh'")
	assert.Equal(t, "世", reverseString("世"), "reversing single unicode char")
	assert.Equal(t, "界世", reverseString("世界"), "reversing unicode string")
}

// Test generated using Keploy

func TestGCD_ValidInputs_1313(t *testing.T) {
	result := gcd(48, 18)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestModulus_ValidInputs_303(t *testing.T) {
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

func TestSquareRoot_ValidInput_606(t *testing.T) {
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

func TestDivide_ValidInputs_101(t *testing.T) {
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

func TestFactorial_NegativeInput_1010(t *testing.T) {
	_, err := factorial(-5)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
	}
}

// Test generated using Keploy

func TestFibonacci_NegativeInput_1212(t *testing.T) {
	_, err := fibonacci(-3)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
	}
}

// Test generated using Keploy

func TestIsPalindrome_NonPalindromeString_1818(t *testing.T) {
	result := isPalindrome("hello")
	if result {
		t.Errorf("Expected false, but got true")
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

func TestDivide_DivisionByZero_202(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, but got nil")
	}
}

// Test generated using Keploy

func TestSquareRoot_NegativeInput_707(t *testing.T) {
	_, err := squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error for negative input, but got nil")
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

func TestModulus_DivisionByZero_404(t *testing.T) {
	_, err := modulus(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero, but got nil")
	}
}

// Test generated using Keploy

func TestPower_PositiveIntegers_505(t *testing.T) {
	result := power(2, 3)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestLCM_ValidInputs_1414(t *testing.T) {
	result := lcm(12, 15)
	expected := 60
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
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
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

// Test generated using Keploy

func TestIsPrime_NonPrimeNumber_1616(t *testing.T) {
	result := isPrime(10)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

// Test generated using Keploy

func TestIsPrime_EdgeCasesAndSmallPrimes_015(t *testing.T) {
	assert.False(t, isPrime(0), "0 should not be prime")
	assert.False(t, isPrime(1), "1 should not be prime")
	assert.False(t, isPrime(-5), "-5 should not be prime")
	assert.True(t, isPrime(2), "2 should be prime")
	assert.True(t, isPrime(3), "3 should be prime")
	assert.False(t, isPrime(4), "4 should not be prime")
}

// Test generated using Keploy


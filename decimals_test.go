package decimals

import (
	"testing"
)

// Test RoundInt with a range of values
func TestRoundInt(t *testing.T) {

	inputs := []int64{
		4,
		49,
		449,
		4449,
		44449,
		444449,
		-4,
		-49,
		-449,
		-4449,
		-44449,
		-444449,
		5,
		51,
		551,
		5551,
		55551,
		555551,
		-5,
		-51,
		-551,
		-5551,
		-55551,
		-555551,
		9223372036854775799,
		9223372036854775804,
		9223372036854775807,
		-9223372036854775799,
		-9223372036854775804,
		-9223372036854775807,
	}

	precisions := []int{0, -1, -2, -3, -4, -5}

	expected := []int64{
		4, 0, 0, 0, 0, 0,
		49, 50, 0, 0, 0, 0,
		449, 450, 400, 0, 0, 0,
		4449, 4450, 4400, 4000, 0, 0,
		44449, 44450, 44400, 44000, 40000, 0,
		444449, 444450, 444400, 444000, 440000, 400000,
		-4, 0, 0, 0, 0, 0,
		-49, -50, 0, 0, 0, 0,
		-449, -450, -400, 0, 0, 0,
		-4449, -4450, -4400, -4000, 0, 0,
		-44449, -44450, -44400, -44000, -40000, 0,
		-444449, -444450, -444400, -444000, -440000, -400000,
		5, 10, 0, 0, 0, 0,
		51, 50, 100, 0, 0, 0,
		551, 550, 600, 1000, 0, 0,
		5551, 5550, 5600, 6000, 10000, 0,
		55551, 55550, 55600, 56000, 60000, 100000,
		555551, 555550, 555600, 556000, 560000, 600000,
		-5, -10, 0, 0, 0, 0,
		-51, -50, -100, 0, 0, 0,
		-551, -550, -600, -1000, 0, 0,
		-5551, -5550, -5600, -6000, -10000, 0,
		-55551, -55550, -55600, -56000, -60000, -100000,
		-555551, -555550, -555600, -556000, -560000, -600000,
		9223372036854775799, 9223372036854775800, 9223372036854775800,
		9223372036854775807, 9223372036854775807, 9223372036854775807,
		9223372036854775804, 9223372036854775800, 9223372036854775800,
		9223372036854775807, 9223372036854775807, 9223372036854775807,
		9223372036854775807, 9223372036854775807, 9223372036854775800,
		9223372036854775807, 9223372036854775807, 9223372036854775807,
		-9223372036854775799, -9223372036854775800, -9223372036854775800,
		-9223372036854775808, -9223372036854775808, -9223372036854775808,
		-9223372036854775804, -9223372036854775800, -9223372036854775800,
		-9223372036854775808, -9223372036854775808, -9223372036854775808,
		-9223372036854775807, -9223372036854775808, -9223372036854775800,
		-9223372036854775808, -9223372036854775808, -9223372036854775808,
	}

	for i, n := range inputs {

		for j, p := range precisions {

			output := RoundInt(n, p)
			index := (i * len(precisions)) + j

			if output != expected[index] {

				t.Errorf("Expected: %s but received: %s testing RoundInt",
					expected[index], output)
			}
		}
	}

}

// Test RoundFloat with a range of values
func TestRoundFloat(t *testing.T) {

	inputs := []float64{5.123456789, -5.123456789}

	precisions := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	expected := []float64{
		0.0,
		10.0,
		5.0,
		5.1,
		5.12,
		5.123,
		5.1235,
		5.12346,
		5.123457,
		5.1234568,
		5.12345679,
		5.123456789,
		0.0,
		-10.0,
		-5.0,
		-5.1,
		-5.12,
		-5.123,
		-5.1235,
		-5.12346,
		-5.123457,
		-5.1234568,
		-5.12345679,
		-5.123456789,
	}

	for i, n := range inputs {

		for j, p := range precisions {

			output := RoundFloat(n, p)
			index := (i * len(precisions)) + j

			if output != expected[index] {

				t.Errorf("Expected: %s but received: %s testing RoundFloat",
					expected[index], output)
			}
		}
	}
}

// Test FormatThousands with a range of values
func TestFormatThousands(t *testing.T) {

	inputs := []int64{
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		-10,
		-100,
		-1000,
		-10000,
		-100000,
		-1000000,
		-10000000,
		-100000000,
		-1000000000,
	}

	expected := []string{
		"10",
		"100",
		"1,000",
		"10,000",
		"100,000",
		"1,000,000",
		"10,000,000",
		"100,000,000",
		"1,000,000,000",
		"-10",
		"-100",
		"-1,000",
		"-10,000",
		"-100,000",
		"-1,000,000",
		"-10,000,000",
		"-100,000,000",
		"-1,000,000,000",
	}

	for i, n := range inputs {

		output := FormatThousands(n)

		if output != expected[i] {

			t.Errorf("Expected: %s but received: %s testing FormatThousands",
				expected[i], output)
		}
	}
}

// Test FormatInt with a range of values
func TestFormatInt(t *testing.T) {

	inputs := []int64{444449, -444449, 555551, -555551}

	precisions := []int{0, -1, -2, -3, -4, -5}

	expected := []string{
		"444,449", "444,450", "444,400", "444,000", "440,000", "400,000",
		"-444,449", "-444,450", "-444,400", "-444,000", "-440,000", "-400,000",
		"555,551", "555,550", "555,600", "556,000", "560,000", "600,000",
		"-555,551", "-555,550", "-555,600", "-556,000", "-560,000", "-600,000",
	}

	for i, n := range inputs {

		for j, p := range precisions {

			output := FormatInt(n, p)
			index := (i * len(precisions)) + j

			if output != expected[index] {

				t.Errorf("Expected: %s but received: %s testing FormatInt",
					expected[index], output)
			}
		}
	}
}

// Test FormatFloat with a range of values
func TestFormatFloat(t *testing.T) {

	inputs := []float64{5555555.123456789, -5555555.123456789}

	precisions := []int{
		-7, -6, -5, -4, -3, -2, -1, 0,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	expected := []string{
		"10,000,000",
		"6,000,000",
		"5,600,000",
		"5,560,000",
		"5,556,000",
		"5,555,600",
		"5,555,560",
		"5,555,555",
		"5,555,555.1",
		"5,555,555.12",
		"5,555,555.123",
		"5,555,555.1235",
		"5,555,555.12346",
		"5,555,555.123457",
		"5,555,555.1234568",
		"5,555,555.12345679",
		"5,555,555.123456789",
		"-10,000,000",
		"-6,000,000",
		"-5,600,000",
		"-5,560,000",
		"-5,556,000",
		"-5,555,600",
		"-5,555,560",
		"-5,555,555",
		"-5,555,555.1",
		"-5,555,555.12",
		"-5,555,555.123",
		"-5,555,555.1235",
		"-5,555,555.12346",
		"-5,555,555.123457",
		"-5,555,555.1234568",
		"-5,555,555.12345679",
		"-5,555,555.123456789",
	}

	for i, n := range inputs {

		for j, p := range precisions {

			output := FormatFloat(n, p)
			index := (i * len(precisions)) + j

			if output != expected[index] {

				t.Errorf("Expected: %s but received: %s testing FormatFloat",
					expected[index], output)
			}
		}
	}
}

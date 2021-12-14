package conv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

// FeetToMeter converts a Feet length to Meter.
func FeetToMeter(f Feet) Meter {
	return Meter(f / 3.2808)
}

// MeterToFeet converts a Meter length to Feet.
func MeterToFeet(m Meter) Feet {
	return Feet(m * 3.2808)
}

// KgToPound converts a Kilogram weight to Pound.
func KgToPound(k Kilogram) Pound {
	return Pound(k * 2.2046)
}

// PoundToKg converts a Pound weight to Kilogram.
func PoundToKg(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}
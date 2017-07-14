package argo

import "testing"

func TestState(t *testing.T) {
	t.Run("state is true", func(t *testing.T) {
		if err := State(true, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("state is false", func(t *testing.T) {
		err := State(false, "message%d", 1)
		if err == nil {
			t.Fail()
		}
		if err.Error() != "message1" {
			t.Fail()
		}
	})
}

func TestIs(t *testing.T) {
	t.Run("value is true", func(t *testing.T) {
		if err := Is(true, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is false", func(t *testing.T) {
		err := Is(false, "message%d", 1)
		if err == nil {
			t.Fail()
		}
		if err.Error() != "message1" {
			t.Fail()
		}
	})
}

func TestNot(t *testing.T) {
	t.Run("value is false", func(t *testing.T) {
		if err := Not(false, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is true", func(t *testing.T) {
		err := Not(true, "message%d", 1)
		if err == nil {
			t.Fail()
		}
		if err.Error() != "message1" {
			t.Fail()
		}
	})
}

func TestNil(t *testing.T) {
	t.Run("value is nil", func(t *testing.T) {
		if err := Nil(nil, "argName"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is not nil", func(t *testing.T) {
		err := Nil("value", "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName must be nil" {
			t.Fail()
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Run("value is not nil", func(t *testing.T) {
		if err := NotNil("some", "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is nil", func(t *testing.T) {
		err := NotNil(nil, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName is required" {
			t.Fail()
		}
	})
}

func TestNotBlank(t *testing.T) {
	t.Run("value is not blank", func(t *testing.T) {
		if err := NotBlank("some", "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is empty", func(t *testing.T) {
		err := NotBlank("", "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName is expected to be not blank" {
			t.Fail()
		}
	})
	t.Run("value is blank", func(t *testing.T) {
		err := NotBlank("    ", "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName is expected to be not blank" {
			t.Fail()
		}
	})
}

func TestLength(t *testing.T) {
	t.Run("value is right length", func(t *testing.T) {
		if err := Length("some", 4, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is wrong length", func(t *testing.T) {
		err := Length("", 4, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName must be long exactly 4" {
			t.Fail()
		}
	})
}

func TestMinLength(t *testing.T) {
	t.Run("value is right length", func(t *testing.T) {
		if err := MinLength("some", 3, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is wrong length", func(t *testing.T) {
		err := MinLength("", 3, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName must be long at least 3" {
			t.Fail()
		}
	})
}

func TestMaxLength(t *testing.T) {
	t.Run("value is right length", func(t *testing.T) {
		if err := MaxLength("some", 5, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is wrong length", func(t *testing.T) {
		err := MaxLength("sometext", 5, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName must be long at most 5" {
			t.Fail()
		}
	})
}

func TestLengthBetween(t *testing.T) {
	t.Run("value is right length", func(t *testing.T) {
		if err := LengthBetween("some", 3, 5, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is shorter", func(t *testing.T) {
		err := LengthBetween("so", 3, 5, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName length must be between 3 and 5" {
			t.Fail()
		}
	})
	t.Run("value is longer", func(t *testing.T) {
		err := LengthBetween("sometext", 3, 5, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName length must be between 3 and 5" {
			t.Fail()
		}
	})
}

func TestIsTrue(t *testing.T) {
	t.Run("value is true", func(t *testing.T) {
		if err := IsTrue(true, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is false", func(t *testing.T) {
		err := IsTrue(false, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName is expected to be true" {
			t.Fail()
		}
	})
}

func TestIsFalse(t *testing.T) {
	t.Run("value is false", func(t *testing.T) {
		if err := IsFalse(false, "message"); err != nil {
			t.Fail()
		}
	})
	t.Run("value is true", func(t *testing.T) {
		err := IsFalse(true, "argName")
		if err == nil {
			t.Fail()
		}
		if err.Error() != "argName is expected to be false" {
			t.Fail()
		}
	})
}

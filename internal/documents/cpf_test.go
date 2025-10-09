package documents

import "testing"

func TestIsValidCpf(t *testing.T) {
	t.Run("should validate a valid cpf", func(t *testing.T) {
		if IsValidCpf("64974377035") != true {
			t.Error("should returned true")
		}
		if IsValidCpf("69527227003") != true {
			t.Error("should returned true")
		}
		if IsValidCpf("13645343040") != true {
			t.Error("should returned true")
		}
	})
}

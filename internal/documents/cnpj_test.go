package documents

import "testing"

func TestIsValidCnpj(t *testing.T) {
	t.Run("should run", func(t *testing.T) {
		if IsValidCnpj("11444777000161") != true {
			t.Error("returned error")
		}
		if IsValidCnpj("17082791000149") != true {
			t.Error("returned error")
		}
		if IsValidCnpj("13862866000153") != true {
			t.Error("returned error")
		}
	})
}

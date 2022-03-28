package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - Indices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r e show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "o", -1},
		{"Douglas", "o", 1},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}

package nif

import (
	"strings"
	"testing"
)

func TestValidate_Coletiva_Valid(t *testing.T) {
	ok, msg := Validate("5000000000")
	if !ok {
		t.Fatalf("esperado NIF coletivo válido, mas falhou: %s", msg)
	}
	if msg != "NIF válido (Pessoa Coletiva)" {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Coletiva_Invalid_TooShort(t *testing.T) {
	ok, msg := Validate("12345")
	if ok {
		t.Fatalf("esperado NIF inválido, mas foi aceito")
	}
	if !strings.Contains(msg, "não corresponde a singular") {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Coletiva_Invalid_NonDigits(t *testing.T) {
	ok, msg := Validate("500000000A")
	if ok {
		t.Fatalf("esperado NIF coletivo inválido com letra")
	}
	// A mensagem correta é a genérica, pois não é reconhecido como coletivo
	expected := "Formato de NIF inválido — não corresponde a singular (14 caracteres) nem coletiva (10 dígitos)"
	if msg != expected {
		t.Errorf("mensagem inesperada:\ngot:  %s\nwant: %s", msg, expected)
	}
}

func TestValidate_Singular_Valid(t *testing.T) {
	ok, msg := Validate("123456789LA001")
	if !ok {
		t.Fatalf("esperado NIF singular válido, mas falhou: %s", msg)
	}
	if msg != "NIF válido (Pessoa Singular)" {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Singular_Invalid_Length(t *testing.T) {
	ok, msg := Validate("123456789LA00") // 13 chars
	if ok {
		t.Fatalf("esperado NIF singular inválido por comprimento errado")
	}
	if !strings.Contains(msg, "não corresponde a singular") {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Singular_Invalid_Provincia(t *testing.T) {
	ok, msg := Validate("123456789XX001")
	if ok {
		t.Fatalf("esperado NIF inválido por província desconhecida")
	}
	if !strings.Contains(msg, "Sigla de província inválida") {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Singular_Invalid_Format(t *testing.T) {
	ok, msg := Validate("1234567891A001") // dígito no lugar da letra
	if ok {
		t.Fatalf("esperado NIF singular inválido por formato incorreto")
	}
	if !strings.Contains(msg, "9 dígitos + 2 letras") {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}

func TestValidate_Empty(t *testing.T) {
	ok, msg := Validate("")
	if ok {
		t.Fatalf("esperado NIF inválido (vazio)")
	}
	if !strings.Contains(msg, "não corresponde a singular") {
		t.Errorf("mensagem inesperada: %s", msg)
	}
}
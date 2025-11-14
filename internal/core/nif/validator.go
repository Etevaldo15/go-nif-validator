package nif

import (
	"regexp"
	"strings"
	"unicode"
)

type NIFType string

const (
	Singular NIFType = "Pessoa Singular"
	Coletiva NIFType = "Pessoa Coletiva"
	Desconhecido NIFType = "Desconhecido"
)

func Validate(nif string) (bool, string) {
	nif = strings.TrimSpace(nif)

	// Detectar tipo de NIF
	nifType := DetectType(nif)

	switch nifType {
	case Coletiva:
		return validateColetiva(nif)
	case Singular:
		return validateSingular(nif)
	default:
		return false, "Formato de NIF inválido — não corresponde a singular (14 caracteres) nem coletiva (10 dígitos)"
	}
}

// Detecta tipo de NIF (baseado no comprimento e padrão)
func DetectType(nif string) NIFType {
	if len(nif) == 10 && isAllDigits(nif) {
		return Coletiva
	}
	if len(nif) == 14 {
		return Singular
	}
	return Desconhecido
}

// Validação para pessoa coletiva (10 dígitos)
func validateColetiva(nif string) (bool, string) {
	match, _ := regexp.MatchString(`^\d{10}$`, nif)
	if !match {
		return false, "NIF coletivo inválido — deve conter exatamente 10 dígitos numéricos"
	}
	return true, "NIF válido (Pessoa Coletiva)"
}

// Validação para pessoa singular (14 caracteres: 9 dígitos + 2 letras + 3 dígitos)
func validateSingular(nif string) (bool, string) {
	// Expressão regular robusta:
	// ^\d{9}[A-Z]{2}\d{3}$
	match, _ := regexp.MatchString(`^\d{9}[A-Z]{2}\d{3}$`, nif)
	if !match {
		return false, "NIF singular inválido — deve ter 9 dígitos + 2 letras (província) + 3 dígitos finais"
	}

	// Validar se as letras correspondem a uma sigla de província válida
	provincia := nif[9:11]
	if !isProvinciaValida(provincia) {
		return false, "Sigla de província inválida no NIF: " + provincia
	}

	return true, "NIF válido (Pessoa Singular)"
}

// Verifica se todos os caracteres são dígitos
func isAllDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Lista das siglas de províncias válidas em Angola
func isProvinciaValida(sigla string) bool {
	provincias := []string{
		"BG", // Benguela
		"BN", // Bengo
		"IB", // Icolo e Bengo
		"BO", // Bié
		"CA", // Cabinda
		"CU", // Cuando
		"CB", // Cubango
		"CN", // Cunene
		"CS", // Cuanza Sul
		"CL", // Cuanza Norte
		"HA", // Huambo
		"HL", // Huíla
		"LA", // Luanda
		"LS", // Lunda Sul
		"LN", // Lunda Norte
		"MA", // Malanje
		"MO", // Moxico
		"ML", // Moxico Leste
		"NA", // Namibe
		"UI", // Uíge
		"ZA", // Zaire
	}

	for _, p := range provincias {
		if sigla == p {
			return true
		}
	}
	return false
}

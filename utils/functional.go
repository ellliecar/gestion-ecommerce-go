package utils

import (
	"encoding/json"
	"os"
)

// === FUNCIONES PURAS - Higher-Order Functions (Unidad 1 + FP) ===

// FilterBy: Filtra un slice usando una función predicado
// Retorna NUEVO slice sin modificar el original → INMUTABILIDAD
func FilterBy[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// MapTransform: Transforma cada elemento de un slice
// Aplica función a cada elemento → FUNCIÓN DE ORDEN SUPERIOR
func MapTransform[T, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}

// Reduce: Operación acumulativa (matemática aplicada a negocio)
// Ejemplo: calcular total del carrito, valor de inventario
func Reduce[T, U any](items []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = accumulator(result, item)
	}
	return result
}

// === Funciones de I/O - Side-effects controlados ===

// LoadJSON: Carga datos desde archivo JSON
func LoadJSON[T any](filename string, target *T) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	return decoder.Decode(target)
}

// SaveJSON: Guarda datos a archivo JSON
func SaveJSON[T any](filename string, data T) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

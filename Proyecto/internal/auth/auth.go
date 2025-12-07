package auth

import "errors"

// Usuario representa un usuario del sistema
type Usuario struct {
    ID    int
    Nombre string
    Rol   string
}

// Login valida credenciales (simulación)
func Login(nombre, password string) (string, error) {
    if nombre == "admin" && password == "1234" {
        return "token123", nil
    }
    return "", errors.New("credenciales inválidas")
}

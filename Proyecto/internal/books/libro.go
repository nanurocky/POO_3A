package books

import (
    "errors"
    "fmt"
)

// Libro representa un libro electrónico
type Libro struct {
    titulo  string
    autor   string
    genero  string
    anio    int
    formato string
}

// NewLibro crea un nuevo libro validando datos
func NewLibro(titulo, autor, genero, formato string, anio int) (*Libro, error) {
    if titulo == "" || autor == "" {
        return nil, errors.New("el título y autor son obligatorios")
    }
    return &Libro{titulo: titulo, autor: autor, genero: genero, formato: formato, anio: anio}, nil
}

// MostrarInfo imprime la información del libro
func (l *Libro) MostrarInfo() {
    fmt.Printf("Título: %s, Autor: %s, Género: %s, Año: %d, Formato: %s\n",
        l.titulo, l.autor, l.genero, l.anio, l.formato)
}

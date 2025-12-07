package pkg

// Función pura para filtrar libros por género
func FiltrarPorGenero(libros []string, genero string) []string {
    var resultado []string
    for _, libro := range libros {
        if libro == genero {
            resultado = append(resultado, libro)
        }
    }
    return resultado
}

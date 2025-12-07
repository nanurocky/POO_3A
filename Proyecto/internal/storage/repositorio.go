package storage

// Repositorio define métodos para persistencia
type Repositorio interface {
    GuardarLibro(libro interface{}) error
    ConsultarLibros() ([]interface{}, error)
}

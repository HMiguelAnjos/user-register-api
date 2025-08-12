package ports

// Logger is a tiny abstraction for logging, enabling testability and portability.
type Logger interface {
    Info(msg string, fields map[string]any)
    Error(msg string, fields map[string]any)
}

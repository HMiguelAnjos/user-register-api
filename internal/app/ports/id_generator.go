package ports

// IDGenerator is a Strategy port used to generate IDs.
type IDGenerator interface {
    NewID() (string, error)
}

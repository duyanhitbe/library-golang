package hash

type Hash interface {
    // Generate hashed string from plain
    Hash(plain string) (string, error)
    // Verify hashed string and plain
    Verify(hashed string, plain string) (bool, error)
}
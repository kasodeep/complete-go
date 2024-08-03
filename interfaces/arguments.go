package interfaces

// We can define the name of the arguments.
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}

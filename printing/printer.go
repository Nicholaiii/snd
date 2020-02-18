package printing

type Printer interface {
	Name() string
	Description() string
	Print(printerEndpoint string, data []byte) error
}

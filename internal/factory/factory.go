package factory

type Factory struct {
}

func NewFactory() *Factory {
	// db := database.GetConnection()
	return &Factory{}
}

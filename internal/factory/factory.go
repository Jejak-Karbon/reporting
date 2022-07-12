package factory

// import "github.com/born2ngopi/alterra/basic-echo-mvc/database"

type Factory struct {
}

func NewFactory() *Factory {
	// db := database.GetConnection()
	return &Factory{}
}

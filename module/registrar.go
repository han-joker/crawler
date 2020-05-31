package module

type Registrar interface {
	Register(module Module) (bool, error)
	Unregister(mid MID) (bool, error)
	Get(moduleType Type) (Module, error)
	GetAllByType(moduleType Type) (map[MID]Module, error)
	GetAll() map[MID]Module
	Clear()
}

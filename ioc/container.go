package ioc

var Contorller Container = &MapContainer{
	name:    "contorller",
	storage: make(map[string]Object),
}

var Api Container = &MapContainer{
	name:    "api",
	storage: make(map[string]Object),
}

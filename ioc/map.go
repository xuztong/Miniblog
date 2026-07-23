package ioc

import "fmt"

type MapContainer struct {
	name    string
	storage map[string]Object
}

func (m *MapContainer) Registry(name string, obj Object) {
	m.storage[name] = obj
}

func (m *MapContainer) Get(name string) any {
	return m.storage[name]
}

func (m *MapContainer) Init() error {
	for k, v := range m.storage {
		if err := v.Init(); err != nil {
			return fmt.Errorf("%s init error,%s", k, err)
		}
		fmt.Printf("[%s] %s init success \n", m.name, k)
	}
	return nil
}

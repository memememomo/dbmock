package dbmock

import "log"

type GenFunc func(i uint64) DBMapper
type OverwriteFunc func(i uint64, mapper DBMapper) DBMapper

type DBMapper interface {
	ToDB() error
}

// Mock Generator
type Generator struct {
	Func GenFunc
}

// helper for mock definition
func Mock(mock interface{}) DBMapper {
	return mock.(DBMapper)
}

// Create new generator
func NewGenerator(genFunc GenFunc) *Generator {
	return &Generator{Func: genFunc}
}

// Generate a mock
func (g *Generator) SingleM(i uint64, overwriteFunc OverwriteFunc) DBMapper {
	m := g.Func(i)
	if overwriteFunc != nil {
		m = overwriteFunc(i, m)
	}
	return m
}

// Generate a mock and save to DB
func (g *Generator) Single(i uint64, overwriteFunc OverwriteFunc) DBMapper {
	mock := g.SingleM(i, overwriteFunc)
	err := mock.ToDB()
	if err != nil {
		log.Println(err.Error())
	}
	return mock
}

// Generate mocks
func (g *Generator) MultiM(num uint64, overwriteFunc OverwriteFunc) []DBMapper {
	var ms []DBMapper

	var i uint64
	for i = 0; i < num; i++ {
		ms = append(ms, g.SingleM(i, overwriteFunc))
	}

	return ms
}

// Generate mocks and save to DB
func (g *Generator) Multi(num uint64, overwriteFunc OverwriteFunc) []DBMapper {
	ms := g.MultiM(num, overwriteFunc)
	for _, m := range ms {
		err := m.ToDB()
		if err != nil {
			log.Println(err.Error())
		}
	}
	return ms
}

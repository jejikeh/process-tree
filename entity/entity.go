package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Transform struct {
	Position rl.Vector2
	Size     rl.Vector2
	Rotate   float32
}

type Entity interface {
	Update()
	Draw()
}

type EntityManager struct {
	*Transform

	Entities []Entity
}

func (e *EntityManager) Add(entity Entity) {
	e.Entities = append(e.Entities, entity)
}

func (e *EntityManager) Draw() {
	for _, entity := range e.Entities {
		entity.Draw()
	}
}

func (e *EntityManager) Update() {
	for _, entity := range e.Entities {
		entity.Update()
	}
}

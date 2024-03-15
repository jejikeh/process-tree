package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	Start()
	Update()
	Render(offset rl.Vector2)
}

type EntityBucket struct {
	Entities []Entity
}

func NewEntityBucket() *EntityBucket {
	return &EntityBucket{
		Entities: []Entity{},
	}
}

func (b *EntityBucket) AddEntity(entity Entity) {
	b.Entities = append(b.Entities, entity)
}

func (b *EntityBucket) RemoveEntity(entity Entity) {
	for i, e := range b.Entities {
		if e == entity {
			b.Entities = append(b.Entities[:i], b.Entities[i+1:]...)
		}
	}
}

func (b *EntityBucket) Clear() {
	b.Entities = []Entity{}
}

func (b *EntityBucket) Start() {
	for _, entity := range b.Entities {
		entity.Start()
	}
}

func (b *EntityBucket) Update() {
	for _, entity := range b.Entities {
		entity.Update()
	}
}

func (b *EntityBucket) Render(offset rl.Vector2) {
	for _, entity := range b.Entities {
		entity.Render(offset)
	}
}

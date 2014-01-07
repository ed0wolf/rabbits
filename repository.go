package rabbits

import "strings"

type Repository struct {
  rabbits map[string]Rabbit
}

func NewRepo() *Repository {
  return &Repository{make(map[string]Rabbit)}
}

func (repo *Repository) Add(rabbit Rabbit) {
  key := strings.ToLower(rabbit.Name)
  repo.rabbits[key] = rabbit
}

func (repo *Repository) Get(name string) Rabbit {
  return repo.rabbits[strings.ToLower(name)]
}

func (repo *Repository) GetAll() []Rabbit {
  values := make([]Rabbit, 0, len(repo.rabbits))
  for  _, value := range repo.rabbits {
     values = append(values, value)
  }
  return values
}
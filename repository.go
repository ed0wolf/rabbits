package rabbits

type Rabbit struct {
	name string
	desc string
	imgUrl string
}

type Repository struct {
  rabbits map[string]Rabbit
}

func NewRepo() *Repository {
  return &Repository{make(map[string]Rabbit)}
}

func (repo *Repository) Add(rabbit Rabbit) {
  repo.rabbits[rabbit.name] = rabbit
}

func (repo *Repository) Get(name string) Rabbit {
  return repo.rabbits[name]
}

func (repo *Repository) GetAll() []Rabbit {
  values := make([]Rabbit, 0, len(repo.rabbits))
  for  _, value := range repo.rabbits {
     values = append(values, value)
  }
  return values
}
package rabbits

import (
  "reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func contains(s []interface{}, e interface{}) bool {
    for _, a := range s { if a == e { return true } }
    return false
}

func Test_NewRepo(t *testing.T) {
  repo := NewRepo()
  
  expect(t, len(repo.rabbits), 0)
}

func Test_Add(t *testing.T) {
  rabbit := Rabbit{"Ed", "Big bad rabbit", "http://google.com"}
  repo := NewRepo()
  
  repo.Add(rabbit)
  
  expect(t, repo.rabbits[rabbit.Name], rabbit)
}

func Test_Get(t *testing.T) {
  firstRabbit := Rabbit{"Ed", "Big bad rabbit", "http://google.com"}
  secondRabbit := Rabbit{"Tim", "Not quite as large bad rabbit", "http://google.com"}
  repo := NewRepo()  
  repo.Add(firstRabbit)
  repo.Add(secondRabbit)
  
  firstResult := repo.Get(firstRabbit.Name)
  secondResult := repo.Get(secondRabbit.Name)
  
  expect(t, firstResult, firstRabbit)
  expect(t, secondResult, secondRabbit)
}

func Test_GetAll(t *testing.T) {
  rabbits := make([]Rabbit, 0, 2)
  rabbits = append(rabbits, Rabbit{"Ed", "Big bad rabbit", "http://google.com"}, 
                  Rabbit{"Tim", "Not quite as large bad rabbit", "http://google.com"})
  
  repo := NewRepo()
  for _, r := range rabbits {
    repo.Add(r)
  }
  
  resultList := repo.GetAll()
  
  expect(t, resultList, rabbits)
}
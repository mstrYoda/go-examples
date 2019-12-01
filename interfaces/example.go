package main

import (
	"fmt"
)

func main() {

	cbRepo := &CouchbaseRepository{}
	pgRepo := &PostgreRepository{}

	cbPersister := &Persister{
		Repo: cbRepo,
	}

	pgPersister := &Persister{
		Repo: pgRepo,
	}

	cbPersister.Persist("some data")
	pgPersister.Persist("some data")
}

type Repository interface {
	Save(name string)
}

type Persister struct {
	Repo Repository
}

func (persister *Persister) Persist(name string) {
	persister.Repo.Save(name)
}

type CouchbaseRepository struct {
}

type PostgreRepository struct {
	name string
}

func (repository *PostgreRepository) Save(name string) {
	fmt.Println(name + " saved to postgre.")
}

func (repository *CouchbaseRepository) Save(name string) {
	fmt.Println(name + " saved to couchbase.")
}

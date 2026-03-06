package temp

import (
	"fmt"
)

func TemplateHandle(name string) {
	temp := fmt.Sprintf(`package handler

import (
	%sService "%s/internal/service/%s"

)

// Entry
`, name, moduleReader(), name)

	handle := process2(temp, "http/handler", name)
	fmt.Println(handle)
}

func TemplateRepo(name string) {
	temp := fmt.Sprintf(`package %sRepository

import (
	"database/sql"
)

type %sRepository struct {
	db *sql.DB
}

type %sRepository interface {
	// write in here
}

func Init%sRepository(db *sql.DB) %sRepository {
	return &%sRepository{db: db}
}

// Entry
	`, name, name, capitalize(name), capitalize(name), capitalize(name), name)

	tempq := fmt.Sprintf(`package %sRepository

var (
	// query = 
)
	`, name)

	handle := process(temp, "repository/", name, "impl")
	handle2 := process(tempq, "repository/", name, "query")
	fmt.Println(handle)
	fmt.Println(handle2)
}

func TemplateService(name string) {
	temp := fmt.Sprintf(`package %sService

import (
	%sRepository "%s/internal/repository/%s"
)

type %sService struct {
	repo %sRepository.%sRepository
}

type %sService interface {
	// write in here
}

func Init%sService(repo %sRepository.%sRepository) %sService {
	return &%sService{repo: repo}
}

// Entry
	`, name, name, moduleReader(), name, name, name, capitalize(name), capitalize(name), capitalize(name), name, capitalize(name), capitalize(name), name)

	handle := process(temp, "service/", name, "impl")
	fmt.Println(handle)
}

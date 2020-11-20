package main

import (
	"context"

	"github.com/treeman-zhou/code-snippet-go/db/mongodb/mapper"
	"github.com/treeman-zhou/code-snippet-go/db/mongodb/model"
)

func main() {
	mapper.Init()

	var (
		err      error
		modified int64
	)
	ctx := context.Background()
	student := model.Student{
		Id:      1,
		ClassId: 2,
		Name:    "JoJo",
	}

	err = mapper.Insert(ctx, student)
	if err != nil {
		panic(err)
	}

	student.Name = "MoMo"
	modified, err = mapper.Update(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}

	student.ClassId = 12
	modified, err = mapper.ReplaceOne(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}

	modified, err = mapper.Delete(ctx, student)
	if err != nil {
		panic(err)
	}
	if modified < 1 {
		panic("not match")
	}
}

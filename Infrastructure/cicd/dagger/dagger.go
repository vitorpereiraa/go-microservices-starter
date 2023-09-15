package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main(){
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	source := client.Container().
		From("golang:1.21.0").
		WithDirectory("TodoList", client.Host().Directory("./TodoList")).
		WithWorkdir("TodoList")
	
	build := source.
		WithExec([]string{"go", "build", "-o", "./todolist.exe"}).
		WithExec([]string{"ls", "-la"})

	result, err := build.Stdout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
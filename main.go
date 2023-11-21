package main

import (
	"context"
	"log"
	"os"
	"time"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func run(ctx context.Context) error {
	dc, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer dc.Close()

	rootDir := dc.Host().Directory(".")

	_, err = dc.Container(dagger.ContainerOpts{
		Platform: "linux/amd64",
	}).From("golang:1.21.4").
		WithMountedDirectory("/src", rootDir).
		WithWorkdir("/src").
		// Bust any cache:
		WithEnvVariable("CACHE_BUSTER", time.Now().Format(time.RFC3339Nano)).
		WithExec([]string{"go", "build"}).
		Sync(ctx)

	return err
}

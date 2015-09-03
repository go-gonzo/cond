package cond

import (
	"github.com/omeid/gonzo"
	"github.com/omeid/gonzo/context"
)

func If(cond bool, do gonzo.Stage) gonzo.Stage {
	if cond {
		return do
	}
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {
		for file := range in {
			out <- file
		}
		return nil
	}
}

func IsString(check, expect string, do gonzo.Stage) gonzo.Stage {
	if check == expect {
		return do
	}
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {
		for file := range in {
			out <- file
		}
		return nil
	}
}

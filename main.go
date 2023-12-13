package main

import (
	"context"
	"fmt"
	"time"

	provider "openfeature-provider-golang/provider"

	"github.com/open-feature/go-sdk/pkg/openfeature"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PrintFlagDetails(flag string, featureFlagEvaluation openfeature.BoolResolutionDetail) {
	fmt.Println("FLAG: ", flag, " \t\t featureFlagEvaluation: ", featureFlagEvaluation)
}
func main() {
	ctx := context.Background()
	flagServerURL := "localhost:8000"

	provider := provider.NewProvider(flagServerURL, "application1")
	openfeature.SetProvider(provider)

	strings := []string{
		"feature-flag-1",
		"feature-flag-2",
		"feature-flag-3",
		"feature-flag-4",
		"feature-flag-5",
		"feature-flag-6",
	}

	ticker := time.NewTicker(1 * time.Second)
	_ = timestamppb.Now()

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("\n\n%s\n", time.Now())

				for _, flag := range strings {
					ff := provider.BooleanEvaluation(
						ctx,
						flag,
						false,
						openfeature.FlattenedContext{},
					)
					PrintFlagDetails(flag, ff)
				}

			}
		}
	}()

	// Keep main function alive indefinitely
	select {}
}

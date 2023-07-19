/*
Copyright © 2023 ras0q

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var ErrCommandNotFound = fmt.Errorf("command not found")

// rootCmd represents the base command when called without any subcommands
func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "kinano-go",
		Short: "I am kinano v2",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("oisu-")
		},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context, outW io.Writer, errW io.Writer, args []string) error {
	rootCmd := rootCmd()
	rootCmd.AddCommand(callCmd())

	rootCmd.SetContext(ctx)
	rootCmd.SetOut(outW)
	rootCmd.SetErr(errW)
	rootCmd.SetArgs(args)

	if _, _, err := rootCmd.Find(args); err != nil {
		return fmt.Errorf("Find: %w: %w", ErrCommandNotFound, err)
	}

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("rootCmd: %w", err)
	}

	return nil
}

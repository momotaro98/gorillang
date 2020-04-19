package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/momotaro98/gorillang"
)

func NewCommandEncode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encode",
		Short: "人間の言葉をゴリラ語に変換します。使い方: gorillang encode おはよう",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("言葉を入れてください。")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			output := gorillang.Encode(args[0])
			fmt.Println(output)
		},
	}
	return cmd
}

func NewCommandDecode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "ゴリラ語を人間の言葉にします。使い方: gorillang decode ウホウォホッ！！ホッ ウホウォうほホホッ！ ウホ？ゥゥオホ！！",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("ゴリラ語を入れてください。")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			input := strings.Join(args, " ")
			output := gorillang.Decode(input)
			fmt.Println(output)
		},
	}
	return cmd
}

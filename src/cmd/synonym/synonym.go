package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/zhuharev/synonym_name.txt/src/synonym"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Synonym"
	app.Commands = cli.Commands{
		cli.Command{
			Name:      "parent",
			ShortName: "p",
			Action:    parent,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "Search parent for this name",
				},
				cli.StringFlag{
					Name:  "file, f",
					Value: "synonym_name.txt",
					Usage: "Input file",
				},
				cli.StringFlag{
					Name:  "encoding, e",
					Value: "dashcoma",
					Usage: "Input file encoding",
				},
			},
		},
	}

	app.Run(os.Args)
}

func parent(ctx *cli.Context) error {
	if ctx.String("name") == "" {
		return fmt.Errorf("Name should not be empty")
	}
	db, e := synonym.New(ctx.String("file"), ctx.String("encoding"))
	if e != nil {
		return e
	}
	parent := db.Parent(ctx.String("name"))
	fmt.Println(parent)
	return nil
}

// Copyright 2021 ZORALab Enterprise (hello@zoralab.com)
// Copyright 2021 "Holloway" Chew, Kean Ho (hollowaykeanho@gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package main is the main "Das Monteur" software application.
//
// This package covers the application's main executions and its command-line
// interfaces. Once it processed a valid command, it will then execute the
// responsible function from monteur/pkg/monteurcore package.
package main

import (
	"fmt"
	"os"

	"gitlab.com/zoralab/cerigo/os/args"
	"gitlab.com/zoralab/monteur/gopkg/monteur"
)

func main() {
	action := ""

	// setup CLI manager
	m := args.NewManager()
	m.Name = "Monteur"
	m.Description = `
A software manufacturing automation and assembling tools in one app.
Das Monteur - Getting the job done locally and remotely at scale!
`
	m.Version = monteur.VERSION
	m.Examples = []string{
		`$ monteur help`,
		`$ monteur setup`,
		`$ monteur test`,
		`$ monteur clean`,
		`$ monteur build`,
		`$ monteur release`,
		`$ monteur package`,
		`$ monteur compose`,
		`$ monteur publish`,
	}

	_ = m.Add(&args.Flag{
		Name:  "Help",
		Label: []string{"help", "--help", "-h"},
		Value: &action,
		Help:  "call for help",
		HelpExamples: []string{
			"$ monteur help",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Setup",
		Label: []string{"setup"},
		Value: &action,
		Help:  "run repository setup for test, develop, and etc",
		HelpExamples: []string{
			"$ monteur setup",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Test",
		Label: []string{"test"},
		Value: &action,
		Help:  "execute the test job",
		HelpExamples: []string{
			"$ monteur test",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Clean",
		Label: []string{"clean"},
		Value: &action,
		Help:  "execute the clean job",
		HelpExamples: []string{
			"$ monteur clean",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Release",
		Label: []string{"release"},
		Value: &action,
		Help:  "execute the release job",
		HelpExamples: []string{
			"$ monteur release",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Build",
		Label: []string{"build"},
		Value: &action,
		Help:  "execute the build job",
		HelpExamples: []string{
			"$ monteur build",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Package",
		Label: []string{"package"},
		Value: &action,
		Help:  "execute the package job",
		HelpExamples: []string{
			"$ monteur package",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Publish",
		Label: []string{"publish"},
		Value: &action,
		Help:  "execute the publish job",
		HelpExamples: []string{
			"$ monteur publish",
		},
	})

	_ = m.Add(&args.Flag{
		Name:  "Compose",
		Label: []string{"compose"},
		Value: &action,
		Help:  "execute the publication composition job",
		HelpExamples: []string{
			"$ monteur compose",
		},
	})

	// parse the CLI arguments
	m.Parse()

	// execute according to action
	switch action {
	case "help", "--help", "-h":
		fmt.Fprintf(os.Stderr, "%s", m.PrintHelp())
		return
	case "setup":
		os.Exit(monteur.Setup())
	case "test":
		os.Exit(monteur.Test())
	case "clean":
		os.Exit(monteur.Clean())
	case "release":
		os.Exit(monteur.Release())
	case "build":
		os.Exit(monteur.Build())
	case "package":
		os.Exit(monteur.Package())
	case "publish":
		os.Exit(monteur.Publish())
	case "compose":
		os.Exit(monteur.Compose())
	default:
		fmt.Fprintf(os.Stderr,
			"[ ERROR ] unknown action. Use 'help' to start.\n",
		)
		os.Exit(1)
	}

	os.Exit(0)
}

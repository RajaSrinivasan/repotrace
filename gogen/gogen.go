package gogen

import (
	"fmt"
	"log"
	"os"
	"time"

	"../repo"
	"../versions"
)

type GoGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (gg GoGen) Generate(v versions.Version, filename string) {

	gofilename := filename + ".go"
	gofile, err := os.Create(gofilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, gofilename)
		return
	}
	defer gofile.Close()

	fmt.Fprintf(gofile, "package %s\n", filename)
	fmt.Fprintln(gofile, "// Go package generator")
	fmt.Fprintf(gofile, "// File: %s\n", gofilename)
	fmt.Fprintf(gofile, "const buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(gofile, "const versionMajor = %d\n", v.Major)
	fmt.Fprintf(gofile, "const versionMinor = %d\n", v.Minor)
	fmt.Fprintf(gofile, "const versionBuild = %d\n", v.Build)
	fmt.Fprintf(gofile, "const repoURL = \"%s\"\n", v.Repo)
	fmt.Fprintf(gofile, "const branchName = \"%s\"\n", v.Branch)
	fmt.Fprintf(gofile, "const shortCommitId = \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(gofile, "const longCommitId = \"%s\"\n", v.LongCommitId)
}

func (gg GoGen) GenerateFromRepo(m *repo.Manifest, v versions.Version, filename string) {

	gofilename := filename + ".go"
	gofile, err := os.Create(gofilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, gofilename)
		return
	}
	defer gofile.Close()

	fmt.Fprintf(gofile, "package %s\n", filename)
	fmt.Fprintln(gofile, "// Go package generator")
	fmt.Fprintf(gofile, "// File: %s\n", gofilename)
	fmt.Fprintf(gofile, "const buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(gofile, "const versionMajor = %d\n", v.Major)
	fmt.Fprintf(gofile, "const versionMinor = %d\n", v.Minor)
	fmt.Fprintf(gofile, "const versionBuild = %d\n", v.Build)

	for _, prj := range m.Projects {
		//repo.Show(prj)

		if len(prj.Revision) == 0 {
			fmt.Fprintf(gofile, "\n// Project %s\n", prj.Name)
			if len(prj.Repo) > 0 {
				fmt.Fprintf(gofile, "const %sRepoURL = \"%s\"\n", prj.Path, prj.Repo)
				if len(prj.Revision) > 0 {
					fmt.Fprintf(gofile, "const %sRevision = \"%s\"\n", prj.Path, prj.Revision)
				}

				fmt.Fprintf(gofile, "const %sShortCommitId = \"%s\"\n", prj.Path, prj.ShortCommitId)
				fmt.Fprintf(gofile, "const %sLongCommitId = \"%s\"\n", prj.Path, prj.LongCommitId)
			}
		} else {
			fmt.Fprintf(gofile, "\n// Project %s Revision %s \n", prj.Name, prj.Revision)
		}
	}
}

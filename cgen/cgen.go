package cgen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/repotrace/repo"
	"github.com/RajaSrinivasan/repotrace/versions"
)

type CGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (cg CGen) Generate(v versions.Version, filename string) {

	hfilename := filename + ".h"
	hfile, err := os.Create(hfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, hfilename)
		return
	}
	defer hfile.Close()

	fmt.Fprintln(hfile, "// C header generator")
	fmt.Fprintf(hfile, "// File: %s.h\n", filename)
	fmt.Fprintf(hfile, "#define BUILD_TIME \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(hfile, "#define VERSION_MAJOR (%d)\n", v.Major)
	fmt.Fprintf(hfile, "#define VERSION_MINOR (%d)\n", v.Minor)
	fmt.Fprintf(hfile, "#define VERSION_BUILD (%d)\n", v.Build)
	fmt.Fprintf(hfile, "#define REPO_URL \"%s\"\n", v.Repo)
	fmt.Fprintf(hfile, "#define BRANCH_NAME \"%s\"\n", v.Branch)
	fmt.Fprintf(hfile, "#define SHORT_COMMIT_ID \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(hfile, "#define LONG_COMMIT_ID \"%s\"\n", v.LongCommitId)
}

func (cg CGen) GenerateFromRepo(m *repo.Manifest, v versions.Version, filename string) {

	outfilename := filename + ".h"
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, outfilename)
		return
	}
	defer outfile.Close()

	fmt.Fprintln(outfile, "// C header generator")
	fmt.Fprintf(outfile, "// File: %s\n", outfilename)

	fmt.Fprintf(outfile, "#define BUILD_TIME \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(outfile, "#define VERSION_MAJOR (%d)\n", v.Major)
	fmt.Fprintf(outfile, "#define VERSION_MINOR (%d)\n", v.Minor)
	fmt.Fprintf(outfile, "#define VERSION_BUILD (%d)\n", v.Build)

	for _, prj := range m.Projects {
		//repo.Show(prj)

		if len(prj.Revision) == 0 {
			fmt.Fprintf(outfile, "\n// Project %s\n", prj.Name)
			if len(prj.Repo) > 0 {
				//prjName := path.Base(prj.Name)
				fmt.Fprintf(outfile, "#define %sRepoURL \"%s\" \n", prj.Path, prj.Repo)
				if len(prj.Revision) > 0 {
					fmt.Fprintf(outfile, "#define %sRevision \"%s\"\n", prj.Path, prj.Revision)
				}

				fmt.Fprintf(outfile, "#define %sShortCommitId \"%s\" \n", prj.Path, prj.ShortCommitId)
				fmt.Fprintf(outfile, "#define %sLongCommitId \"%s\" \n", prj.Path, prj.LongCommitId)
			}
		} else {
			fmt.Fprintf(outfile, "\n// Project %s Revision %s \n", prj.Name, prj.Revision)
		}
	}
}

package inigen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/repotrace/repo"
	"github.com/RajaSrinivasan/repotrace/versions"
)

type IniGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (ig IniGen) Generate(v versions.Version, filename string) {

	inifile, err := os.Create(filename + ".ini")
	if err != nil {
		log.Printf("%v creating %s\n", err, filename)
		return
	}
	defer inifile.Close()

	fmt.Fprintf(inifile, "[versions]\n")
	fmt.Fprintf(inifile, "buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(inifile, "versionMajor = %d\n", v.Major)
	fmt.Fprintf(inifile, "versionMinor = %d\n", v.Minor)
	fmt.Fprintf(inifile, "versionBuild = %d\n", v.Build)
	fmt.Fprintf(inifile, "repoURL = \"%s\"\n", v.Repo)
	fmt.Fprintf(inifile, "branchName = \"%s\"\n", v.Branch)
	fmt.Fprintf(inifile, "shortCommitId = \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(inifile, "longCommitId = \"%s\"\n", v.LongCommitId)
}

func (ig IniGen) GenerateFromRepo(m *repo.Manifest, v versions.Version, filename string) {

	outfilename := filename + ".ini"
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, outfilename)
		return
	}
	defer outfile.Close()

	fmt.Fprintf(outfile, "[versions]\n")
	fmt.Fprintf(outfile, "buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(outfile, "versionMajor = %d\n", v.Major)
	fmt.Fprintf(outfile, "versionMinor = %d\n", v.Minor)
	fmt.Fprintf(outfile, "versionBuild = %d\n", v.Build)

	for _, prj := range m.Projects {
		//repo.Show(prj)

		if len(prj.Revision) == 0 {
			if len(prj.Repo) > 0 {
				//prjName := path.Base(prj.Name)
				fmt.Fprintf(outfile, "%sRepoURL= \"%s\" \n", prj.Path, prj.Repo)
				if len(prj.Revision) > 0 {
					fmt.Fprintf(outfile, "%sRevision= \"%s\"\n", prj.Path, prj.Revision)
				}

				fmt.Fprintf(outfile, "%sShortCommitId= \"%s\" \n", prj.Path, prj.ShortCommitId)
				fmt.Fprintf(outfile, "%sLongCommitId= \"%s\" \n", prj.Path, prj.LongCommitId)
			}
		}
	}
}

package adagen

import (
	"fmt"
	"log"
	"os"
	"time"
	"path"

	"../repo"
	"../versions"
)

type AdaGen int

func (a AdaGen) Generate(v versions.Version, filename string) {

	specfilename := filename + ".ads"
	specfile, err := os.Create(specfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, specfilename)
		return
	}
	defer specfile.Close()

	fmt.Fprintf(specfile, "package %s is\n", filename)
	fmt.Fprintln(specfile, "-- Ada spec generator")
	fmt.Fprintf(specfile, "-- File: %s.ads\n", filename)
	fmt.Fprintf(specfile, "    BUILD_TIME : constant String := \"%s\" ;\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(specfile, "    VERSION_MAJOR : constant := %d ;\n", v.Major)
	fmt.Fprintf(specfile, "    VERSION_MINOR : constant := %d ;\n", v.Minor)
	fmt.Fprintf(specfile, "    VERSION_BUILD : constant := %d ;\n", v.Build)
	fmt.Fprintf(specfile, "    REPO_URL : constant String := \"%s\" ;\n", v.Repo)
	fmt.Fprintf(specfile, "    BRANCH_NAME : constant String := \"%s\" ;\n", v.Branch)
	fmt.Fprintf(specfile, "    SHORT_COMMIT_ID : constant String := \"%s\" ;\n", v.ShortCommitId)
	fmt.Fprintf(specfile, "    LONG_COMMIT_ID : constant String := \"%s\" ;\n", v.LongCommitId)
	fmt.Fprintf(specfile, "end %s ;\n", filename)
}

func (a AdaGen) GenerateFromRepo(m *repo.Manifest, v versions.Version, filename string) {

	outfilename := filename + ".ads"
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, outfilename)
		return
	}
	defer outfile.Close()

	fmt.Fprintf(outfile, "package %s is\n", filename)
	fmt.Fprintln(outfile, "-- Ada spec generator")
	fmt.Fprintf(outfile, "-- File: %s.ads\n", filename)

	fmt.Fprintf(outfile, "    BUILD_TIME : constant String := \"%s\" ;\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(outfile, "    VERSION_MAJOR : constant := %d ;\n", v.Major)
	fmt.Fprintf(outfile, "    VERSION_MINOR : constant := %d ;\n", v.Minor)
	fmt.Fprintf(outfile, "    VERSION_BUILD : constant := %d ;\n", v.Build)

	for _, prj := range m.Projects {
		//repo.Show(prj)

		if len(prj.Revision) == 0 {
			fmt.Fprintf(outfile, "\n-- Project %s\n", prj.Name)
			if len(prj.Repo) > 0 {
				prjName := path.Base(prj.Name)
				fmt.Fprintf(outfile, "    %sRepoURL : constant := \"%s\" ;\n", prjName, prj.Repo)
				if len(prj.Revision) > 0 {
		           fmt.Fprintf(outfile, "    %sRevision : constant := \"%s\" ;\n", prjName, prj.Revision)
				}

				fmt.Fprintf(outfile, "    %sShortCommitId : constant := \"%s\" ;\n", prjName, prj.ShortCommitId)
				fmt.Fprintf(outfile, "    %sLongCommitId : constant := \"%s\" ;\n", prjName, prj.LongCommitId)
			}
		} else {
			fmt.Fprintf(outfile, "\n-- Project %s Revision %s \n", prj.Name, prj.Revision)
		}
	}
	fmt.Fprintf(outfile, "end %s ;\n", filename)	
}

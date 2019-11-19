package repo

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Default struct {
	XMLName  xml.Name `xml:"default"`
	Sync     string   `xml:"sync-j,attr"`
	Revision string   `xml:"revision,attr"`
	Upstream string   `xml:"master,attr"`
}

type Remote struct {
	XMLName xml.Name `xml:"remote"`
	Fetch   string   `xml:"fetch,attr"`
	Name    string   `xml:"name,attr"`
}

type Copyfile struct {
	XMLName xml.Name `xml:"copyfile"`
	Dest    string   `xml:"dest,attr"`
	Src     string   `xml:"src,attr"`
}

type Project struct {
	XMLName  xml.Name `xml:"project"`
	Remote   string   `xml:"remote,attr"`
	Name     string   `xml:"name,attr"`
	Revision string   `xml:"revision,attr"`
	Path     string   `xml:"path,attr"`
	//Copyfile Copyfile
}

type Manifest struct {
	XMLName  xml.Name  `xml:"manifest"`
	Default  Default   `xml:"default"`
	Remotes  []Remote  `xml:"remote"`
	Projects []Project `xml:"project"`
}

var manifest Manifest

func traceProject(prj Project) {
	if len(prj.Revision) == 0 {
		log.Printf("Project %s has no revision spec %s\n", prj.Name, prj.Revision)
	}
	log.Printf("Project %s Path %s Revision %s length %d\n", prj.Name, prj.Path, prj.Revision, len(prj.Revision))

}
func LoadManifest(mfpath string) error {
	mfname := mfpath
	mf, err := os.Open(mfname)
	if err != nil {
		log.Fatal(err)
	}
	defer mf.Close()

	mfdata, err := ioutil.ReadAll(mf)
	if err != nil {
		panic(err)
	}
	xml.Unmarshal(mfdata, &manifest)
	log.Printf("Default Repository Name %s url %s\n", manifest.Remotes[0].Name, manifest.Remotes[0].Fetch)

	for _, prj := range manifest.Projects {
		traceProject(prj)
	}

	return nil
}

package repo

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	"../versions"
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
	XMLName       xml.Name `xml:"project"`
	Remote        string   `xml:"remote,attr"`
	Name          string   `xml:"name,attr"`
	Revision      string   `xml:"revision,attr"`
	Path          string   `xml:"path,attr"`
	Repo          string
	Branch        string
	ShortCommitId string
	LongCommitId  string
	//Copyfile Copyfile
}

type Manifest struct {
	XMLName  xml.Name  `xml:"manifest"`
	Default  Default   `xml:"default"`
	Remotes  []Remote  `xml:"remote"`
	Projects []Project `xml:"project"`
}

type Generator interface {
	Generate(v versions.Version, filename string)
	GenerateFromRepo(m *Manifest, v versions.Version, filename string)
}

func Show(prj Project) {
	log.Printf("Project %s Path %s Revision %s Repo %s\n", prj.Name, prj.Path, prj.Revision, prj.Repo)
}
func LoadManifest(mfpath string) (Manifest, error) {

	var manifest Manifest

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

	return manifest, nil
}

func fillGapsProject(prj *Project) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	err := os.Chdir(prj.Path)
	if err != nil {
		log.Printf("Error (%s)\n", err)
		return
	}
	log.Printf("Working from %s\n", prj.Path)
	rem := versions.GetRemoteURL(".")
	prj.Repo = rem

	br := versions.GetBranchWithHead(".")
	prj.Branch = br
	cid, lcid := versions.GetCommitId(".", "qqq")
	prj.ShortCommitId = cid
	prj.LongCommitId = lcid
	Show(*prj)
	//log.Printf("Project %s Repo %s Branch %s Commit Id %s Long %s\n", prj.Name, prj.Repo, prj.Branch, prj.ShortCommitId, prj.LongCommitId)
}

func FillGaps(manifest *Manifest) {
	log.Println("Begin Fillgap")
	for id, _ := range manifest.Projects {
		fillGapsProject(&manifest.Projects[id])
	}
	log.Println("End Fillgap")
}

# repotrace - generate source code details for traceability

## Background

Applications distributed in binary form should be traceable back to its source. When using a revision control system like git, this would mean being able to extract the source regardless of how old or how many revisions it has undergone since.

It is particularly challenging in the case of large systems where source code from different repositories are combined to produce the full package. Building embedded linux systems with [yocto](https://www.yoctoproject.org) is an example. Typically the [repo](https://gerrit.googlesource.com/git-repo) is used to orchestrate these builds. This tool supports the processing of a repo manifest.xml file to walk through the component directories and gather source code details. The output file will include the details for each component.

This projectlet generates a fragment of source code that can be compiled into the application and retrieved at runtime to report the details as necessary.

## Usage

        ../bin/repotrace --help
        usage: repotrace [-h|--help] [-v|--verbose] [-r|--report-version]
                        [-f|--manifest "<value>"] [-m|--major <integer>] [-n|--minor
                        <integer>] [-b|--build <integer>] [-L|--language
                        (go|C|Ada|ini)] [-o|--output "<value>"]

                        generate source trace info

        Arguments:

        -h  --help            Print help information
        -v  --verbose         Verbose. Default: false
        -r  --report-version  report version. Default: false
        -f  --manifest        Repo manifest file 
        -m  --major           Major version. Default: 0
        -n  --minor           Minor version. Default: 0
        -b  --build           Build Number. Default: 999
        -L  --language        Language to output
        -o  --output          Output file base name. Default: revisions

## Basic Examples

### Go Application

In the case of go, the output file name is used as the package name. 

        ../../bin/repotrace -L go

        cat revisions.go 
        package revisions
        // Go package generator
        // File: revisions.go
        const buildTime = "Sat Nov 23 2019 05:19:35"
        const versionMajor = 0
        const versionMinor = 0
        const versionBuild = 999
        const repoURL = "git@github.com:RajaSrinivasan/repotrace.git"
        const branchName = "master"
        const shortCommitId = "89f2267"
        const longCommitId = "89f2267b90c09bab344ddbb3a2c0afcee3785850"

## repo examples

In the following examples, a simple [manifest](https://github.com/RajaSrinivasan/myprojects.git) file is used :

        <?xml version="1.0" encoding="UTF-8"?>
        <manifest>

        <default sync-j="4" revision="master" upstream="master"/>

        <remote fetch="https://gitlab.com/"  name="gitlab"/>
        <remote fetch="git://github.com/"  name="github"/>

        <project remote="github"  name="RajaSrinivasan/srctrace.git"        path="srctrace"/>
        <project remote="gitlab"  name="privatetutor/projectlets/go.git"    path="go"/>
        </manifest>

In this example 2 different projects are managed one from gitlab.com and another from github.com. repotrace can process this manifest file and gather the revision info for each project.

### Go Application

        ../../../bin/repotrace -f .repo/manifest.xml -L go
        2019/11/23 05:40:57 Reporting versions for .repo/manifest.xml
        2019/11/23 05:40:57 Default Repository Name gitlab url https://gitlab.com/
        2019/11/23 05:40:57 Begin Fillgap
        2019/11/23 05:40:57 Working from srctrace
        2019/11/23 05:40:57 Project RajaSrinivasan/srctrace.git Path srctrace Revision  Repo git://github.com/RajaSrinivasan/srctrace.git
        2019/11/23 05:40:57 Working from go
        2019/11/23 05:40:57 Project privatetutor/projectlets/go.git Path go Revision  Repo https://gitlab.com/privatetutor/projectlets/go.git
        2019/11/23 05:40:57 End Fillgap

        cat revisions.go
        package revisions
        // Go package generator
        // File: revisions.go
        const buildTime = "Sat Nov 23 2019 05:40:57"
        const versionMajor = 0
        const versionMinor = 0
        const versionBuild = 999

        // Project RajaSrinivasan/srctrace.git
        const srctraceRepoURL = "git://github.com/RajaSrinivasan/srctrace.git"
        const srctraceShortCommitId = "5090875"
        const srctraceLongCommitId = "5090875cd44ce4be91b042e66a5122c5eec90adb"

        // Project privatetutor/projectlets/go.git
        const goRepoURL = "https://gitlab.com/privatetutor/projectlets/go.git"
        const goShortCommitId = "1f80b2a"
        const goLongCommitId = "1f80b2a0023a6b69d958b7c415472a325b4b0ba8"

As can be seen above, for each project, the repository info is gathered and the code fragment is generated.

## Complex manifests

Typical yocto based manifests tend to be more complex. for example, many projects are included by specifying exact commits. For example:


        <project remote="github" revision="c2b641c8a0c4fd71fcb477d788a740c2c26cddce" upstream="rocko"  name="YOCTO/poky"                    path="os/sources/poky"/>
        <project remote="github" revision="470cd54e44913f81c76538641bbdd80574624677" upstream="rocko"  name="YOCTO/meta-freescale"          path="os/sources/meta-freescale"/>
  

In such cases, since the revision is explicitly specified, the information is not generated in the output.

## Installing

        git clone http://github.com/RajaSrinivasan/repotrace.git
        cd repotrace
        make dependencies
        make all
        cp ../bin/repotrace <desired location>



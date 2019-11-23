# repotrace - generate source code details for traceability

## Background

Applications distributed in binary form should be traceable back to its source. When using a revision control system like git, this would mean being able to extract the source regardless of how old or home any revisions it has undergone since.

It is particularly challenging in the case of large systems where source code from different repositories are combined to produce the full package. Building embedded linux systems with [yocto](https://www.yoctoproject.org) is an example. Typically the [repo](https://gerrit.googlesource.com/git-repo) is used to orchestrate these builds. This tool supports the processing of a repo manifest.xml file to walk through the component directories and gather source code details. The output file will include the details for each component.

This projectlet generates a fragment of source code that can be compiled into the application and retrieved at runtime to report the details as necessary.

## Usage

        usage: repotrace [-h|--help] [-v|--verbose] [-m|--major <integer>] [-n|--minor
                        <integer>] [-b|--build <integer>] [-L|--language
                        (C|C++|Ada|python|go)] [-o|--output "<value>"]

                        generate source trace info

        Arguments:

        -h  --help      Print help information
        -v  --verbose   Verbose. Default: false
        -m  --major     Major version. Default: 0
        -n  --minor     Minor version. Default: 0
        -b  --build     Build Number. Default: 999
        -L  --language  Language to output
        -o  --output    Output file base name. Default: revisions

## Examples
### Go Application

In the case of go, the output file name is used as the package name. 

        ../../bin/repotrace -L go
        2019/11/15 11:20:19 repotrace

        cat revisions.go
        package revisions
        // Go package generator
        // File: revisions.h
        const buildTime = "Fri Nov 15 2019 11:20:19"
        const versionMajor = 0
        const versionMinor = 0
        const versionBuild = 999
        const repoURL = "git@gitlab.com:privatetutor/projectlets/go.git"
        const branchName = "master"
        const shortCommitId = "3e5cabc"
        const longCommitId = "3e5cabcc45a90d03202a04251713d573fbf6a807"

### C Application

        ../../bin/repotrace --language C -m 2 -n 3 -b 234
        2019/11/07 05:46:03 repotrace
        $ cat revisions.h
        // C header generator
        // File: revisions.h
        #define BUILD_TIME "Thu Nov 7 2019 05:46:03"
        #define VERSION_MAJOR (2)
        #define VERSION_MINOR (3)
        #define VERSION_BUILD (234)
        #define REPO_URL "git@gitlab.com:privatetutor/projectlets/go.git"
        #define BRANCH_NAME "master"
        #define SHORT_COMMIT_ID "cefa722"
        #define LONG_COMMIT_ID “cefa72267cc4d3a07fbf5e672b0053116d582aa7”

### Ada Application

In the case of Ada, the output filename is used to generate an entire package spec.

        ../../bin/repotrace --language Ada -m 2 -n 3 -b 234
        2019/11/07 05:48:39 repotrace
        $ cat revisions.ads
        package revisions is
        -- Ada spec generator
        -- File: revisions.ads
            BUILD_TIME : constant String := "Thu Nov 7 2019 05:48:39" ;
            VERSION_MAJOR : constant := 2 ;
            VERSION_MINOR : constant := 3 ;
            VERSION_BUILD : constant := 234 ;
            REPO_URL : constant String := "git@gitlab.com:privatetutor/projectlets/go.git" ;
            BRANCH_NAME : constant String := "master" ;
            SHORT_COMMIT_ID : constant String := "cefa722" ;
            LONG_COMMIT_ID : constant String := "cefa72267cc4d3a07fbf5e672b0053116d582aa7" ;
        end revisions ;

### Ini Files

        ../bin/repotrace --language ini
        2019/11/15 21:31:19 repotrace

        cat revisions
        [versions]
        buildTime = "Fri Nov 15 2019 21:31:20"
        versionMajor = 0
        versionMinor = 0
        versionBuild = 999
        repoURL = "git@gitlab.com:privatetutor/projectlets/go.git"
        branchName = "master"
        shortCommitId = "ec60c09"
        longCommitId = "ec60c0924b215757ed2ba1edb2ff3cdd0aac84d3"

## Installing

        git clone http://github.com/RajaSrinivasan/repotrace.git
        cd repotrace
        make dependencies
        make all
        cp ../bin/repotrace <desired location>



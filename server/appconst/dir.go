package appconst

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/futil"
)

// errors
var (
	ErrProjectRootNotFound1 = errutil.NewFactory("project root not found from: %q")
)

// ProjectRootDir project root dir, search from current working directory
var ProjectRootDir = func() (dir string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if dir, err = filepath.Abs(pwd); err != nil {
		panic(err)
	}

	for {
		switch dir {
		case "", ".", "/":
			panic(ErrProjectRootNotFound1.New(nil, pwd))
		}

		// must exist dir
		serverdir := filepath.Join(dir, "server")
		dbmigratedir := filepath.Join(serverdir, "dbmigrate")
		webdir := filepath.Join(dir, "web")

		// optional set (dev)
		gitdir := filepath.Join(dir, ".git")
		dockerdir := filepath.Join(dir, "docker")

		// optional set (prod)
		projectbin := filepath.Join(dir, strings.ToLower(AppName))

		if futil.IsDir(serverdir) && futil.IsDir(dbmigratedir) && futil.IsExist(webdir) {
			// dev
			if filepath.Base(dir) == strings.ToLower(AppName) && futil.IsDir(gitdir) && futil.IsDir(dockerdir) {
				return dir
			}

			// production
			if futil.IsExist(projectbin) {
				return dir
			}
		}

		dir = filepath.Dir(dir)
	}
}()

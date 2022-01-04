package libdeb

import (
	"fmt"

	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/archive/deb"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/liblog"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libmonteur"
	"gitlab.com/zoralab/monteur/gopkg/monteur/internal/libpackager"
)

func Changelog(pkg *libmonteur.TOMLPackage,
	variables *map[string]interface{},
	log *liblog.Logger) (err error) {
	var d *deb.Data
	var ok bool
	var packagePath string
	var app *libmonteur.Software

	log.Info("Creating %s changelog ...", libmonteur.CHANGELOG_DEB)

	// process critical variables
	app, ok = (*variables)[libmonteur.VAR_APP].(*libmonteur.Software)
	if !ok {
		panic("MONTEUR DEV: why is VAR_APP missing?")
	}

	// process package pathing
	packagePath, err = libpackager.UpdatePackagePath(variables,
		pkg,
		libmonteur.CHANGELOG_DEB,
		log.Info,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	// generate deb app data
	(*variables)[libmonteur.VAR_PACKAGE_OS] = pkg.OS
	(*variables)[libmonteur.VAR_PACKAGE_ARCH] = pkg.Arch

	d, err = createAppData(app, *variables, pkg, packagePath)
	if err != nil || d == nil {
		return err
	}

	// generate changelog entries
	err = d.CreateChangelog(packagePath)
	if err != nil {
		return fmt.Errorf("%s: %s", libmonteur.ERROR_CHANGELOG, err)
	}

	log.Info("Creating %s changelog ➤ OK", libmonteur.CHANGELOG_DEB)
	return nil
}

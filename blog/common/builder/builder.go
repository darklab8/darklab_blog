package builder

import (
	"os"

	"github.com/darklab8/blog/blog/common/types"
	"github.com/darklab8/blog/blog/settings"

	"github.com/darklab8/go-utils/goutils/utils/utils_cp"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_os"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Builder struct {
	components []*Component
	build_path utils_types.FilePath
}

type BuilderOption func(b *Builder)

func NewBuilder(opts ...BuilderOption) *Builder {
	b := &Builder{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *Builder) RegComps(components ...*Component) {
	b.components = append(b.components, components...)
}

func (b *Builder) build(params types.GlobalParams) {
	os.RemoveAll(params.Buildpath.ToString())
	os.MkdirAll(params.Buildpath.ToString(), os.ModePerm)

	for _, comp := range b.components {
		comp.Write(params)
	}

	folders := utils_os.GetRecursiveDirs(settings.ProjectFolder)
	for _, folder := range folders {
		if utils_filepath.Base(folder) == "static" {
			utils_cp.Dir(folder.ToString(),
				utils_filepath.Join(settings.ProjectFolder, utils_types.FilePath(params.Buildpath.ToString()), "static").ToString())
		}
	}
}

func (b *Builder) BuildAll() {

	var siteRoot = settings.GetSiteRoot()
	b.build(types.GlobalParams{
		Buildpath:         "build",
		Theme:             types.ThemeDark,
		SiteRoot:          siteRoot,
		StaticRoot:        siteRoot + settings.StaticPrefix,
		OppositeThemeRoot: siteRoot + "light/",
	})
	b.build(types.GlobalParams{
		Buildpath:         "build/light",
		Theme:             types.ThemeLight,
		SiteRoot:          siteRoot + "light/",
		StaticRoot:        siteRoot + "light/" + settings.StaticPrefix,
		OppositeThemeRoot: siteRoot,
	})

}

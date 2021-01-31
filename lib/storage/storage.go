package storage

import (
	"github.com/spf13/afero"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
)

// UseDefault Return the filesystem interface with config name
func UseDefault() Fs {
	name, has := config.Setting.Default["storage"]
	if !has {
		exception.New("the default storage config does not set!", 404).
			Ctx(json.M{"Default": config.Setting.Default}).
			Throw()
	}
	return Use(name)
}

// Use Return the filesystem interface with config name
func Use(name string) Fs {
	setting, has := config.Setting.Storage[name]
	if !has {
		exception.New("Storage config not found!", 404).
			Ctx(json.M{"name": name, "Storage": config.Setting.Storage}).
			Throw()
	}
	return UseSetting(setting)
}

// UseSetting Get or create connections from given settings
func UseSetting(setting config.Storage) Fs {
	switch setting.Engine {
	case "osfs":
		if setting.Readonly {
			return OsFsRead(setting.Root)
		}
		return OsFs(setting.Root)
	case "mem":
		if setting.Readonly {
			return MemFsRead(setting.Namespace)
		}
		return MemFs(setting.Namespace)
	case "cos":
		if setting.Readonly {
			return CosFsRead(setting.Root)
		}
		return CosFs(setting.Root)
	default:
		return MemFsRead("default")
	}
}

// OsFs Return the Operating System Native filesystem interface
func OsFs(basepath string) Fs {
	return Fs{
		root: basepath,
		Afero: afero.Afero{
			Fs: afero.NewBasePathFs(afero.NewOsFs(), basepath),
		},
	}
}

// OsFsRead Return the readonly Operating System Native filesystem interface
func OsFsRead(basepath string) Fs {
	return Fs{
		root: basepath,
		Afero: afero.Afero{
			Fs: afero.NewReadOnlyFs(afero.NewBasePathFs(afero.NewOsFs(), basepath)),
		},
	}
}

// CosFs Return the Qcloud COS  interface
func CosFs(name string) Fs {
	return Fs{
		root: name,
		Afero: afero.Afero{
			Fs: afero.NewBasePathFs(afero.NewOsFs(), "/base/path"),
		},
	}
}

// CosFsRead Return the readonly Qcloud COS interface
func CosFsRead(name string) Fs {
	return Fs{
		root: name,
		Afero: afero.Afero{
			Fs: afero.NewReadOnlyFs(afero.NewBasePathFs(afero.NewOsFs(), "/base/path")),
		},
	}
}

// MemFs Return the Memory Backed Storage filesystem interface
func MemFs(namespace string) Fs {
	return Fs{
		root: namespace,
		Afero: afero.Afero{
			Fs: afero.NewBasePathFs(afero.NewMemMapFs(), "/"+namespace),
		},
	}
}

// MemFsRead Return the readonly Memory Backed Storage filesystem interface
func MemFsRead(namespace string) Fs {
	return Fs{
		root: namespace,
		Afero: afero.Afero{
			Fs: afero.NewReadOnlyFs(afero.NewBasePathFs(afero.NewMemMapFs(), "/"+namespace)),
		},
	}
}

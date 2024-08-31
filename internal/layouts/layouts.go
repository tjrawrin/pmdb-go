package layouts

import (
	"encoding/json"
	"pmdb-go/web"
)

type ViteManifest struct {
	WebAssetsScriptJs struct {
		File    string `json:"file"`
		Name    string `json:"name"`
		Src     string `json:"src"`
		IsEntry bool   `json:"isEntry"`
	} `json:"web/assets/script.js"`
	WebAssetsStyleCSS struct {
		File    string `json:"file"`
		Name    string `json:"name"`
		Src     string `json:"src"`
		IsEntry bool   `json:"isEntry"`
	} `json:"web/assets/style.css"`
}

func GetAssets(filePath string) string {
	jsonBytes, err := web.WebFS.ReadFile("dist/.vite/manifest.json")
	if err != nil {
		panic("Error parsing manifest.json " + err.Error())
	}
	var manifest ViteManifest
	err = json.Unmarshal(jsonBytes, &manifest)
	if err != nil {
		panic("Error parsing manifest.json " + err.Error())
	}
	if filePath == "/public/style.css" {
		return "/public/" + manifest.WebAssetsStyleCSS.File
	}
	if filePath == "/public/script.js" {
		return "/public/" + manifest.WebAssetsScriptJs.File
	}
	return ""
}

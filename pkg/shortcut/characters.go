package shortcut

import "github.com/genshinsim/gcsim/pkg/core/keys"

var CharNameToKey = map[string]keys.Char{
	"traveleranemo":   keys.LumineAnemo,
	"travelergeo":     keys.LumineGeo,
	"travelerelectro": keys.LumineElectro,
	"travelerdendro":  keys.LumineDendro,
	"travelerhydro":   keys.LumineHydro,
	"travelerpyro":    keys.LuminePyro,
	"travelercryo":    keys.LumineCryo,
	"aether-anemo":    keys.AetherAnemo,
	"lumine-anemo":    keys.LumineAnemo,
	"aether-geo":      keys.AetherGeo,
	"lumine-geo":      keys.LumineGeo,
	"aether-electro":  keys.AetherElectro,
	"lumine-electro":  keys.LumineElectro,
	"aether-dendro":   keys.AetherDendro,
	"lumine-dendro":   keys.LumineDendro,
	"aether-hydro":    keys.AetherHydro,
	"lumine-hydro":    keys.LumineHydro,
	"aether-pyro":     keys.AetherPyro,
	"lumine-pyro":     keys.LuminePyro,
	"aether-cryo":     keys.AetherCryo,
	"lumine-cryo":     keys.LumineCryo,
	"aetheranemo":     keys.AetherAnemo,
	"lumineanemo":     keys.LumineAnemo,
	"aethergeo":       keys.AetherGeo,
	"luminegeo":       keys.LumineGeo,
	"aetherelectro":   keys.AetherElectro,
	"lumineelectro":   keys.LumineElectro,
	"aetherdendro":    keys.AetherDendro,
	"luminedendro":    keys.LumineDendro,
	"aetherhydro":     keys.AetherHydro,
	"luminehydro":     keys.LumineHydro,
	"aetherpyro":      keys.AetherPyro,
	"luminepyro":      keys.LuminePyro,
	"aethercryo":      keys.AetherCryo,
	"luminecryo":      keys.LumineCryo,
	"albedo":          keys.Albedo,
	"aloy":            keys.Aloy,
	"amber":           keys.Amber,
	"arlecchino":      keys.Arlecchino,
	"arle":            keys.Arlecchino,
	"barbara":         keys.Barbara,
	"barb":            keys.Barbara,
	"beidou":          keys.Beidou,
	"bennett":         keys.Bennett,
	"charlotte":       keys.Charlotte,
	"chongyun":        keys.Chongyun,
	"chong":           keys.Chongyun,
	"cyno":            keys.Cyno,
	"diluc":           keys.Diluc,
	"diona":           keys.Diona,
	"eula":            keys.Eula,
	"fischl":          keys.Fischl,
	"fish":            keys.Fischl,
	"amy":             keys.Fischl,
	"ganyu":           keys.Ganyu,
	"hutao":           keys.Hutao,
	"tao":             keys.Hutao,
	"ht":              keys.Hutao,
	"jean":            keys.Jean,
	"kaedeharakazuha": keys.Kazuha,
	"kazuha":          keys.Kazuha,
	"kaz":             keys.Kazuha,
	"kaeya":           keys.Kaeya,
	"kamisatoayaka":   keys.Ayaka,
	"ayaka":           keys.Ayaka,
	"kamisatoayato":   keys.Ayato,
	"ayato":           keys.Ayato,
	"keqing":          keys.Keqing,
	"keq":             keys.Keqing,
	"klee":            keys.Klee,
	"kujousara":       keys.Sara,
	"kujosara":        keys.Sara,
	"sara":            keys.Sara,
	"lisa":            keys.Lisa,
	"mona":            keys.Mona,
	"ningguang":       keys.Ningguang,
	"ning":            keys.Ningguang,
	"noelle":          keys.Noelle,
	"qiqi":            keys.Qiqi,
	"raidenshogun":    keys.Raiden,
	"raiden":          keys.Raiden,
	"herexcellencythealmightynarukamiogoshogodofthunder": keys.Raiden,
	"razor":                   keys.Razor,
	"rosaria":                 keys.Rosaria,
	"rosa":                    keys.Rosaria,
	"sangonomiyakokomi":       keys.Kokomi,
	"kokomi":                  keys.Kokomi,
	"koko":                    keys.Kokomi,
	"sayu":                    keys.Sayu,
	"sucrose":                 keys.Sucrose,
	"tartaglia":               keys.Tartaglia,
	"childe":                  keys.Tartaglia,
	"thoma":                   keys.Thoma,
	"venti":                   keys.Venti,
	"xiangling":               keys.Xiangling,
	"xl":                      keys.Xiangling,
	"xianyun":                 keys.Xianyun,
	"liuyun":                  keys.Xianyun,
	"cloudretainer":           keys.Xianyun,
	"xiao":                    keys.Xiao,
	"xingqiu":                 keys.Xingqiu,
	"xq":                      keys.Xingqiu,
	"xinyan":                  keys.Xinyan,
	"yanfei":                  keys.Yanfei,
	"yoimiya":                 keys.Yoimiya,
	"yoi":                     keys.Yoimiya,
	"yunjin":                  keys.Yunjin,
	"zhongli":                 keys.Zhongli,
	"zhong":                   keys.Zhongli,
	"zl":                      keys.Zhongli,
	"gorou":                   keys.Gorou,
	"aratakiitto":             keys.Itto,
	"itto":                    keys.Itto,
	"aratakitheoneandoniitto": keys.Itto,
	"shenhe":                  keys.Shenhe,
	"yae":                     keys.YaeMiko,
	"yaemiko":                 keys.YaeMiko,
	"yelan":                   keys.Yelan,
	"kukishinobu":             keys.Kuki,
	"kuki":                    keys.Kuki,
	"shikanoinheizou":         keys.Heizou,
	"heizou":                  keys.Heizou,
	"tighnari":                keys.Tighnari,
	"collei":                  keys.Collei,
	"dori":                    keys.Dori,
	"candace":                 keys.Candace,
	"nilou":                   keys.Nilou,
	"kusanali":                keys.Nahida,
	"lesserlordkusanali":      keys.Nahida,
	"nahida":                  keys.Nahida,
	"layla":                   keys.Layla,
	"faruzan":                 keys.Faruzan,
	"faru":                    keys.Faruzan,
	"wanderer":                keys.Wanderer,
	"scaramouche":             keys.Wanderer,
	"scara":                   keys.Wanderer,
	"kunikuzushi":             keys.Wanderer,
	"kuni":                    keys.Wanderer,
	"kabukimono":              keys.Wanderer,
	"alhaitham":               keys.Alhaitham,
	"haitham":                 keys.Alhaitham,
	"baizhu":                  keys.Baizhu,
	"dehya":                   keys.Dehya,
	"yaoyao":                  keys.Yaoyao,
	"mika":                    keys.Mika,
	"kaveh":                   keys.Kaveh,
	"kirara":                  keys.Kirara,
	"lyney":                   keys.Lyney,
	"lynette":                 keys.Lynette,
	"neuvillette":             keys.Neuvillette,
	"neuv":                    keys.Neuvillette,
	"chiefjusticeoffontaine":  keys.Neuvillette,
	"freminet":                keys.Freminet,
	"furina":                  keys.Furina,
	"furinadefontaine":        keys.Furina,
	"navia":                   keys.Navia,
	"demoiselle":              keys.Navia,
	"wriothesley":             keys.Wriothesley,
	"wrio":                    keys.Wriothesley,
	"chevreuse":               keys.Chevreuse,
	"chev":                    keys.Chevreuse,
	"gaming":                  keys.Gaming,
	"chiori":                  keys.Chiori,
	"clorinde":                keys.Clorinde,
	"emilie":                  keys.Emilie,
	"mualani":                 keys.Mualani,
}

package config

/*
Net - network data.
	- Data  - contains:
		- MaxConnections <int> - connected clients limit.
		- Port <int> - WS port.
*/
type Net struct {
	MaxConnections int `json:"max_connections"`
	Port           int `json:"port"`
}

/*
Engine - game engine.
	- Data  - contains:
		- FPS <int> - time interval for calculating the states of objects.
		- TimeStep <int> - step interval.
*/
type Engine struct {
	FPS      int `json:"fps"`
	TimeStep int `json:"time_step"`
}

/*
GameObjects - data about objects in the game world.
	- Data  - contains:
		- Players <Players> - data of characteristics objects of the Players.
		- Scene <Scene> - object within which game events take place.
		- Weapons <[]Weapon> - array of weapons settings.
		- GunParts <[]GunPart> - array of settings for the part object for assembling weapons.
		- GunModules <[]GunModule> - array of settings for the weapon module object.
		- Loots <[]Loot> - array of settings for an object that contains ammo, first aid kits and armor.
		- Runes <[]Rune> - array of objects with useful effects.
		- GunCases <[]GunCase> - array of settings for weapon part chests.
		- GunChests <[]GunChest> - array of settings for weapon module chests.
		- AmmoCases <[]AmmoCase> - array of settings for ammunition chests.
		- Blocks <[]Block> - array of settings for blocks(walls).
*/
type GameObjects struct {
	Players    Players     `json:"players"`
	Scene      Scene       `json:"scene"`
	Weapons    []Weapon    `json:"weapons"`
	GunParts   []GunPart   `json:"gun_parts"`
	GunModules []GunModule `json:"gun_modules"`
	Loots      []Loot      `json:"loots"`
	Runes      []Rune      `json:"runes"`
	GunCases   []GunCase   `json:"gun_cases"`
	GunChests  []GunChest  `json:"gun_chests"`
	AmmoCases  []AmmoCase  `json:"ammo_cases"`
	Blocks     []Block     `json:"blocks"`
}

/*
Players - data of characteristics objects of the Players.
	- Data  - contains:
		- Width <int> - player object's width.
		- Height <int> - player object's height.
		- MaxHealth <int> - limit of health.
		- Armor <int> - limit of armor.
		- Speed <int> - speed of player.
*/
type Players struct {
	Width     int `json:"width"`
	Height    int `json:"height"`
	MaxHealth int `json:"max_health"`
	MaxArmor  int `json:"max_armor"`
	Speed     int `json:"speed"`
}

/*
Weapon - weapon settings.
	- Data  - contains:
		- Kind <string> - weapon name.
		- RateOfFire <int> - rate of fire of the weapon is the firing interval.
		- BulletSpread <int> - bullet spread.
		- MaxParts <int> - number of weapon parts to assemble.
		- ReloadTime <int> - weapon reload time.
		- CartridgeCapacity <int> - maximum number of cartridges in the store.
		- CartridgeStores <int> - maximum number of weapon stores.
		- IsDefault <bool> - default weapon.
		- Modules <[]Module> - possible weapon modules.
		- BulletSpeed <int> - bullet speed.
		- BulletWidth <int> - bullet width.
		- BulletHeight <int> - bullet height.
		- BulletLifeTime <int> - bullet flight time.
		- BulletDamage <int> - bullet damage.
		- BulletEffects <[]Effect> - possible bullet effects.
		- BulletRadiusOfDestruction <int> - bullet radius of destrcution.
		- BulletTriggerRadius <int> - bullet triegger radius.
		- BulletTriggerDelay <int> - bullet response delay.
		- BulletRicochet <int> - count of bullet ricochet.
*/
type Weapon struct {
	Kind                      string   `json:"kind"`
	RateOfFire                int      `json:"rate_of_fire"`
	BulletSpread              int      `json:"bullet_spread"`
	MaxParts                  int      `json:"max_parts"`
	ReloadTime                int      `json:"reload_time"`
	CartridgeCapacity         int      `json:"cartridge_capacity"`
	CartridgeStores           int      `json:"cartridge_stores"`
	IsDefault                 bool     `json:"is_default"`
	Modules                   []Module `json:"modules"`
	BulletSpeed               int      `json:"bullet_speed"`
	BulletWidth               int      `json:"bullet_width"`
	BulletHeight              int      `json:"bullet_height"`
	BulletLifeTime            int      `json:"bullet_life_time"`
	BulletDamage              int      `json:"bullet_damage"`
	BulletEffects             []Effect `json:"bullet_effects"`
	BulletRadiusOfDestruction int      `json:"bullet_radius_of_destruction"`
	BulletTriggerRadius       int      `json:"bullet_trigger_radius"`
	BulletTriggerDelay        int      `json:"bullet_trigger_delay"`
	BulletRicochet            int      `json:"bullet_ricochet"`
}

/*
Effect - effect data for Players and Bullets.
Responsible for temporarily changing the characteristics of these objects.
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Speed <int> - variable speed of the object. Default value is 0.
		- Rate <int> - variable rate of the object is weapon. Default value is 0.
		- Chance <int> - probability of effect.
		- Duration <int> - duration of the effect.
*/
type Effect struct {
	Kind     string `json:"kind"`
	Speed    int    `json:"speed"`
	Rate     int    `json:"rate"`
	Chance   int    `json:"chance"`
	Duration int    `json:"duration"`
}

/*
Module - data about the modules that can be installed in the weapon.
	- Data  - contains:
		- Kind <string> - weapon name.
		- RateOfFire <int> - rate of fire of the weapon is the firing interval.
		- BulletSpread <int> - bullet spread.
		- ReloadTime <int> - weapon reload time.
		- CartridgeCapacity <int> - maximum number of cartridges in the store.
		- CartridgeStores <int> - maximum number of weapon stores.
		- BulletSpeed <int> - bullet speed.
		- BulletWidth <int> - bullet width.
		- BulletHeight <int> - bullet height.
		- BulletLifeTime <int> - bullet flight time.
		- BulletDamage <int> - bullet damage.
		- IsInstalled <int> - is pre-installed the module.
*/
type Module struct {
	Kind              string `json:"kind"`
	BulletSpeed       int    `json:"bullet_speed"`
	BulletWidth       int    `json:"bullet_width"`
	BulletHeight      int    `json:"bullet_height"`
	BulletLifeTime    int    `json:"bullet_life_time"`
	BulletDamage      int    `json:"bullet_damage"`
	ReloadTime        int    `json:"reload_time"`
	CartridgeCapacity int    `json:"cartridge_capacity"`
	CartridgeStores   int    `json:"cartridge_stores"`
	RateOfFire        int    `json:"rate_of_fire"`
	BulletSpread      int    `json:"bullet_spread"`
	IsInstalled       bool   `json:"is_installed"`
}

/*
Scene - object within which game events take place.
	- Data  - contains:
		- Width <int> - scene width.
		- Height <int> - scene height.
		- SizeCell <int> - size of scene's cell.
*/
type Scene struct {
	WidthCountCell  int           `json:"width_count_cell"`
	HeightCountCell int           `json:"height_count_cell"`
	SizeCell        int           `json:"size_cell"`
	PlayersDistResp int           `json:"players_dist_resp"`
	ObjectLimits    []ObjectLimit `json:"object_limit"`
}

type ObjectLimit struct {
	Kind      string `json:"kind"`
	Count     int    `json:"count"`
	RespDelay int    `json:"resp_delay"`
}

/*
GunPart - settings for the part object for assembling weapons.
	- Data  - contains:
		- Kind <string> - kind of weapon.
		- Width <int> - scene width.
		- Height <int> - scene height.
		- Chance <int> - probability of gun part.
*/
type GunPart struct {
	Kind   string `json:"kind"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Chance int    `json:"chance"`
}

/*
GunModule - settings for the weapon module object.
	- Data  - contains:
		- Kind <string> - kind of weapon.
		- Width <int> - scene width.
		- Height <int> - scene height.
		- Chance <int> - probability of gun module.
*/
type GunModule struct {
	Kind   string `json:"kind"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Chance int    `json:"chance"`
}

/*
Rune - object with useful effects.
	- Data  - contains:
		- Width <int> - scene width.
		- Height <int> - scene height.
		- Effect <Effect> - effect data for Players and Bullets.
*/
type Rune struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Effect Effect `json:"effect"`
}

/*
Loot - settings for an object that contains ammo, first aid kits and armor.
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Width <int> - scene width.
		- Height <int> - scene height.
		- Chance <int> - probability of loot.
*/
type Loot struct {
	Kind   string `json:"kind"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Value  []int  `json:"value"`
	Chance int    `json:"chance"`
}

/*
AmmoCase - settings for ammunition chests.
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Width <int> - object width.
		- Height <int> - object height.
		- Health <int> - object health.
		- LootList <[]string> - list of supports loots.
		- Chance <int> - probability of ammo case.
*/
type AmmoCase struct {
	Kind     string   `json:"kind"`
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Health   int      `json:"health"`
	LootList []string `json:"loot_list"`
	Chance   int      `json:"chance"`
}

/*
GunCase - settings for weapon part chests.
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Width <int> - object width.
		- Height <int> - object height.
		- Health <int> - object health.
		- GunPartList <[]string> - list of supports gun parts.
		- Chance <int> - probability of gun case.
*/
type GunCase struct {
	Kind        string   `json:"kind"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Health      int      `json:"health"`
	GunPartList []string `json:"gun_part_list"`
	Chance      int      `json:"chance"`
}

/*
GunChest - settings for weapon module chests.
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Width <int> - object width.
		- Height <int> - object height.
		- Health <int> - object health.
		- GunPartList <[]string> - list of supports gun parts.
		- Chance <int> - probability of gun chest.
*/
type GunChest struct {
	Kind          string   `json:"kind"`
	Width         int      `json:"width"`
	Height        int      `json:"height"`
	Health        int      `json:"health"`
	GunModuleList []string `json:"gun_module_list"`
	Chance        int      `json:"chance"`
}

/*
Block - settings for blocks(walls).
	- Data  - contains:
		- Kind <string> - kind of effect.
		- Width <int> - object width.
		- Height <int> - object height.
		- Health <int> - object health.
		- GunPartList <[]string> - list of supports gun parts.
		- Chance <int> - respawn chance.
*/
type Block struct {
	Kind        string   `json:"kind"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Health      int      `json:"health"`
	ContentList []string `json:"content_list"`
	Chance      int      `json:"chance"`
}

/*
configData - game server configuration data.
	- Data  - contains:
		- Net <Net> - network data.
		- Engine <Engine> - game engine.
		- GameObjects <GameObjects> - data about objects in the game world.
*/
type configData struct {
	Net         Net         `json:"net"`
	Engine      Engine      `json:"engine"`
	GameObjects GameObjects `json:"game_objects"`
}

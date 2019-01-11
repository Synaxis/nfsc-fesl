package model

func NewHeroStats() HeroStats {
	return HeroStats{
		Medals:       newMedals(),
		InGameEvents: newInGameEvents(),
		Inventory:    newInventory(),
		ByFaction:    newByFaction(),
		ByKit:        newByKit(),
		Kits:         newKits(),
		Weapons:      newWeapons(),
		Vehicles:     newVehicles(),
		Maps:         newMaps(),
	}
}

type HeroStats struct {
	// FacialHairType: 0 = None, 102..109
	FacialHairType string `json:"c_ft" stats:"c_ft" default:"0"`

	// Faction: 1 = National, 2 = Royal
	Faction string `json:"c_team" stats:"c_team" default:"1"`

	// Hair Color: 1..5
	HairColor string `json:"c_hrc" stats:"c_hrc" default:"1"`

	// HairStyle 0 = bald, 82..87 some hair
	HairStyle string `json:"c_hrs" stats:"c_hrs" default:"0"`

	// SkinColor 1..9, 1 = darkest, 9 = lighest
	SkinColor string `json:"c_skc" stats:"c_skc" default:"1" min:"1" max:"9"`

	// ut=0 k=c_ltp v=9301.0000 t= pt=0
	CLtp  string `json:"c_ltp" stats:"c_ltp" default:"9301.0000"` // 9260.0000
	CLtm  string `json:"c_ltm" stats:"c_ltm" default:"0.0000"`    // 9260.0000
	CFhrs string `json:"c_fhrs" stats:"c_fhrs" default:"0.0000"`  // Capture-Flag Hours played?
	CSlm  string `json:"c_slm" stats:"c_slm" default:"0.0000"`    // 0.0000
	Cdm   string `json:"cdm" stats:"cdm" default:"0.0000"`
	Edm   string `json:"edm" stats:"edm" default:"0.0000"`

	PlayerType  string `json:"c_kit" default:"0" stats:"c_kit"` // 0 = Commando, 1 = Soldier, 2 = Gunner
	HeroPoints  string `json:"c_wallet_hero" stats:"c_wallet_hero" default:"0.0000"`
	ValorPoints string `json:"c_wallet_valor" stats:"c_wallet_valor" default:"0.0000"`
	GamesPlayed string `json:"games" stats:"games" default:"0.000"`
	Elo         string `json:"elo" stats:"elo" default:"1000.0000"`
	Level       string `json:"level" stats:"level" default:"1"`
	XP          string `json:"xp" stats:"xp" default:"0.0000"`

	TimePlayed   string `json:"ct" stats:"ct" default:"0.0000"`
	Kills        string `json:"ki" stats:"ki" default:"0.0000"`
	Deaths       string `json:"dt" stats:"dt" default:"0.0000"`
	Suicides     string `json:"su" stats:"su" default:"0.0000"`
	Wins         string `json:"win" stats:"win" default:"0.0000"`
	Losses       string `json:"los" stats:"los" default:"0.0000"`
	BulletsFired string `json:"fi" stats:"fi" default:"0.0000"`
	BulletsHit   string `json:"hi" stats:"hi" default:"0.0000"`

	// Score
	Score        string `json:"rs" stats:"rs" default:"0.0000"`
	TeamScore    string `json:"ts" stats:"ts" default:"0.0000"`
	SkillScore   string `json:"ss" stats:"ss" default:"0.0000"`
	CaptureScore string `json:"cs" stats:"cs" default:"0.0000"`

	PrestigeLevel  string `json:"prs" stats:"prs" default:"0.0000"` // reseted on levelup
	PrestigePoints string `json:"ppt" stats:"ppt" default:"0.0000"` // reseted on levelup

	PlayedTutorial         string `json:"c_tut" stats:"c_tut" default:"1.0000"`
	AwayBonusMedal         string `json:"awybt" stats:"awybt" default:"0.0000"`
	DailyMissionsCompleted string `json:"dmc" stats:"dmc" default:"0.0000"`

	Globals
	Medals
	Rounded
	Captures
	Streaks
	ByFaction
	ByKit
	Maps
	Vehicles
	Weapons
	Kits
	Teamplay
	Inventory
	InGameEvents
}

type Globals struct {
	GlobalScore     string `json:"gsco" stats:"gsco" default:"0.0000"`
	GlobalXP        string `json:"expts" stats:"expts" default:"0.0000"`
	GlobalAwayBonus string `json:"bnspt" stats:"bnspt" default:"0.0000"`
}

// Medals, awards and missions
type Medals struct {
	WebKey           map[int]string `json:"c_aprweb" stats:"c_aprweb" default:"0" start:"0" end:"6"` //newwer clients
	Awards            map[int]string `json:"aw" stats:"aw" default:"6000" start:"6000" end:"6182"` //when lvl 30 more missions are avail
	Missions          map[int]string `json:"mid" stats:"mid" default:"0" start:"0" end:"2"`
	PlayedMissions    map[int]string `json:"c_mid" stats:"c_mid" default:"0.0000" start:"0" end:"2"`
	CompletedMissions map[int]string `json:"c_cmid" stats:"c_cmid" default:"0.0000" start:"0" end:"2"`
	MissionOne        map[int]string `json:"m0c" stats:"m0c" default:"0.0000" start:"0" end:"5"`
	MissionTwo        map[int]string `json:"m1c" stats:"m1c" default:"0.0000" start:"0" end:"5"`
	MissionThree      map[int]string `json:"m2c" stats:"m2c" default:"0.0000" start:"0" end:"5"`

	// ut=0 k=c_wmid0 v=6000.0000 t= pt=0
	WishedMissions map[int]string `json:"c_wmid" stats:"c_wmid" default:"6000.0000" start:"0" end:"2"`
}

func newMedals() Medals {
	return Medals{
		WebKey:           map[int]string{},
		Awards:            map[int]string{},
		Missions:          map[int]string{},
		PlayedMissions:    map[int]string{},
		CompletedMissions: map[int]string{},
		MissionOne:        map[int]string{},
		MissionTwo:        map[int]string{},
		MissionThree:      map[int]string{},
		WishedMissions:    map[int]string{},
	}
}

type Rounded struct {
	StartRank                string `json:"startLVL" stats:"startLVL" default:"0.0000"`
	RoundXP                  string `json:"roundXP" stats:"roundXP" default:"0.0000"`
	RoundBonusXP             string `json:"roundBXP" stats:"roundBXP" default:"0.0000"`
	RoundValorPoints         string `json:"roundVP" stats:"roundVP" default:"0.0000"`
	RoundBonusValorPoints    string `json:"roundBVP" stats:"roundBVP" default:"0.0000"`
	RoundHeroPoints          string `json:"roundHP" stats:"roundHP" default:"0.0000"`
	RoundPrestigePoints      string `json:"roundPP" stats:"roundPP" default:"0.0000"`
	RoundBonusPrestigePoints string `json:"roundBPP" stats:"roundBPP" default:"0.0000"`
	TotalPrestigePoints      string `json:"totalPP" stats:"totalPP" default:"0.0000"`
}

// Captures
type Captures struct {
	Captures        string `json:"cpc" stats:"cpc" default:"0.0000"` // With neutralizations
	CapturesAssists string `json:"cpa" stats:"cpa" default:"0.0000"` // With neutralization assists
	CaptureDefends  string `json:"cpd" stats:"cpd" default:"0.0000"`
	RocketCaptures  string `json:"rc" stats:"rc" default:"0.0000"`
}

type Streaks struct {
	KillStreak string `json:"ks" stats:"ks" default:"0.0000"` // Bests
	DeadStreak string `json:"ds" stats:"ds" default:"0.0000"` // Bests
}

// c_team
type ByFaction struct {
	ScoreByFaction  map[int]string `json:"ft_rs" stats:"ft_rs" default:"0.0000" start:"0" end:"1"`
	KillsByFaction  map[int]string `json:"ft_ki" stats:"ft_ki" default:"0.0000" start:"0" end:"1"`
	DeathsByFaction map[int]string `json:"ft_dt" stats:"ft_dt" default:"0.0000" start:"0" end:"1"`
	WinsByFaction   map[int]string `json:"ft_win" stats:"ft_win" default:"0.0000" start:"0" end:"1"`
	LossesByFaction map[int]string `json:"ft_los" stats:"ft_los" default:"0.0000" start:"0" end:"1"`
}

func newByFaction() ByFaction {
	return ByFaction{
		ScoreByFaction:  map[int]string{},
		KillsByFaction:  map[int]string{},
		DeathsByFaction: map[int]string{},
		WinsByFaction:   map[int]string{},
		LossesByFaction: map[int]string{},
	}
}

// c_kit
type ByKit struct {
	ScoreByKit  map[int]string `json:"fc_rs" stats:"fc_rs" default:"0.0000" start:"0" end:"2"`
	KillsByKit  map[int]string `json:"fc_ki" stats:"fc_ki" default:"0.0000" start:"0" end:"2"`
	DeathsByKit map[int]string `json:"fc_dt" stats:"fc_dt" default:"0.0000" start:"0" end:"2"`
	WinsByKit   map[int]string `json:"fc_win" stats:"fc_win" default:"0.0000" start:"0" end:"2"`
	LossesByKit map[int]string `json:"fc_los" stats:"fc_los" default:"0.0000" start:"0" end:"2"`
}

func newByKit() ByKit {
	return ByKit{
		ScoreByKit:  map[int]string{},
		KillsByKit:  map[int]string{},
		DeathsByKit: map[int]string{},
		WinsByKit:   map[int]string{},
		LossesByKit: map[int]string{},
	}
}

type Maps struct {
	MapTimePlayed map[int]string `json:"m_ct"`  //  stats:"m_ct" start:? end:?
	MapWins       map[int]string `json:"m_win"` //  stats:"m_win" start:? end:?
	MapLoses      map[int]string `json:"m_los"` //  stats:"m_los" start:? end:?
}

func newMaps() Maps {
	return Maps{
		MapTimePlayed: map[int]string{},
		MapWins:       map[int]string{},
		MapLoses:      map[int]string{},
	}
}

type Vehicles struct {
	VehicleTimeInObject map[int]string `json:"tv"`    //  stats:"tv" start:? end:?
	VehicleKills        map[int]string `json:"kv"`    //  stats:"kv" start:? end:?
	VehicleKilledBy     map[int]string `json:"dfv"`   //  stats:"dfv" start:? end:?
	VehicleRoadKills    map[int]string `json:"kvr"`   //  stats:"kvr" start:? end:?
	VehicleDestroyed    map[int]string `json:"dstrv"` //  stats:"dstrv" start:? end:?
	VehicleDeaths       map[int]string `json:"div"`   //  stats:"div" start:? end:?
}

func newVehicles() Vehicles {
	return Vehicles{
		VehicleTimeInObject: map[int]string{},
		VehicleKills:        map[int]string{},
		VehicleKilledBy:     map[int]string{},
		VehicleRoadKills:    map[int]string{},
		VehicleDestroyed:    map[int]string{},
		VehicleDeaths:       map[int]string{},
	}
}

type Weapons struct {
	WeaponTimeInObject map[int]string `json:"tw"`  // stats:"tw" start:? end:?
	WeaponTimeInKit    map[int]string `json:"twk"` // stats:"twk" start:? end:?
	WeaponKills        map[int]string `json:"kw"`  // stats:"kw" start:? end:?
	WeaponKilledBy     map[int]string `json:"dfw"` // stats:"dfw" start:? end:?
	WeaponBulletsFired map[int]string `json:"sw"`  // stats:"sw" start:? end:?
	WeaponBulletsHit   map[int]string `json:"hw"`  // stats:"hw" start:? end:?
	WeaponDeaths       map[int]string `json:"dww"` // stats:"dww" start:? end:?
}

func newWeapons() Weapons {
	return Weapons{
		WeaponTimeInObject: map[int]string{},
		WeaponTimeInKit:    map[int]string{},
		WeaponKills:        map[int]string{},
		WeaponKilledBy:     map[int]string{},
		WeaponBulletsFired: map[int]string{},
		WeaponBulletsHit:   map[int]string{},
		WeaponDeaths:       map[int]string{},
	}
}

type Kits struct {
	KitKills    map[int]string `json:"kk" stats:"kk" default:"0.0000" start:"0" end:"2"`
	KitKilledBy map[int]string `json:"kkb" stats:"kkb" default:"0.0000" start:"0" end:"2"`
}

func newKits() Kits {
	return Kits{
		KitKills:    map[int]string{},
		KitKilledBy: map[int]string{},
	}
}

type Teamplay struct {
	DamageAssists string `json:"ka" stats:"ka" default:"0.0000"`
	HealSelf      string `json:"he" stats:"he" default:"0.0000"`
	DriverAssists string `json:"drka" stats:"drka" default:"0.0000"`
}

type Inventory struct {
	// ut=0 k=c_apr v=0.0000 t=10 pt=0
	Appearance []string `json:"c_apr" stats:"c_apr" default:"10;978;980"`
	Emotes     []string `json:"c_emo" stats:"c_emo" default:"5000;5007;5016;0;0;0;0;0;0"`
	// ut=0 k=c_eqp v=0.0000 t=3190;0;0;3155;0;0;0;0;0;0 pt=0
	Equipment []string `json:"c_eqp" stats:"c_eqp" default:"3190;0;0;3155;0;0;0;0;0;0"` // 3012;3010;2075;3155;2005;0;0;0;0;0
	Items     []string `json:"c_items" stats:"c_items" default:"2026;2027;2028;2031;2032;2033;2046;2047;2048;2055;2056;2057;2091"`
}

func newInventory() Inventory {
	return Inventory{
		Appearance: []string{},
		Emotes:     []string{},
		Equipment:  []string{},
		Items:      []string{},
	}
}

type InGameEvents struct {
	InGameEvents map[int]string `json:"ige"` // stats:"ige" start:? end:?
}

func newInGameEvents() InGameEvents {
	return InGameEvents{
		InGameEvents: map[int]string{},
	}
}

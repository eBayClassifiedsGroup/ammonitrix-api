package config

var DefaultConfig = &Config{
	Listen: Listen{
		Port: ":5859",
	},
	Elastic: Elastic{
		Host:          "localhost",
		Port:          ":9200",
		IndexName:     "ammonitrix",
		MetaDataIndex: "ammonitrix_meta",
	},
}

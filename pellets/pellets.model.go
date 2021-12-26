package pellets

type PelletPrice struct {
	Error                        string `json:"error"`
	BestPreis                    int    `json:"bestPreis"`
	Bundesland                   string
	Fahrzeit                     string
	Lieferzeit                   string
	Gebiet                       string
	Lager                        string
	Lager_id                     string
	NettoEinzelpreisSK           float32 `json:"nettoEinzelpreisSK"`
	NettoEinzelpreis             float32
	NettoEinzelpreisCal          float32
	NettoEinzelpreisMindestmenge float32
}

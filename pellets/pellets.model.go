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
	NettoEinzelpreisSK           float64 `json:"nettoEinzelpreisSK"`
	NettoEinzelpreis             float64
	NettoEinzelpreisCal          float64
	NettoEinzelpreisMindestmenge float64
}

const (
	Loose       = "P101"
	Sacked      = "P104"
	UnitKG      = "kg"
	UnitPallete = "Pal."
)

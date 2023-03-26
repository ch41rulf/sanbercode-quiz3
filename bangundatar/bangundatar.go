package bangundatar

type SegitigaSamaSisi struct {
	Alas   string  `json:"alas"`
	Tinggi string  `json:"tinggi"`
	Hitung float64 `json:"hitung"`
}

type Persegi struct {
	Sisi   string  `json:"sisi"`
	Hitung float64 `json:"hitung"`
}

type PersegiPanjang struct {
	Panjang string  `json:"panjang"`
	Lebar   string  `json:"lebar"`
	Hitung  float64 `json:"hitung"`
}

type Lingkara struct {
	JariJari string `json:"jariJari"`
	Hitung   string `json:"hitung"`
}

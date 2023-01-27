package models

type LabaRugi struct {
	Penjualan                float64
	ReturPenjualan           float64
	PotonganPenjualan        float64
	Pembelian                float64
	BebanAngkutPembelian     float64
	ReturPembelian           float64
	PotonganPembelian        float64
	PersediaanAwal           float64
	BarangTersedia           float64
	PersediaanAkhir          float64
	BebanOperasional         float64
	PendapatanNonOperasional float64
	BebanNonOperasional      float64
	LabaBersihSebelumPajak   float64
	PajakPenghasilan         int16
	LabaBersihSetelahPajak   float64
}

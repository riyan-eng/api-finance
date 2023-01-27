package entities

type Penjualan struct {
	Penjualan         float64
	ReturPenjualan    float64
	PotonganPenjualan float64
	PenjualanBersih   float64
}

type Pembelian struct {
	Pembelian            float64
	BebanAngkutPembelian float64
	ReturPembelian       float64
	PotonganPembelian    float64
	PembelianBersih      float64
}

type HPP struct {
	PersediaanAwal  float64
	Pembelian       Pembelian
	BarangTersedia  float64
	PersediaanAkhir float64
	HPP             float64
}

type LabaRugiOperasional struct {
	Penjualan           Penjualan
	HPP                 HPP
	BebanOperasional    float64
	LabaRugiOperasional float64
}

type LabaRugiNonOperasional struct {
	PendapatanNonOperasional float64
	BebanNonOperasional      float64
	LabaRugiNonOperasional   float64
}

type LabaRugi struct {
	LabaRugiOperasional    LabaRugiOperasional
	LabaRugiNonOperasional LabaRugiNonOperasional
	LabaBersihSebelumPajak float64
	PajakPenghasilan       float64
	LabaBersihSetelahPajak float64
}

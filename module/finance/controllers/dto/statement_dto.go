package dto

type Penjualan struct {
	Penjualan         float64 `json:"penjualan"`
	ReturPenjualan    float64 `json:"retur_penjualan"`
	PotonganPenjualan float64 `json:"potongan_penjualan"`
	PenjualanBersih   float64 `json:"penjualan_bersih"`
}

type Pembelian struct {
	Pembelian            float64 `json:"pembelian"`
	BebanAngkutPembelian float64 `json:"beban_angkut_pembelian"`
	ReturPembelian       float64 `json:"retur_pembelian"`
	PotonganPembelian    float64 `json:"potongan_pembelian"`
	PembelianBersih      float64 `json:"pembelian_bersih"`
}

type HPP struct {
	PersediaanAwal  float64   `json:"persediaan_awal"`
	Pembelian       Pembelian `json:"pembelian"`
	BarangTersedia  float64   `json:"barang_tersedia"`
	PersediaanAkhir float64   `json:"persediaan_akhir"`
	HPP             float64   `json:"hpp"`
}

type LabaRugiOperasional struct {
	Penjualan           Penjualan `json:"penjualan"`
	HPP                 HPP       `json:"hpp"`
	BebanOperasional    float64   `json:"beban_operasional"`
	LabaRugiOperasional float64   `json:"laba_rugi_operasional"`
}

type LabaRugiNonOperasional struct {
	PendapatanNonOperasional float64 `json:"pendapatan_non_operasional"`
	BebanNonOperasional      float64 `json:"beban_non_operasional"`
	LabaRugiNonOperasional   float64 `json:"laba_rugi_non_operasional"`
}

type LabaRugi struct {
	LabaRugiOperasional    LabaRugiOperasional    `json:"laba_rugi_operasional"`
	LabaRugiNonOperasional LabaRugiNonOperasional `json:"laba_rugi_non_operasional"`
	LabaBersihSebelumPajak float64                `json:"laba_bersih_sebelum_pajak"`
	PajakPenghasilan       int16                  `json:"pajak_penghasilan"`
	LabaBersihSetelahPajak float64                `json:"laba_bersih_setelah_pajak"`
}

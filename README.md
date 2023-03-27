# sanbercode-quiz3
API Routes
Berikut ini adalah titik akhir API yang tersedia:

Titik Akhir Bangun Datar
Segitiga Sama Sisi
Titik Akhir GET /bangun-datar/segitiga-sama-sisi

Endpoint ini mengambil informasi tentang Segitiga Sama Sisi.

Persegi
Endpoint: GET /bangun-datar/persegi

Endpoint ini mengambil informasi tentang Persegi.

Persegi Panjang
Titik akhir: GET /bangun-datar/persegi-panjang

Endpoint ini mengambil informasi tentang Persegi Panjang.

Lingkaran
Titik akhir: GET /bangun-datar/lingkaran

Endpoint ini mengambil informasi tentang Lingkaran.

Categories Endpoints
Dapatkan Semua Kategori
Endpoint: GET /kategori

Endpoint  ini mengambil semua kategori yang tersedia.

Insert Categories
Endpoint  POST /categories

Ruas akhir ini menyisipkan kategori baru.

Update Categories
Endpoint : PUT /kategori/:id

Ruas akhir ini memperbarui kategori berdasarkan ID.

Delete Categories
Endpoint : DELETE /categories/:id

Endpoint  ini menghapus kategori berdasarkan ID.

Dapatkan Buku Berdasarkan ID Kategori
Titik akhir: GET /categories/:id

Endpoint ini mengambil semua buku yang termasuk dalam sebuah kategori berdasarkan ID.

Books Endpoints
Get Books
Endpoint GET /books

Endpoint ini mengambil semua buku yang tersedia.

Endpoint POST /books
Endpoint ini menambahkan buku baru.

Update Books
Endpoint : PUT /books/:id

Endpoint  ini memperbarui buku berdasarkan ID.

Delete Books
Endpoint DELETE /buku/:id

Endpoint ini menghapus buku berdasarkan ID.

Authentication Middleware
Gunakan Username dan Password yang diberikan untuk proses POST,PUT,DELETE
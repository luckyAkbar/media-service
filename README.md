# IMAGE SERVICE

Sebuah tempat dimana kita bisa menyimpan (hosting) gambar secara gratis

## Apa saja yang dapat dilakukan?

Beberapa hal yang bisa anda lakukan di layanan ini adalah

## 1. Menyimpan gambar
Layanan ini tersedia gratis bagi semua orang, namun gambar yang akan diupload akan dibatasi besarannya oleh server. <br>
Untuk menyimpan gambar Anda, dapat dilakukan dengan ketentuan request berikut:<br>

- **POST image.service.luckyakbar.tech**
- Ektensi gambar yang diperbolehkan: jpg, jpeg, png
- Payload: <br>
  - Content-Type: *multipart/form-data*
  - Body: image -> berisi gambar yang akan anda upload
- Respon Server Type: *application/json*
- Respon Server Body: *delete_key, update_key, image_key*

Catatan: <br>
- *delete_key* dari server bisa digunakan apabila anda ingin menghapus gambar yang telah anda upload. Hal ini dilakukan<br>
untuk mencegah adanya orang yang ingin menghapus gambar tanpa sepengetahuan anda.
- *update_key* dapat anda gunakan untuk mengubah gambar yang telah anda simpan.
- *image_key* adalah nama unik dari gambar anda. Anda memerlukan ini untuk melihat gambar yang telah anda simpan

## 2. Melihat gambar
Layanan ini digunakan untuk melihat gambar yang telah di upload di layanan ini. <br>
Untuk melihat gambar yang telah diupload, dapat dilakukan dengan cara berikut:<br>

- **GET image.service.luckyakbar.tech/*image_key***
- Response Server: menyesuaikan ekstensi & mimetype gambar

Catatan: <br>
- Apabila *image_key* yang anda masukan adalah tidak valid atau tidak dimiliki oleh gambar manapun, server akan meresponse *404 Not Found*

## 3. Mengubah gambar
Layanan ini digunakan apabila anda ingin mengubah gambar yang anda simpan. Misal anda menyimpan gambar peliharaan kesayangan anda, dan suatu saat <br>
anda mendapatkan peliharaan baru, maka anda bisa mengubah gambar peliharaan anda. <br>
Untuk melakukan hal ini, anda bisa melakukan request seperti berikut:<br>

- **PATCH image.service.luckyakbar.tech/*image_key***
- Ektensi gambar yang diperbolehkan: jpg, jpeg, png
- Payload: <br>
  - Content-Type: *multipart/form-data*
  - Body: image -> berisi gambar yang akan anda ganti
- Respon Server Type: *application/json*
- Respon Server Body: *delete_key, update_key, image_key*

Catatan: <br>
*image_key* dari gambar lama anda tidak akan berubah ketika ingin mengakses gambar baru anda. Namun *delete_key* dan *update_key* akan berubah


## 4. Menghapus gambar
Layanan ini bisa anda gunakan ketika anda ingin menghapus gambar yang telah anda upload. <br>
Caranya adalah dengan melakukan request seperti berikut: <br>

- **DELETE image.service.luckyakbar.tech/*image_key***
- Payload: <br>
  - Content-Type: *application/json*
  - Body: *delete_key*
- Response Server Type: *application/json*
- Response Server Body: *image_key*

Catatan: <br>
Ketika anda menghapus gambar yang telah anda simpan, maka anda tidak akan bisa mengembalikannya lagi.
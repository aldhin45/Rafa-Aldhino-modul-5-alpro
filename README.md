SOAL1

Fungsi fibonacci: Fungsi ini menerima satu parameter n dan mengembalikan nilai Fibonacci ke-n. Jika n bernilai 0 atau 1, fungsi langsung mengembalikan nilai n karena suku pertama dan kedua dalam deret Fibonacci adalah 0 dan 1. Untuk nilai n yang lebih besar dari 1, fungsi ini memanggil dirinya sendiri dua kali dengan argumen n-1 dan n-2, dan menjumlahkan hasilnya, mengikuti aturan rekursif dari deret Fibonacci.
Fungsi main: Di dalam fungsi utama (main), variabel n diatur bernilai 10, yang berarti kita akan menghitung deret Fibonacci hingga suku ke-10. Program kemudian mencetak pesan untuk menunjukkan batas deret yang akan ditampilkan.
SOAL2

Fungsi printBintang: Fungsi ini bertugas mencetak sejumlah bintang dalam satu baris. Fungsi ini menerima parameter n yang menunjukkan jumlah bintang yang harus dicetak pada baris tersebut. Jika n > 0, maka fungsi akan mencetak satu bintang (*) dan memanggil dirinya sendiri (printStars) dengan nilai n-1. Proses ini akan berlanjut hingga n mencapai 0, yang menandakan bahwa jumlah bintang yang diperlukan untuk baris tersebut sudah tercetak.
Fungsi printPattern: Fungsi ini bertugas mencetak pola bintang per baris. Fungsi menerima dua parameter: n: Jumlah total baris yang diinginkan. current: Baris yang sedang dicetak saat ini.
SOAL3

Fungsi faktor: Fungsi ini bertanggung jawab untuk mencari dan mencetak semua faktor dari bilangan N. Fungsi ini memiliki dua parameter: n: Bilangan yang akan dicari faktornya. i: Angka yang digunakan untuk memeriksa apakah i adalah faktor dari n
Fungsi main: Fungsi utama ini menjalankan : Pertama, program meminta input dari pengguna untuk memasukkan sebuah bilangan positif n. Setelah itu, program mencetak pesan "Faktor dari n:", dan memanggil fungsi findFactors dengan n sebagai bilangan yang akan dicari faktornya dan 1 sebagai nilai awal dari i.
SOAL4 -Fungsi printSequence: Fungsi ini digunakan untuk mencetak urutan angka dari N ke 1 dan kembali lagi ke N. Parameter n: Merupakan batas nilai awal yang dimasukkan oleh pengguna (contoh, N). Parameter current: Merupakan nilai yang sedang dicetak dalam urutan rekursif.

Fungsi main: Fungsi utama menjalankan program sebagai berikut: Meminta input dari pengguna untuk menentukan nilai N. Memanggil fungsi printSequence dengan nilai N sebagai batas akhir, dan current dimulai dengan N. Setelah seluruh urutan dicetak, fungsi fmt.Println() digunakan untuk membuat baris baru.
SOAL5

Fungsi cetakbilangan: Fungsi ini bertanggung jawab untuk mencetak bilangan ganjil dari 1 hingga N. Parameter n: Merupakan batas atas yang dimasukkan oleh pengguna. Fungsi akan mencetak bilangan ganjil dari 1 sampai N. Parameter current: Merupakan angka yang sedang diperiksa dan dicetak jika ganjil.
Fungsi main: Fungsi utama menjalankan program dengan urutan berikut: Meminta input dari pengguna untuk memasukkan nilai N. Memanggil fungsi printOdd dengan N sebagai batas atas dan memulai current dari 1.

SOAL6 -Fungsi nilai: Fungsi ini bertanggung jawab untuk menghitung hasil dari x pangkat y secara rekursif. Parameter x: Bilangan pokok yang akan dipangkatkan. Parameter y: Eksponen yang menentukan pangkat dari x

Fungsi main: Fungsi utama ini meminta input dari pengguna dan menampilkan hasil perhitungan pangkat. Program meminta pengguna memasukkan dua bilangan, x (bilangan pokok) dan y (eksponen). Program kemudian memanggil fungsi power(x, y) dan menyimpan hasilnya di variabel result.


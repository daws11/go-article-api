-- migrations/000002_seed_posts_table.up.sql

-- File ini berisi data awal (seed) untuk tabel 'posts'.
-- Ini digunakan untuk mengisi database dengan data dummy saat migrasi.

-- Artikel 1
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-1 (Ini judul yang panjangnya lebih dari 20 karakter)',
    'Ini adalah isi konten dummy untuk artikel ke-1. Konten ini harus cukup panjang, setidaknya 200 karakter agar valid saat diuji nanti. Kita bisa menambahkan teks berulang atau lorep ipsum singkat untuk memenuhi persyaratan panjang minimum ini. Testing data seeding penting untuk memastikan aplikasi berjalan sesuai harapan dengan data awal. Ini adalah paragraf tambahan untuk memastikan panjangnya lebih dari 200 karakter dengan mudah dan cepat saat pengujian.',
    'Teknologi',
    'Publish'
);

-- Artikel 2
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-2 (Contoh judul lain yang juga valid panjangnya)',
    'Konten dummy artikel kedua juga perlu memenuhi syarat panjang 200 karakter. Menambahkan data dummy seperti ini membantu dalam pengembangan frontend dan backend untuk simulasi data nyata. Kita bisa variasikan status dan kategori untuk setiap entri agar data lebih beragam dan mencakup semua kemungkinan kasus penggunaan yang ada di aplikasi artikel ini. Paragraf ekstra ditambahkan di sini untuk mencapai panjang minimum.',
    'Berita',
    'Draft'
);

-- Artikel 3
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-3 (Pastikan setiap judul unik dan panjang)',
    'Artikel ketiga dengan status Thrash sebagai contoh. Isi kontennya pun dibuat sepanjang mungkin, minimal 200 karakter. Proses seeding ini bisa memakan waktu jika dilakukan manual, oleh karena itu otomatisasi via migrasi sangat membantu efisiensi developer dalam menyiapkan environment development atau testing. Jangan lupa untuk terus menambahkan teks hingga syarat karakter minimum terpenuhi.',
    'Gaya Hidup',
    'Thrash'
);

-- Artikel 4
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-4 (Variasi kategori dan status sangat penting)',
    'Konten artikel keempat dengan kategori Olahraga dan status Publish. Mengisi database dengan data awal yang representatif mempercepat proses debugging dan validasi fitur, terutama fitur yang berkaitan dengan penampilan data seperti tabel, list, dan pagination di sisi frontend maupun logika bisnis di backend. Pastikan konten ini juga memenuhi syarat 200 karakter.',
    'Olahraga',
    'Publish'
);

-- Artikel 5
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-5 (Lebih banyak data untuk pengujian pagination)',
    'Artikel kelima ini melanjutkan proses seeding. Kontennya harus lebih dari 200 karakter. Pengujian fitur seperti pagination atau pencarian memerlukan sejumlah data yang cukup. Kategori artikel ini adalah Hiburan dan statusnya Draft. Variasi ini penting untuk menguji filter dan tampilan yang berbeda di aplikasi. Jangan lupa memastikan panjang konten selalu memadai.',
    'Hiburan',
    'Draft'
);

-- Artikel 6
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-6 (Memastikan konsistensi data dummy)',
    'Ini adalah artikel keenam dalam rangkaian data dummy. Kategori diatur ke Bisnis dan statusnya Thrash. Konten ini juga harus panjang, minimal 200 karakter. Konsistensi dalam format data dummy membantu developer lain memahami struktur data yang diharapkan dan mempercepat adaptasi mereka terhadap proyek yang sedang berjalan. Teks tambahan untuk mencapai batas karakter.',
    'Bisnis',
    'Thrash'
);

-- Artikel 7
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-7 (Contoh lain untuk kategori Teknologi)',
    'Artikel ketujuh kembali menggunakan kategori Teknologi, namun dengan status Publish. Konten ini, seperti yang lainnya, dibuat sepanjang mungkin (minimal 200 karakter) untuk simulasi data nyata. Data seeding yang baik mencakup berbagai kombinasi kategori dan status untuk pengujian yang komprehensif. Teks ini ditambahkan untuk memenuhi syarat panjang minimum.',
    'Teknologi',
    'Publish'
);

-- Artikel 8
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-8 (Menggunakan kategori Berita lagi)',
    'Konten dummy untuk artikel kedelapan. Kategori Berita digunakan lagi di sini, dengan status Draft. Pastikan konten ini memiliki panjang minimal 200 karakter. Simulasi data yang beragam membantu mengidentifikasi potensi bug atau masalah tampilan pada kondisi data yang berbeda-beda. Pengujian menyeluruh adalah kunci kualitas aplikasi. Paragraf ini untuk memperpanjang teks.',
    'Berita',
    'Draft'
);

-- Artikel 9
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-9 (Status Thrash untuk Gaya Hidup)',
    'Artikel kesembilan dengan kategori Gaya Hidup dan status Thrash. Konten harus tetap memenuhi syarat minimal 200 karakter. Memiliki data dengan status Thrash penting untuk menguji fitur pemulihan data atau penghapusan permanen jika ada. Setiap status merepresentasikan state yang berbeda dalam lifecycle sebuah artikel. Teks tambahan disertakan di sini.',
    'Gaya Hidup',
    'Thrash'
);

-- Artikel 10
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-10 (Satu dekade artikel dummy!)',
    'Mencapai artikel kesepuluh dalam proses seeding. Kategori Olahraga dengan status Publish. Panjang konten minimal 200 karakter tetap menjadi syarat utama. Data dummy ini sangat berguna untuk demonstrasi aplikasi kepada klien atau stakeholder, menunjukkan bagaimana aplikasi akan terlihat dan berfungsi dengan data sungguhan. Pastikan panjang konten cukup.',
    'Olahraga',
    'Publish'
);

-- Artikel 11
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-11 (Masih berlanjut dengan variasi)',
    'Artikel kesebelas, kategori Hiburan dan status Draft. Seperti biasa, konten harus lebih dari 200 karakter. Variasi data terus dilakukan untuk mencakup lebih banyak skenario pengujian. Pengembangan fitur pencarian dan filter sangat bergantung pada keberagaman data awal seperti ini. Teks ini ditambahkan untuk mencapai panjang minimum yang diperlukan.',
    'Hiburan',
    'Draft'
);

-- Artikel 12
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-12 (Kategori Bisnis dengan status Thrash)',
    'Ini artikel kedua belas. Kategori Bisnis dan status Thrash. Konten harus panjang, setidaknya 200 karakter. Menguji bagaimana sistem menangani artikel yang dihapus (Thrash) dari berbagai kategori adalah bagian penting dari quality assurance. Data dummy ini memfasilitasi pengujian tersebut. Teks tambahan untuk memenuhi persyaratan panjang.',
    'Bisnis',
    'Thrash'
);

-- Artikel 13
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-13 (Teknologi lagi dengan status Publish)',
    'Artikel ketiga belas. Kategori Teknologi dengan status Publish. Konten harus minimal 200 karakter. Semakin banyak data yang mirip dengan data produksi, semakin baik proses pengujian dan pengembangan. Ini membantu menemukan masalah performa atau bug logika lebih awal. Teks ini memastikan konten cukup panjang.',
    'Teknologi',
    'Publish'
);

-- Artikel 14
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-14 (Berita terbaru dalam format Draft)',
    'Artikel keempat belas dengan kategori Berita dan status Draft. Konten minimal 200 karakter. Status Draft penting untuk menguji alur kerja editorial, di mana artikel ditulis tetapi belum dipublikasikan. Data dummy harus mencerminkan alur kerja nyata ini. Teks tambahan untuk mencapai panjang minimum.',
    'Berita',
    'Draft'
);

-- Artikel 15
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-15 (Setengah jalan! Gaya Hidup Thrash)',
    'Artikel kelima belas, menandai setengah jalan proses seeding 30 artikel. Kategori Gaya Hidup dan status Thrash. Konten tetap harus minimal 200 karakter. Memiliki cukup banyak data di setiap status dan kategori membantu menguji agregasi data dan laporan statistik jika fitur tersebut ada. Teks tambahan disertakan.',
    'Gaya Hidup',
    'Thrash'
);

-- Artikel 16
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-16 (Olahraga terbit lagi!)',
    'Artikel keenam belas. Kategori Olahraga dengan status Publish. Panjang konten minimal 200 karakter. Data yang dipublikasikan adalah yang paling sering dilihat pengguna, jadi penting untuk memiliki banyak contoh untuk menguji tampilan dan performa halaman daftar artikel. Teks tambahan untuk memenuhi syarat panjang.',
    'Olahraga',
    'Publish'
);

-- Artikel 17
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-17 (Hiburan dalam status Draft)',
    'Artikel ketujuh belas. Kategori Hiburan, status Draft. Konten minimal 200 karakter. Proses seeding yang baik juga mempertimbangkan distribusi data yang realistis, meskipun dalam skala kecil. Ini membantu mengantisipasi bagaimana aplikasi akan berperilaku dengan data nyata. Teks ini untuk memastikan panjang konten.',
    'Hiburan',
    'Draft'
);

-- Artikel 18
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-18 (Artikel Bisnis masuk keranjang sampah)',
    'Artikel kedelapan belas. Kategori Bisnis dengan status Thrash. Konten minimal 200 karakter. Pengujian fitur empty trash atau pemulihan item memerlukan data dengan status Thrash. Data dummy ini memungkinkan pengujian fitur tersebut secara efektif. Teks tambahan untuk mencapai batas karakter.',
    'Bisnis',
    'Thrash'
);

-- Artikel 19
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-19 (Teknologi yang sudah dipublikasi)',
    'Artikel kesembilan belas. Kategori Teknologi dengan status Publish. Konten minimal 200 karakter. Memastikan data publish dari berbagai kategori ditampilkan dengan benar adalah prioritas. Pengujian tampilan, sorting, dan filtering pada data publish sangat penting. Teks ini ditambahkan untuk memenuhi syarat panjang.',
    'Teknologi',
    'Publish'
);

-- Artikel 20
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-20 (Dua puluh artikel, kategori Berita Draft)',
    'Artikel kedua puluh. Kategori Berita dengan status Draft. Konten minimal 200 karakter. Jumlah data yang cukup banyak (seperti 20 atau 30) mulai berguna untuk menguji performa query database dan efisiensi pengambilan data, terutama dengan pagination. Teks tambahan untuk mencapai panjang minimum.',
    'Berita',
    'Draft'
);

-- Artikel 21
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-21 (Gaya Hidup masuk Thrash lagi)',
    'Artikel kedua puluh satu. Kategori Gaya Hidup dengan status Thrash. Konten minimal 200 karakter. Terus menambahkan data dengan status dan kategori yang bervariasi memastikan cakupan pengujian yang luas untuk berbagai fitur aplikasi. Teks tambahan disertakan di sini.',
    'Gaya Hidup',
    'Thrash'
);

-- Artikel 22
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-22 (Olahraga publish, persiapan akhir)',
    'Artikel kedua puluh dua. Kategori Olahraga dengan status Publish. Konten minimal 200 karakter. Semakin mendekati target 30 artikel, database pengembangan menjadi semakin representatif untuk pengujian fitur-fitur kompleks. Teks tambahan untuk memenuhi syarat panjang.',
    'Olahraga',
    'Publish'
);

-- Artikel 23
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-23 (Hiburan masih dalam Draft)',
    'Artikel kedua puluh tiga. Kategori Hiburan dengan status Draft. Konten minimal 200 karakter. Pengujian alur kerja dari Draft ke Publish bisa disimulasikan dengan data seperti ini. Developer bisa mengubah status data dummy ini untuk menguji transisi state. Teks ini untuk memastikan panjang konten.',
    'Hiburan',
    'Draft'
);

-- Artikel 24
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-24 (Bisnis dihapus sementara - Thrash)',
    'Artikel kedua puluh empat. Kategori Bisnis dengan status Thrash. Konten minimal 200 karakter. Memastikan fitur pengelolaan sampah (trash) berfungsi baik untuk semua kategori adalah penting, data ini membantu verifikasi tersebut. Teks tambahan untuk mencapai batas karakter.',
    'Bisnis',
    'Thrash'
);

-- Artikel 25
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-25 (Teknologi publish lagi, seperempat akhir!)',
    'Artikel kedua puluh lima. Kategori Teknologi dengan status Publish. Konten minimal 200 karakter. Tinggal beberapa artikel lagi. Data publish ini menambah jumlah data yang bisa dilihat oleh pengguna akhir simulasi. Teks ini ditambahkan untuk memenuhi syarat panjang.',
    'Teknologi',
    'Publish'
);

-- Artikel 26
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-26 (Berita Draft mendekati final)',
    'Artikel kedua puluh enam. Kategori Berita dengan status Draft. Konten minimal 200 karakter. Artikel draft ini bisa digunakan untuk menguji fitur preview sebelum publish atau kolaborasi antar penulis jika ada. Teks tambahan untuk mencapai panjang minimum.',
    'Berita',
    'Draft'
);

-- Artikel 27
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-27 (Gaya Hidup di Thrash, hampir selesai)',
    'Artikel kedua puluh tujuh. Kategori Gaya Hidup dengan status Thrash. Konten minimal 200 karakter. Data Thrash terakhir sebelum mencapai 30 artikel. Pengujian komprehensif mencakup semua state dan kategori. Teks tambahan disertakan di sini.',
    'Gaya Hidup',
    'Thrash'
);

-- Artikel 28
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-28 (Olahraga publish, tinggal dua lagi)',
    'Artikel kedua puluh delapan. Kategori Olahraga dengan status Publish. Konten minimal 200 karakter. Data publish ini penting untuk menguji tampilan akhir dan performa list artikel. Teks tambahan untuk memenuhi syarat panjang.',
    'Olahraga',
    'Publish'
);

-- Artikel 29
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-29 (Hiburan Draft, satu sebelum terakhir)',
    'Artikel kedua puluh sembilan. Kategori Hiburan dengan status Draft. Konten minimal 200 karakter. Artikel draft terakhir dalam batch seeding ini. Penting untuk menguji bagaimana artikel terakhir dalam daftar atau halaman ditampilkan. Teks ini untuk memastikan panjang konten.',
    'Hiburan',
    'Draft'
);

-- Artikel 30
INSERT INTO posts (title, content, category, status) VALUES
(
    'Judul Artikel Dummy ke-30 (Artikel terakhir dalam seeding data ini)',
    'Ini adalah konten untuk artikel dummy yang ke-30. Proses seeding hampir selesai. Konten ini, seperti yang lain, harus memiliki panjang minimal 200 karakter sesuai aturan validasi yang ditetapkan. Penggunaan data dummy yang konsisten dan valid memperlancar siklus pengembangan perangkat lunak secara keseluruhan. Teks tambahan untuk memastikan syarat panjang terpenuhi.',
    'Teknologi', -- Menggunakan Teknologi lagi sebagai contoh terakhir
    'Draft'      -- Status Draft untuk contoh terakhir
);

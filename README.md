# BILL CLI

BILL CLI adalah tool sederhana berbasis Go yang membantu developer melakukan proses Git seperti commit, pull, dan push dengan lebih cepat melalui satu perintah.

Project ini dibuat sebagai latihan membangun CLI tool dan memahami bagaimana Git automation bekerja.

---

## 📅 Development Log

### 06 March 2026

Hari pertama mengembangkan **BILL CLI**

Progress saat ini:

- ✅ Mendeteksi apakah folder memiliki `.git`
- ✅ Commit otomatis
- ✅ Pull otomatis sebelum push
- ✅ Push ke repository
- ✅ Validasi token GitHub
- ✅ Menyimpan token ke `config.json`

---

## 🚀 Fitur Saat Ini

### `bill push`

Perintah ini akan:

1. Mengecek apakah folder memiliki repository Git (`.git`)
2. Jika ada:
   - meminta pesan commit
   - menjalankan `git add .`
   - menjalankan `git commit`
   - menjalankan `git pull`
   - menjalankan `git push`
3. Jika belum ada repository:
   - membuat repository baru
   - meminta token GitHub
   - menyimpan token ke `config.json`

---

## 📦 Contoh Penggunaan

Menjalankan program:

```bash
go run cmd/bill/main.go
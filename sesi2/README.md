# Dasar Pemrograman Backend

## Live Coding - Student Management System

### Implementation technique

Siswa akan melaksanakan sesi live code di 15 menit terakhir dari sesi mentoring dan di awasi secara langsung oleh Mentor. Dengan penjelasan sebagai berikut:

- **Durasi**: 15 menit pengerjaan
- **Submit**: Maximum 10 menit setelah sesi mentoring menggunakan `grader-cli`
- **Obligation**: Wajib melakukan _share screen_ di breakout room yang akan dibuatkan oleh Mentor pada saat mengerjakan Live Coding.

### Description

Proyek ini adalah sebuah program _command-line interface (CLI)_ sederhana yang ditulis dalam bahasa Go, yang memungkinkan kita untuk mengelola data mahasiswa, seperti menambahkan mahasiswa baru, menghapus mahasiswa, menampilkan data mahasiswa, dan mencari mahasiswa berdasarkan nilai.

Program ini memiliki 4 fungsi utama:

- `ViewStudents`: untuk menampilkan data mahasiswa
- `AddStudent`: untuk menambahkan mahasiswa baru ke dalam sistem
- `RemoveStudent`: untuk menghapus mahasiswa dari sistem
- `HighestScore`: untuk mencari dan menampilkan mahasiswa dengan nilai tertinggi

Program ini menggunakan `map` untuk menyimpan data mahasiswa, dengan key berupa nama mahasiswa dan value berupa slice yang berisi alamat, nomor telepon, dan nilai mahasiswa.

### Constraints

Program ini dibagi menjadi 4 bagian:

- **Main**: mengontrol keseluruhan aliran program dan memanggil fungsi lainnya
- **ViewStudents**: menampilkan data mahasiswa dalam format tabel
- **AddStudent**: menambahkan mahasiswa baru ke dalam sistem
- **RemoveStudent**: menghapus mahasiswa dari sistem
- **HighestScore**: mencari mahasiswa dengan nilai tertinggi dan menampilkannya

Program ini menggunakan switch statement untuk memproses input pengguna dan memanggil fungsi yang sesuai untuk mengelola data mahasiswa. Jika pengguna memasukkan input yang tidak valid, program akan menampilkan pesan kesalahan.

Berikut adalah penjelasan dari fungsi-fungsi yang harus diimplementasi:

- Fungsi **RemoveStudent** menerima parameter `*map[string][]interface{}` dan mengembalikan sebuah fungsi yang menerima parameter string, yang akan menghapus data mahasiswa dengan nama yang diberikan dari `map`.

  ```go
  func RemoveStudent(students *map[string][]interface{}) func(string) {
    // TODO: answer here
  }
  ```

- Fungsi **HighestScore** menerima parameter `*map[string][]interface{}` dan mengembalikan nama mahasiswa dengan nilai tertinggi dan nilai tersebut. Jika tidak ada mahasiswa, fungsi ini mengembalikan string kosong dan nilai **0**.

  ```go
  func HighestScore(students *map[string][]interface{}) (string, int) {
    // TODO: answer here
  }
  ```

### Test Case Examples

**Input/Output**:

```bash
> Enter command (add, remove, high-score, view): add
> Enter name: John
> Enter address: Bogor
> Enter phone: 081234567890
> Enter score: 90
> Enter command (add, remove, high-score, view): add
> Enter name: Jane
> Enter address: Cianjur
> Enter phone: 081234567891
> Enter score: 85
> Enter command (add, remove, high-score, view): view
> Student data:
> Name    Address         Phone           Score
> John    Bogor           081234567890    90
> Jane    Cianjur         081234567891    85

> Enter command (add, remove, high-score, view): high-score
> High Score:  John 90
> Enter command (add, remove, high-score, view): remove
> Enter name: John
> Enter command (add, remove, high-score, view): view
> Student data:
> Name    Address         Phone           Score
> Jane    Cianjur         081234567891    85

> Enter command (add, remove, high-score, view): invalid-command
> Invalid command. Available commands: add, remove, high-score, view
> Enter command (add, remove, high-score, view):
```

**Explanation**:

> Test case ini menunjukkan fungsionalitas dasar dari program. Pengguna menambahkan siswa baru bernama John dengan alamat, nomor telepon, dan skor. Kita kemudian kita mencari siswa dengan nilai tertinggi (Mike dengan nilai 90). Kemudian kita melihat data siswa, menghapus John dari sistem, dan melihat data siswa yang diperbarui. Terakhir, pengguna memasukkan perintah yang tidak valid dan menerima pesan error.

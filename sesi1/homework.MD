
1. fork repo ini ke akun github masing2
2. di dalam folder `sesi1`, copas folder `example` beserta isinya. folder hasil copian diberi nama dengan format namamentee
3. kerjakan tugasnya di dalam file `main.go` pada folder yang sudah dicopy dan di-rename.
4. tugasnya adalah buat fungsi untuk mencari tahu siapa saja siswa yang nilainya sudah memenuhi syarat lulus dan yang belum
    - nama fungsinya GroupingStundentByPoint
    - inputnya berupa map dengan: key adalah nama siswa,  dan valuenya adalah nilai siswa. misal: `{"ikhsan": 30, "sans": 70, "ikh": 80}`
    - return value dari fungsi tersebut berupa map yang berisi dua item, dengan: key  `lulus` dan `tidak_lulus` saja, value masing2 item tersebut adalah slice yang berisi nama2 siswa yg memenuhi syarat masing2 group/key. misal: `{"lulus": {"namasiswa1", "namasiswa2"}, "tidak_lulus": {"namasiswa4"}}`
    - untuk syarat yang dimaksud bisa mengacu pada flowchart di bawah
5. jika sudah selesai, buat PR ke branch `master` pada repo sumber, paling lambat 3 April 2023, pukul 23.00 WIB.
6. Skor maksimal dalam penilaian tugas ini adalah 10 poin keaktifan.
7. Boleh dikerjakan bersama dengan teman2 kelompok dalam satu asmen yang sama

```mermaid
flowchart TD
    A([Start]) --> B[Create empty slice for lulus and tidak_lulus]
    B --> C[Loop over students and their scores from the input]
    C --> D{Is student score > 60?}
    D -- Yes --> E[Add student to lulus slice] --> H{Are there more students?}
    D -- No --> F[Add student to tidak_lulus slice] --> H
    H -- Yes --> I[Move to next student]
    I --> C
    H -- No --> J[create map for the return value]
    J --> K[Return result] --> END([End])
```
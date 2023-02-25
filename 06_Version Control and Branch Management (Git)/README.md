# (6) Version Control and Branch Management (Git)
## Perintah dalam git
### 1. git init
Perintah ini digunakan untuk membuat repositori baru. Caranya:
```git
git init
```

### 2. git clone
Perintah git clone digunakan untuk checkout repositori. Jia repositori berada di remove server, gunakan:
```git
git clone alex@93.188.160.58:/path/to/repository
```

### 3. git commit
Perintah git commit digunakan untuk melakukan commit pada perubahan ke head. Ingat bahwa perubahan apapun yang di-commit tidak akan langsung ke remote repository. Gunakan:
```git
git commit –m “Isi dengan keterangan untuk commit”
```

### 4. git push
git push adalah perintah git dasar lainnya. Push akan mengirimkan perubahan ke master branch dari remote repository yang berhubungan dengan direktori kerja Anda. Misalnya:
```git
git push origin master
```

### 5. git chekout
Perintah git checkout bisa digunakan untuk membuat branch atau untuk berpindah diantaranya. Misalnya, perintah berikut ini akan membuat branch baru dan berpindah ke dalamnya:
```git
command git checkout -b <nama-branch>
```

### 6. git branch
Perintah git branch bisa digunakan untuk me-list, membuat atau menghapus branch. Untuk menampilkan semua branch yang ada di repository, gunakan:
```git
git branch
```
Untuk menghapus branch:
```git
git branch -d <branch-name>
```

### 7. git pull
Untuk menggabungkan semua perubahan yang ada di remote repository ke direktori lokal, gunakan perintah pull ;
```git
git pull
```

### 8. git fetch
Perintah ini digunakan untuk menampilkan semua object dari remote repository yang tidak berada di direktori kerja lokal. Contohnya:
```git
git fetch origin
```

## Git Flow 
Alur kerja Git Flow adalah alur kerja dengan cabang fitur. 

Perbedaan dari alur kerja lainnya ada pada cara pengembang untuk membuat cabang.

Dari alur ini, developer tidak diperbolehkan untuk membuat cabang langsung dari cabang master, hal ini dikarenakan untuk menghapus code review yang menyinggung cabang master. 

## Merge Conflict Pada GitHub
Merge conflict biasanya terjadi jika ada dua orang yang mengedit file yang sama.
Kenapa bisa begitu, "kan mereka sudah punya cabang masing-masing?"
Bisa jadi, di cabang yang mereka kerjakan ada file yang sama dengan cabang lain. Kemudian, saat digabungkan terjadi conflict.
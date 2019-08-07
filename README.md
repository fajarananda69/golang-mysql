# golang-mysql

## Install Driver
Unduh driver mysql menggunakan go get.
```bash
go get -u "github.com/go-sql-driver/mysql"
```
## Connection
```golang
 func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1)/DBName")
	if err != nil {
		fmt.Println("gagal")
		return nil, err
	}
	fmt.Println("success login")
	return db, nil
}
  ```
Skema connection string untuk driver mysql yang kita gunakan cukup unik, ``root:pass@tcp(127.0.0.1:3306)/DBName``. Dibawah ini merupakan skema connection string yang bisa digunakan pada driver Go MySQL Driver. Jika anda menggunakan driver mysql lain, skema koneksinya bisa saja berbeda tergantung driver yang digunakan.
```
user:password@tcp(host:port)/dbname
user@tcp(host:port)/dbname
```
Di bawah ini adalah penjelasan mengenai connection string yang digunakan pada fungsi `connect()`.
```
root@tcp(127.0.0.1:3306)/db_belajar_golang
// user     => root
// password =>
// host     => 127.0.0.1 atau localhost
// port     => 3306
// dbname   => DBName
```

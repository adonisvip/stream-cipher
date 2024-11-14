## Build code

go build

## Test Encrypt and Decrypt

```
stream-cipher --help
Usage: stream-cipher <input_file> <output_file> <key> <mode>

Options:
  <input_file>    Tên file đầu vào chứa văn bản cần mã hóa hoặc giải mã
  <output_file>   Tên file đầu ra để lưu trữ kết quả sau khi mã hóa hoặc giải mã
  <key>           Khóa dùng để mã hóa và giải mã (nhập dưới dạng chuỗi)
  <mode>          Chế độ hoạt động: 'encrypt' để mã hóa hoặc 'decrypt' để giải mã

Example:
  stream-cipher plaintext.txt encrypted.txt mysecretkey encrypt
  stream-cipher encrypted.txt decrypted.txt mysecretkey decrypt
```
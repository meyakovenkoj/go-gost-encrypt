# Модуль для стохастического преобразования в Go

Основано на [библиотеке на языке C](https://github.com/meyakovenkoj/kuznechik_cuda)

Для работы модуля требуется установить библиотеку:
```bash
git clone https://github.com/meyakovenkoj/kuznechik_cuda gostcrypt
cd gostcrypt
mkdir build && cd _
cmake ..
make && make install
```

### 1 Ответ

Программа выведет: [77, 78, 79].

В созданном массиве элементы нумеруются следующим образом:

- значение:76; индекс:0
- значение:77; индекс:1
- значение:78; индекс:2
- значение:79; индекс:3
- значение:80; индекс:4

Вывод происходит среза элементов, начиная с индекса 1(включая), заканчивая 4(не включая).

Тем самым в поток попадают числа 76, 77, 78, 79. 


### 2 Ответ

Программа выведет: 2 \n 1.

В первом случае функция test возвращает значение 2, т.к. отложенная функция может считывать и присваивать возвращаемое значение. 

В функции же anotherTest defer работает предсказуемо. Она не влияет на переменную x, что не сказывается на ее значении. 


### 3 Ответ

Программа выведет: nil \n false. Это объясняется тем, что интерфейс error не равен nil. 

Интерфейс равен nil тогда и только тогда, когда тип и значение равны nil. 


Интерфейс представляет структуру из двух указателей, один ссылается на хранимое значение, второй же на описание самого интерфейса. 

Описание содержит информацию о типе и указатель на виртуальную таблицу. 

Пустой же интерфейс содержит лишь информацию о типе и указатель на хранимое значение. 

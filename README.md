# short-string
## Описание
В Cloud очень много серверов и, чтобы видеть более наглядно, какие модели серверов и в каком количестве есть в нашей инфраструктуре, мы заменяем набор моделей на сокращения.

## Задание
Напишите функцию, которая на вход получает строку вида `aaabbbccccc` и сворачивает ее до `a3b3c5`. Длина строки может быть не ограничена.  

**Важно!**

Если строка содержит повторяющиеся символы, их нужно сгруппировать в один объект и отсортировать. 
Например, строка `aaabbbcccccaaaaa` должна свернутся в `a8b3с5`, а строка `zzzzcccUUUuu` в `U3c3u2z4`
Также возможно использование русских букв: `ЯЯЯБББддд` свернется в `Б3Я3д3`
Вы можете использовать любые стандартные библиотеки Go.

### Sample Input:
aaabbbccccc
### Sample Output:
a3b3c5
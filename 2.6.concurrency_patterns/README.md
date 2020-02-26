.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]


---
class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Concurrency Patterns

### Юрочко Юрий


---
class: top white
background-image: url(tmp/sound.svg)
background-size: 130%
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **–** если есть проблемы cо звуком или с видео
]


---
# Цель занятия 
- познакомиться с некоторыми паттернами конкурентного программирования в Go
- посмотреть на примеры реализации в коде


---
# Конкурентный код
Глобально мы обеспечивавем безопаность за счет:
- примитивов синхронизации (e.g. sync.Mutex, etc)
- каналы
- confinement-техники


---
# Confinement-техники
Варианты:
- неизменяемые данные (идеально, но далеко не всегда возможно)
- ad hock
- lexical


---
# Ad hoc
По сути, неявная договоренность, что "я - читаю, а ты пишешь", поэтому мы не используем никакие средства синхронизации.
```
data := make([]int, 4)

loopData := func(handleData chan<- int) {
        defer close(handleData)
        for i := range data {
                handleData <- data[i]
        }
}

handleData := make(chan int)
go loopData(handleData)

for num := range handleData {
        fmt.Println(num)
}
```

---
# Lexical
Никакой договоренности нет, по сути, она неявно создана кодом.
```
chanOwner := func() <-chan int {
        results := make(chan int, 5)
        go func() {
                defer close(results)
                for i := 0; i <= 5; i++ {
                        results <- i
                }
        }()
        return results
}

consumer := func(results <-chan int) {
        for result := range results {
                fmt.Printf("Received: %d\n", result)
                }
        fmt.Println("Done receiving!")
}

results := chanOwner()
consumer(results)
```


---
# For-select цикл
Пример 1
```
for _, i := range []int{1, 2, 3, 4, 5} {
        select {
        case <-done:
                return
        case intStream <- i:
        }
}
```

Пример 2
```
for {
        select {
                case <- done:
                        return
                default:
        }
}
```


---
# Как предотвратить утечку горутин
Проблема:
```
doWork := func(strings <-chan string) <-chan interface{} {
        completed := make(chan interface{})
        go func() {
                defer fmt.Println("doWork exited.")
                defer close(completed)
                for s := range strings {
                        fmt.Println(s)
                }
        }()
        return completed
}

doWork(nil)

time.Sleep(time.Second * 5)
fmt.Println("Done.")
```


---
# Как предотвратить утечку горутин
Решение - явный индиктор того, что пора завершаться:
```
doWork := func(done <-chan interface{}, strings <-chan string)
        <-chan interface{} {
        terminated := make(chan interface{})
        go func() {
                defer fmt.Println("doWork exited.")
                defer close(terminated)
                for {
                        select {
                        case s := <-strings:
                                fmt.Println(s)
                        case <-done:
                                return
                        }
                }
        }()
        return terminated
}
...
```


---
# Or-channel
А что, если источников несколько?

Можно воспользоваться идеей выше и применить ее к нескольким каналам.


---
# Обработка ошибок
Главный вопрос - кто ответственнен за обработку ошибок?

Варианты:
- просто логировать (имеет право на жизнь)
- падать (плохой вариант, но встречается)
- возвращать ошибку туда, где больше контекста для обработки


---
# Обработка ошибок
Пример:
```
checkStatus := func(done <-chan interface{}, urls ...string)
        <-chan Result {
        results := make(chan Result)
        go func() {
                defer close(results)
                for _, url := range urls {
                        var result Result
                        resp, err := http.Get(url)
                        result = Result{Error: err, Response: resp}
                        select {
                        case <-done:
                                return
                        case results <- result:
                        }
                }
        }()
        return results
}
```


---
# Pipeline
Некая концепкия.

Суть - разбиваем работу, которую нужно выполнить, на некие этапы.

Каждый этап получает какие-то данные, обрабатывает, и отсылает их дальше.

Можно легко менять каждый этап, не задевая остальные.

https://blog.golang.org/pipelines

https://medium.com/statuscode/pipeline-patterns-in-go-a37bb3a7e61d


---
# Pipeline
Свойства, обычно применимые к этапу (stage)
- входные и выходные данные имеют один тип
- должна быть возможность передавать этап (например, фукнции в го - подходят)


---
# Простой пример (batch processing)
Stage 1
```
multiply := func(values []int, multiplier int) []int {
        multipliedValues := make([]int, len(values))
        for i, v := range values {
                multipliedValues[i] = v * multiplier
        }
        return multipliedValues
}
```
Stage 2
```
add := func(values []int, additive int) []int {
        addedValues := make([]int, len(values))
        for i, v := range values {
                addedValues[i] = v + additive
        }
        return addedValues
}
```


---
# Простой пример (batch processing)
Использование:
```
ints := []int{1, 2, 3, 4}
for _, v := range add(multiply(ints, 2), 1) {
        fmt.Println(v)
}
```


---
# Тот же пайплайн, но с горутинами
Генератор
```
generator := func(done <-chan interface{}, integers ...int) <-chan int {
        intStream := make(chan int)
        go func() {
                defer close(intStream)
                for _, i := range integers {
                        select {
                        case <-done:
                                return
                        case intStream <- i:
                        }
                }
        }()
        return intStream
}
```


---
# Тот же пайплайн, но с горутинами
Горутина с умножением
```
multiply := func(done <-chan interface{}, intStream <-chan int,
        multiplier int) <-chan int {
        multipliedStream := make(chan int)
        go func() {
                defer close(multipliedStream)
                for i := range intStream {
                        select {
                        case <-done:
                                return
                        case multipliedStream <- i*multiplier:
                        }
                }
        }()
        return multipliedStream
}
```


---
# Тот же пайплайн, но с горутинами
Горутина с добавлением
```
add := func(
done <-chan interface{},intStream <-chan int, additive int) <-chan int {
        addedStream := make(chan int)
        go func() {
                defer close(addedStream)
                for i := range intStream {
                        select {
                        case <-done:
                                return
                        case addedStream <- i+additive:
                        }
                }
        }()
        return addedStream
}
```


---
# Тот же пайплайн, но с горутинами
Использование:
```
done := make(chan interface{})
defer close(done)

intStream := generator(done, 1, 2, 3, 4)
pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

for v := range pipeline {
        fmt.Println(v)
}
```


---
# Полезные генераторы - Repeat
```
repeatFn := func(done <-chan interface{}, fn func() interface{})
        <-chan interface{} {
        valueStream := make(chan interface{})
        go func() {
                defer close(valueStream)
                for {
                        select {
                        case <-done:
                                return
                        case valueStream <- fn():
                        }
                }
        }()
        return valueStream
}
```


---
# Полезные генераторы - Take
```
take := func(done <-chan interface{}, valueStream <-chan interface{},
        num int) <-chan interface{} {
        takeStream := make(chan interface{})
        go func() {
                defer close(takeStream)
                for i := 0; i < num; i++ {
                        select {
                        case <-done:
                                return
                        case takeStream <- <-valueStream:
                        }
                }
        }()
        return takeStream
}
```


---
# Fan-Out
Процесс запуска нескольки горутин для обработки входных данных.


---
# Fan-In
Процесс слияния нескольких источников результов в один канал.


---
# Fan-Out & Fan-In
Смотрим на примере нахождения простых чисел.


---
# Выводы
- старайтесь писать максимально простой и понятный код
- пораждая горутину, всегда используйте done канал для управления
- не игнорируйте ошибки, старайтесь вернуть их туда, где больше контекста
- использование пайплайнов делает код более читаемым
- использование пайплайнов позволяет легко менять отдельные этапы


---
# Дополнительные материалы
https://blog.golang.org/pipelines

https://github.com/golang/go/wiki/LearnConcurrency

https://github.com/KeKe-Li/book/blob/master/Go/go-in-action.pdf

http://s1.phpcasts.org/Concurrency-in-Go_Tools-and-Techniques-for-Developers.pdf


---
# Итоги занятия 
- познакомились с некоторыми паттернами конкурентного программирования в Go
- посмотрели на примеры реализации в коде


---
## Вопросы?


---
class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!

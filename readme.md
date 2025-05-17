# Компилятор языка Do(мое видиние того как надо развивать Go)

## Как собирать

**NOTE:**
При первой сборке требуется подключение к Интернету

После клонирования репозитория необходимо выполнить кодогенерацию с помощью `antlr4`

```bash
    antlr4 -Dlanguage=Go Do.g4 -package parser -o parser
```

Для запуска

```bash
    go run . {put your source files here} [-o executable name(optional)] 
```

или

```bash
    go build .
    DoCompiler [put your source files here] [-o executable name(optional)] 
```

## Тестирование

Выполнить в корне репозитория

```bash
    go test
```

## Примеры

В папке `/doexamples` представлены примеры использования языка Do
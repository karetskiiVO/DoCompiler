# Компилятор языка Do(мое видиние того как надо развивать Go)

## Как собирать

После клонирования репозитория необходимо выполнить кодогенерацию с помощью `antlr4`

```bash
    antlr4 -Dlanguage=Go Do.g4 -package parser -o parser
```

Для запуска

```bash
    go run . [put your source files here]
```

или

```bash
    go build .
    DoCompiler [put your source files here]
```

## Основные концепты

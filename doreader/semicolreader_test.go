package doreader

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSemicolReader(t *testing.T) {
	code := `

act f () {
    /* объвим переменных */
    var a, b int

    a =
    1               // все верно
    a, b = 1            // ошибка: справа 1 значение - слева 2
    a = 1, 2            // ошибка: справа 2 значения - слева 1
    a, _ = 1, 2         // все верно - лишние значения подавляются псевдопеременной
    _, b = 1, 2         // тоже все верно
   

    a, b = 1, _         // ошибка: справа 1 значение - слева 2
    a, b = 1, _, 2      // а тут все корректно

    a = h()             // снова ошибка - функция возвращает больше значений
    a, _ = h()          // правильно - нельзя игнорировать значения просто так

    g()                 // тут нехватает аргументов в вызове функции

    a, b = f1(), f2(), f1(), f2() // а вот и то, ради чего все затевалось - после этой строки (a, b) = (1, 1)
}

act h () (int, int) {}

act g (a, b int) (int, int, int, int, int) {}

var buf int

act f1 () {
}

act f2 () int {
    var tmp int
    buf = 2
    return tmp
}

type A struct {
    a int
    b int
}
`

	want := `
act f () {
    
    var a, b int;

    a =
    1;               
    a, b = 1;            
    a = 1, 2;            
    a, _ = 1, 2;         
    _, b = 1, 2;         
   

    a, b = 1, _;         
    a, b = 1, _, 2;      

    a = h();             
    a, _ = h();          

    g();                 

    a, b = f1(), f2(), f1(), f2(); 
}

act h () (int, int) {}

act g (a, b int) (int, int, int, int, int) {}

var buf int;

act f1 () {
}

act f2 () int {
    var tmp int;
    buf = 2;
    return tmp;
}

type A struct {
    a int;
    b int;
}
`
	res, err := io.ReadAll(NewSemicolReader(NewCommentDeletor(strings.NewReader(code))))
	fmt.Println(string(res))
	require.NoError(t, err)
	require.Equal(t, want, string(res))
}

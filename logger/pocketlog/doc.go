/*
Package pocketlog expões uma API para logs.

Para usar, você precisa instanciar um logger com pocketlog.New selecionando um
nível, logs menos críticos que esse nível não serão registrados.

A reutilização do logger é de responsabilidade de quem usa o pacote.
*/
package pocketlog

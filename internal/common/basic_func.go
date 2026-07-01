package common

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IBasicRepository interface {
	CreateLog()
}

type BasicRepository[T any] struct {
	Pool *pgxpool.Pool
	self T
}

func (b *BasicRepository[T]) CreateLog() {
	fmt.Println("Save log")
}

func (b *BasicRepository[T]) GetAllEntries(table string) {
	stringConf := fmt.Sprintf("select * from %s", table)
	rows, err := b.Pool.Query(context.Background(), stringConf)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		entrie := b.self
		rows.Scan(&entrie)
	}
}

func multiply(a, b int) int {
	return a * b
}

func CreateTmpl(tmpl string, data interface{}) (string, error) {
	var buf bytes.Buffer
	t, err := template.New("new").
		Funcs(template.FuncMap{"multiply": multiply}).
		Parse(tmpl)
	if err != nil {
		return "", INerr
	}
	t.Execute(&buf, data)
	return buf.String(), nil
}

func Map[T any, R any, M any](items []T, meta M, fn func(T, M) R) []R {
	result := make([]R, 0, len(items))

	for _, item := range items {
		result = append(result, fn(item, meta))
	}

	return result
}

func JoinIds[T int | int16 | int32 | int64](ids []T) string {
	res := make([]string, len(ids))
	for i, v := range ids {
		res[i] = strconv.FormatInt(int64(v), 10)
	}
	return strings.Join(res, ", ")
}

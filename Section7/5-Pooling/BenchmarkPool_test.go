package main

import "testing"

func BenchmarkWithOutPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100_000_000; i++ {
			connection = &DbConnection{
				Host:     "localhost",
				DbName:   "Db",
				User:     "root",
				Password: "root",
			}
			connection.DbName = "Change"
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var connection *DbConnection
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100_000_000; i++ {
			connection = connectionPool.Get().(*DbConnection)
			connection.DbName = "Change"
			connectionPool.Put(connection)
		}
	}
}

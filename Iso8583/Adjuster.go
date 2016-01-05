package Iso8583

type Adjuster interface  {
	Get(string) string
	Set(string) string
}
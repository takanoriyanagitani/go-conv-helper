package main

import (
	"context"
	"log"
	"time"

	g9 "github.com/redis/go-redis/v9"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	ir "github.com/takanoriyanagitani/go-conv-helper/external/buffer/redis"
	gr "github.com/takanoriyanagitani/go-conv-helper/external/buffer/redis/go-redis"
)

const redisAddr string = "localhost:6379"
const redisKey string = "cafef00d-dead-beaf-face-864299792458"

var resKey string = ir.SimpleResponseKeyBuilderStaticDefault(redisKey)

const timeout time.Duration = time.Second

var client *g9.Client = g9.NewClient(&g9.Options{
	Addr:     redisAddr,
	Password: "",
	DB:       0,
})

var llen ir.LLen[string] = gr.LLenNew(client).ToLLen()
var lget ir.LengthGetter[string] = ir.LengthGetter[string](llen)

var lchk ir.LengthChecker = ir.LengthCheckerNewStatic(10)

var lpush gr.LPush[[]byte] = gr.LPushNew[[]byte](client)
var limited ir.LPush[string, []byte] = lpush.
	ToLPush().
	CreateLimited(lget)(lchk)

var brpop ir.BRPop[string] = gr.
	BRPopNew(client).
	ToBRPop()

var str2bytes ir.SerializeToBytes[string] = func(input string) ([]byte, error) {
	return []byte(input), nil
}

var str2str ir.ParseFromString[string] = func(val string) (parsed string, e error) {
	return val, nil
}

var builder ir.ConverterBuilder[string, string, string, []byte] = ir.
	ConverterBuilder[string, string, string, []byte]{}.
	WithSerializer(ir.Serialize[string, []byte](str2bytes)).
	WithLPush(limited).
	WithBRPop(brpop).
	WithParse(ir.Parse[string, string](str2str)).
	WithRequestKey(redisKey).
	WithResponseKey(resKey)
var conv ch.Converter[string, string] = builder.Build(timeout)

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds)
	var ctx context.Context = context.Background()
	const input string = "helo"
	log.Printf("converting...\n")
	var started time.Time = time.Now()
	converted, e := conv(ctx, input)
	var finished time.Time = time.Now()
	log.Printf("converted.\n")
	var diff time.Duration = finished.Sub(started)
	log.Printf("diff: %v\n", diff)
	if nil != e {
		log.Fatalf("Unable to convert: %v\n", e)
	}
	log.Printf("input:     %s\n", input)
	log.Printf("converted: %s\n", converted)
}

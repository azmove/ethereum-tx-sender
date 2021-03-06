package gas

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFetchAndGet(t *testing.T){
	price, err := fetchPrice(context.Background())

	assert.Nil(t, err)
	spew.Dump(price)

	normalPrice, urgentPrice := GetCurrentGasPrice()
	spew.Dump(normalPrice, urgentPrice)
}

func TestStartFetcher(t *testing.T) {
	ctx,cancel := context.WithTimeout(context.Background(), time.Second * 10)

	go StartFetcher(ctx)
	<- time.After(time.Second * 5)
	cancel()
	spew.Dump(GetCurrentGasPrice())
}

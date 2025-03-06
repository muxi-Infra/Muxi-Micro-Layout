package data

import (
	"context"
	"greeter/conf"
	log "greeter/pkg/logger"

	"greeter/internal/biz"
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Conf, logger log.Logger) *Data {
	return &Data{}
}

type greeterRepo struct {
	data *Data
	log  *log.ZapLogger
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, log *log.ZapLogger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, s string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}

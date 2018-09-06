package elastic

import (
	"context"
	"net/url"
	"regexp"
	"time"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/full-stack-demo/server/appconst"
	"github.com/tsaikd/full-stack-demo/server/slacklog"
	elastic "gopkg.in/olivere/elastic.v5"
)

// errors
var (
	ErrSkipElasticLog = errutil.NewFactory("skip elastic log")
	ErrParseURLFailed = errutil.NewFactory("parse elastic prefix url failed")
)

var esClient *Client

// Elastic return elastic client instance
func Elastic() *Client {
	if esClient == nil {
		return &Client{}
	}
	return esClient
}

// Config init elastic config
func Config(ctx context.Context, elasticPrefix string) (err error) {
	if esClient, err = newClient(ctx, elasticPrefix); err != nil {
		return
	}

	return
}

// Close release used resource
func Close() (err error) {
	if esClient != nil {
		err = esClient.close()
		esClient = nil
	}
	return
}

// Client of elastic connection
type Client struct {
	indexPrefix string
	client      *elastic.Client
	processor   *elastic.BulkProcessor
	indexChan   chan indexDataType
	closeChan   chan struct{}
}

func (t *Client) close() (err error) {
	if t.indexChan != nil {
		close(t.indexChan)
	}
	if t.processor != nil {
		err = t.processor.Close()
		t.processor = nil
	}
	if t.client != nil {
		t.client = nil
	}
	t.indexPrefix = ""
	if t.closeChan != nil {
		<-t.closeChan // wait for runIndexLoop finished
	}
	return
}

func (t *Client) runIndexLoop(ctx context.Context) {
	if t.processor == nil {
		slacklog.Trace(ErrSkipElasticLog.New(nil))
	}

	defer func() {
		t.closeChan <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-t.indexChan:
			if !ok {
				return
			}
			if t.processor != nil {
				req := elastic.NewBulkIndexRequest().
					Index(t.indexPrefix + data.IndexName).
					Type(data.IndexType).
					Id(data.ID).
					Doc(data.Doc)
				t.processor.Add(req)
			}
		}
	}
}

// Index insert document to elastic
func (t *Client) Index(idxname string, idxtype string, id string, doc interface{}) {
	if t.indexChan != nil {
		t.indexChan <- indexDataType{
			IndexName: idxname,
			IndexType: idxtype,
			ID:        id,
			Doc:       doc,
		}
	}
}

var (
	regPrefixTrim = regexp.MustCompile(`^/`)
	regSuffixDash = regexp.MustCompile(`-$`)
)

// newClient return a new elastic client
// elasticPrefix format {protocol}://{host}:{port}/{indexprefix}
// indexprefix is optional
// set elasticPrefix == "-" or empty to disable elastic
func newClient(ctx context.Context, elasticPrefix string) (client *Client, err error) {
	switch elasticPrefix {
	case "", "-":
		return &Client{}, nil
	}

	urlInfo, err := url.Parse(elasticPrefix)
	if err != nil {
		return nil, ErrParseURLFailed.New(err)
	}

	indexprefix := regPrefixTrim.ReplaceAllString(urlInfo.Path, "")
	if indexprefix != "" && !regSuffixDash.MatchString(indexprefix) {
		indexprefix += "-"
	}

	clientURLInfo, err := urlInfo.Parse("/")
	if err != nil {
		return nil, ErrParseURLFailed.New(err)
	}

	esClient, err := elastic.NewClient(
		elastic.SetURL(clientURLInfo.String()),
		elastic.SetSniff(false),
		elastic.SetErrorLog(slacklog.Logger()),
	)
	if err != nil {
		return
	}

	bulkActions := 100
	flushInterval := 30 * time.Second

	processor, err := esClient.BulkProcessor().
		Name(appconst.AppNameLower + "-elastic").
		BulkActions(bulkActions).
		BulkSize(5 << 20). // 5MB
		FlushInterval(flushInterval).
		Do(ctx)
	if err != nil {
		return
	}

	client = &Client{
		indexPrefix: indexprefix,
		client:      esClient,
		processor:   processor,
		indexChan:   make(chan indexDataType, 50000),
		closeChan:   make(chan struct{}),
	}
	go client.runIndexLoop(ctx)
	return
}

type indexDataType struct {
	IndexName string
	IndexType string
	ID        string
	Doc       interface{}
}

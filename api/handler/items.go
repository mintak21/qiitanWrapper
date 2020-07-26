package handler

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	apiModel "github.com/mintak21/qiitaWrapper/api/model"
	"github.com/mintak21/qiitaWrapper/api/util/client"
	genModel "github.com/mintak21/qiitaWrapper/gen/models"
	"github.com/mintak21/qiitaWrapper/gen/restapi/qiitawrapper/items"
	log "github.com/sirupsen/logrus"
)

const (
	perPage = 50
)

func init() {
	strfmt.MarshalFormat = strfmt.RFC3339Millis
}

// NewGetTagItemsHandler handles a request for getting tag items
func NewGetTagItemsHandler() items.GetTagItemsHandler {
	return &tagItemsHandler{
		client: client.NewQiitaClient(),
	}
}

type tagItemsHandler struct {
	client client.QiitaClient
}

// NewSyncTagItemsHandler handles a request for getting target day tag items
func NewSyncTagItemsHandler() items.SyncTagItemsHandler {
	return &syncTagItemsHandler{
		client: client.NewQiitaClient(),
	}
}

type syncTagItemsHandler struct {
	client client.QiitaClient
}

// Handle the get entry request
func (h *tagItemsHandler) Handle(params items.GetTagItemsParams) middleware.Responder {
	query := fmt.Sprintf("tag:%s", params.Tag)
	response, hasNext, err := sendGetItemRequest(h.client, int(*params.Page), query)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error(("failed to send request to Qiita API"))
		return items.NewGetTagItemsInternalServerError().WithPayload(&genModel.Error{Message: err.Error()})
	}
	return items.NewGetTagItemsOK().WithPayload(toModel(response, *params.Page, hasNext))
}

func (h *syncTagItemsHandler) Handle(params items.SyncTagItemsParams) middleware.Responder {
	var targetDate string
	if params.Date == nil {
		targetDate = time.Now().Format(strfmt.RFC3339FullDate)
	} else {
		targetDate = params.Date.String()
	}
	query := fmt.Sprintf("tag:%s created:<=%s created:>=%s", params.Tag, targetDate, targetDate)
	response, hasNext, err := sendGetItemRequest(h.client, int(*params.Page), query)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error(("failed to send request to Qiita API"))
		return items.NewSyncTagItemsInternalServerError().WithPayload(&genModel.Error{Message: err.Error()})
	}
	return items.NewSyncTagItemsOK().WithPayload(toModel(response, *params.Page, hasNext))
}

func sendGetItemRequest(cl client.QiitaClient, page int, query string) ([]*apiModel.QiitaItem, bool, error) {
	parameter := client.NewGetItemsParameter(page, perPage+1, query)
	qiitaItems, err := cl.GetItems(parameter)
	if err != nil {
		return nil, false, err
	}
	if perPage < len(qiitaItems) {
		return qiitaItems[0 : len(qiitaItems)-1], true, err
	}
	return qiitaItems, false, nil
}

func toModel(resItems []*apiModel.QiitaItem, page int64, hasNext bool) *genModel.Items {
	var items []*genModel.Item
	for _, resItem := range resItems {
		item := genModel.Item{
			Title: resItem.Title,
			Link:  resItem.URL,
			User: &genModel.User{
				Name:          resItem.User.Name,
				ThumbnailLink: resItem.User.ProfileImageURL,
			},
			Statistics: &genModel.Statistics{
				Lgtms: int64(resItem.LikesCount),
			},
			CreatedAt: strfmt.DateTime(resItem.CreatedAt),
		}
		items = append(items, &item)
	}
	return &genModel.Items{
		HasNext: hasNext,
		Page:    page,
		Items:   items,
	}
}

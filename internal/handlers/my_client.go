package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mdayat/fullstack2024-test/go/configs"
	"github.com/mdayat/fullstack2024-test/go/internal/dbutil"
	"github.com/mdayat/fullstack2024-test/go/internal/dtos"
	"github.com/mdayat/fullstack2024-test/go/internal/httputil"
	"github.com/mdayat/fullstack2024-test/go/repository"
)

type MyClientHandler interface {
	GetMyClients(res http.ResponseWriter, req *http.Request)
	CreateMyClient(res http.ResponseWriter, req *http.Request)
	UpdateMyClient(res http.ResponseWriter, req *http.Request)
	DeleteMyClient(res http.ResponseWriter, req *http.Request)
}

type myClient struct {
	configs configs.Configs
}

func NewMyClientHandler(configs configs.Configs) MyClientHandler {
	return &myClient{
		configs: configs,
	}
}

func (mc *myClient) GetMyClients(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("GetMyClients"))
}

func (mc *myClient) CreateMyClient(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var reqBody dtos.CreateMyClientRequest
	if err := httputil.DecodeAndValidate(req, mc.configs.Validate, &reqBody); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var address pgtype.Text
	if reqBody.Address != "" {
		address = pgtype.Text{String: reqBody.Address, Valid: true}
	}

	var phoneNumber pgtype.Text
	if reqBody.PhoneNumber != "" {
		phoneNumber = pgtype.Text{String: reqBody.PhoneNumber, Valid: true}
	}

	var city pgtype.Text
	if reqBody.City != "" {
		city = pgtype.Text{String: reqBody.City, Valid: true}
	}

	now := time.Now()
	retryableFunc := func(qtx *repository.Queries) (repository.MyClient, error) {
		myClient, err := qtx.InsertMyClient(ctx, repository.InsertMyClientParams{
			Name:         reqBody.Name,
			Slug:         reqBody.Slug,
			IsProject:    reqBody.IsProject,
			SelfCapture:  reqBody.SelfCapture,
			ClientPrefix: reqBody.ClientPrefix,
			ClientLogo:   reqBody.ClientLogo,
			Address:      address,
			PhoneNumber:  phoneNumber,
			City:         city,
			CreatedAt:    pgtype.Timestamp{Time: now, Valid: true},
		})

		if err != nil {
			return repository.MyClient{}, err
		}

		encodedMyClient, err := json.Marshal(myClient)
		if err != nil {
			return repository.MyClient{}, err
		}

		if err = mc.configs.Redis.Set(ctx, myClient.Slug, encodedMyClient, 0).Err(); err != nil {
			return repository.MyClient{}, err
		}

		return myClient, err
	}

	myClient, err := dbutil.RetryableTxWithData(ctx, mc.configs.Db.Conn, mc.configs.Db.Queries, retryableFunc)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody := dtos.MyClientResponse{
		Id:           myClient.ID,
		Name:         myClient.Name,
		Slug:         myClient.Slug,
		IsProject:    myClient.IsProject,
		SelfCapture:  myClient.SelfCapture,
		ClientPrefix: myClient.ClientPrefix,
		ClientLogo:   myClient.ClientLogo,
		Address:      myClient.Address.String,
		PhoneNumber:  myClient.PhoneNumber.String,
		City:         myClient.City.String,
		CreatedAt:    myClient.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:    myClient.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:    myClient.DeletedAt.Time.Format(time.RFC3339),
	}

	params := httputil.SendSuccessResponseParams{
		StatusCode: http.StatusCreated,
		ResBody:    resBody,
	}

	if err := httputil.SendSuccessResponse(res, params); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (mc *myClient) UpdateMyClient(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	myClientIdStr := chi.URLParam(req, "myClientId")
	myClientId, err := strconv.Atoi(myClientIdStr)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	var reqBody dtos.CreateMyClientRequest
	if err := httputil.DecodeAndValidate(req, mc.configs.Validate, &reqBody); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var address pgtype.Text
	if reqBody.Address != "" {
		address = pgtype.Text{String: reqBody.Address, Valid: true}
	}

	var phoneNumber pgtype.Text
	if reqBody.PhoneNumber != "" {
		phoneNumber = pgtype.Text{String: reqBody.PhoneNumber, Valid: true}
	}

	var city pgtype.Text
	if reqBody.City != "" {
		city = pgtype.Text{String: reqBody.City, Valid: true}
	}

	now := time.Now()
	retryableFunc := func(qtx *repository.Queries) (repository.MyClient, error) {
		oldMyClient, err := qtx.SelectMyClientById(ctx, int32(myClientId))
		if err != nil {
			return repository.MyClient{}, err
		}

		if err = mc.configs.Redis.Del(ctx, oldMyClient.Slug).Err(); err != nil {
			return repository.MyClient{}, err
		}

		newMyClient, err := qtx.UpdateMyClient(ctx, repository.UpdateMyClientParams{
			Name:         reqBody.Name,
			Slug:         reqBody.Slug,
			IsProject:    reqBody.IsProject,
			SelfCapture:  reqBody.SelfCapture,
			ClientPrefix: reqBody.ClientPrefix,
			ClientLogo:   reqBody.ClientLogo,
			Address:      address,
			PhoneNumber:  phoneNumber,
			City:         city,
			UpdatedAt:    pgtype.Timestamp{Time: now, Valid: true},
		})

		if err != nil {
			return repository.MyClient{}, err
		}

		encodedMyClient, err := json.Marshal(newMyClient)
		if err != nil {
			return repository.MyClient{}, err
		}

		if err = mc.configs.Redis.Set(ctx, newMyClient.Slug, encodedMyClient, 0).Err(); err != nil {
			return repository.MyClient{}, err
		}

		return newMyClient, err
	}

	newMyClient, err := dbutil.RetryableTxWithData(ctx, mc.configs.Db.Conn, mc.configs.Db.Queries, retryableFunc)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	resBody := dtos.MyClientResponse{
		Id:           newMyClient.ID,
		Name:         newMyClient.Name,
		Slug:         newMyClient.Slug,
		IsProject:    newMyClient.IsProject,
		SelfCapture:  newMyClient.SelfCapture,
		ClientPrefix: newMyClient.ClientPrefix,
		ClientLogo:   newMyClient.ClientLogo,
		Address:      newMyClient.Address.String,
		PhoneNumber:  newMyClient.PhoneNumber.String,
		City:         newMyClient.City.String,
		CreatedAt:    newMyClient.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:    newMyClient.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:    newMyClient.DeletedAt.Time.Format(time.RFC3339),
	}

	params := httputil.SendSuccessResponseParams{
		StatusCode: http.StatusCreated,
		ResBody:    resBody,
	}

	if err := httputil.SendSuccessResponse(res, params); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (mc *myClient) DeleteMyClient(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	myClientIdStr := chi.URLParam(req, "myClientId")
	myClientId, err := strconv.Atoi(myClientIdStr)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	now := time.Now()
	retryableFunc := func(qtx *repository.Queries) (int64, error) {
		myClient, err := qtx.SelectMyClientById(ctx, int32(myClientId))
		if err != nil {
			return 0, err
		}

		if err = mc.configs.Redis.Del(ctx, myClient.Slug).Err(); err != nil {
			return 0, err
		}

		return qtx.DeleteMyClient(ctx, repository.DeleteMyClientParams{
			ID:        myClient.ID,
			DeletedAt: pgtype.Timestamp{Time: now, Valid: true},
		})
	}

	affectedRows, err := dbutil.RetryableTxWithData(ctx, mc.configs.Db.Conn, mc.configs.Db.Queries, retryableFunc)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if affectedRows == 0 {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

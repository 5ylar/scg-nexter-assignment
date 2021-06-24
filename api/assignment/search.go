package assignment

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) multipleSearchFromDataset(ctx echo.Context) error {
	var req multipleSearchFromDatasetReq

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	result := s.search.SearchMultipleKeysFromStringDataset(req.Keys, req.Dataset, req.CaseSensitive)

	return ctx.JSON(http.StatusOK, result)
}

type multipleSearchFromDatasetReq struct {
	Keys          []string `json:"keys"`
	Dataset       []string `json:"dataset"`
	CaseSensitive bool     `json:"caseSensitive"`
}

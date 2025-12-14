package handlers

import (
	"stock-ingestor/internal/application/usecases"

	"strings"

	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	getStockTimeSeriesUseCase *usecases.GetStockTimeSeriesUseCase
	getCompanyOverviewUseCase *usecases.GetCompanyOverviewUseCase
}

func NewStockHandler(
	getStockTimeSeriesUseCase *usecases.GetStockTimeSeriesUseCase,
	getCompanyOverviewUseCase *usecases.GetCompanyOverviewUseCase,
) *StockHandler {
	return &StockHandler{
		getStockTimeSeriesUseCase: getStockTimeSeriesUseCase,
		getCompanyOverviewUseCase: getCompanyOverviewUseCase,
	}
}

func (h *StockHandler) GetStockTimeSeries(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	period := c.Query("period", "weekly")

	timeSeries, err := h.getStockTimeSeriesUseCase.Execute(symbol, period)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(timeSeries)
}

func (h *StockHandler) GetCompanyOverview(c *fiber.Ctx) error {
	symbol := c.Params("symbol")

	overview, err := h.getCompanyOverviewUseCase.Execute(symbol)
	if err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "rate limit") {
			status = fiber.StatusTooManyRequests
		} else if strings.Contains(err.Error(), "overview request failed") {
			status = fiber.StatusBadGateway
		}

		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(overview)
}

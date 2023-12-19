package web

import (
	"demo-golang/internal/services"
	"demo-golang/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

type CompanyRouter struct {
	companyService services.CompanyService
}

func NewCompanyRouter(companyService services.CompanyService) *CompanyRouter {
	return &CompanyRouter{companyService: companyService}
}

func (r *CompanyRouter) RegisterRoutes(e *echo.Echo) {
	e.POST("/companies", r.createCompany)
	e.GET("/companies/:id", r.getCompany)
	e.PUT("/companies/:id", r.updateCompany)
	e.DELETE("/companies/:id", r.deleteCompany)
	e.GET("/companies", r.listCompanies)
}

func (r *CompanyRouter) createCompany(c echo.Context) error {
	context := c.Request().Context()
	var company model.Company
	if err := c.Bind(&company); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := r.companyService.CreateCompany(context, company)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Company created successfully")
}

func (r *CompanyRouter) getCompany(c echo.Context) error {
	context := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	company, err := r.companyService.GetCompany(context, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, company)
}

func (r *CompanyRouter) updateCompany(c echo.Context) error {
	context := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var company model.Company
	if err := c.Bind(&company); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	company.Id = id
	err = r.companyService.UpdateCompany(context, company)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Company updated successfully")
}

func (r *CompanyRouter) deleteCompany(c echo.Context) error {
	context := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = r.companyService.DeleteCompany(context, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Company deleted successfully")
}

func (r *CompanyRouter) listCompanies(c echo.Context) error {
	context := c.Request().Context()
	ids := c.QueryParam("ids")
	idList := strings.Split(ids, ",")

	companies, err := r.companyService.ListCompanies(context, idList)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, companies)
}

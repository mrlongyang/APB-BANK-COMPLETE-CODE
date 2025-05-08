
package handlers

import (
    "database/sql"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/mrlongyang/apbbank/models"
)

type ProductHandler struct {
    DB *sql.DB
}

func (h *ProductHandler) Create(c echo.Context) error {
    p := new(models.Product)
    if err := c.Bind(p); err != nil {
        return err
    }
    _, err := h.DB.Exec("INSERT INTO product(product_id, product_name) VALUES (?, ?)", p.ProductID, p.ProductName)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, p)
}

func (h *ProductHandler) Read(c echo.Context) error {
    rows, err := h.DB.Query("SELECT product_id, product_name FROM product")
    if err != nil {
        return err
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var p models.Product
        rows.Scan(&p.ProductID, &p.ProductName)
        products = append(products, p)
    }
    return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Update(c echo.Context) error {
    p := new(models.Product)
    if err := c.Bind(p); err != nil {
        return err
    }
    _, err := h.DB.Exec("UPDATE product SET product_name=? WHERE product_id=?", p.ProductName, p.ProductID)
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Delete(c echo.Context) error {
    id := c.Param("id")
    _, err := h.DB.Exec("DELETE FROM product WHERE product_id=?", id)
    if err != nil {
        return err
    }
    return c.NoContent(http.StatusNoContent)
}

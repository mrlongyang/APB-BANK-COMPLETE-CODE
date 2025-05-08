
package handlers

import (
    "database/sql"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/mrlongyang/apbbank/models"
    "github.com/mrlongyang/apbbank/utils"
)

type UserHandler struct {
    DB *sql.DB
}

func (h *UserHandler) Register(c echo.Context) error {
    u := new(models.User)
    if err := c.Bind(u); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    _, err := h.DB.Exec("INSERT INTO user(user_id, user_name, user_password) VALUES (?, ?, ?)", u.UserID, u.UserName, u.UserPassword)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) Login(c echo.Context) error {
    u := new(models.User)
    if err := c.Bind(u); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    row := h.DB.QueryRow("SELECT user_password FROM user WHERE user_id = ?", u.UserID)
    var dbPassword string
    if err := row.Scan(&dbPassword); err != nil || dbPassword != u.UserPassword {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
    }

    token, _ := utils.GenerateToken(u.UserID)
    return c.JSON(http.StatusOK, map[string]string{"token": token})
}

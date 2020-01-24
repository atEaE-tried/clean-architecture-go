package controllers

import (
	"strconv"

	"github.com/atEaE-tried/clean-architecture-go/interfaces/database"
	"github.com/atEaE-tried/clean-architecture-go/usecase"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserControlller(sqlHandler database.SQLHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SQLHandler: sqlHandler
            }
        },
    }
}

func (ctrl *UserController) Create(c Context) {
    u := domain.User{}
    c.Bind(&u)
    err := ctrl.Interactor.Add(u)
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(201)
}

func (ctrl *UserController) Index (c Context){
    users, err := ctrl.Interactor.Users()
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(200, users)
}

func (ctrl *UserController) Show(c Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := ctrl.Interactor.UserById(id)
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(200, user)
}

package httpGin

import (
	"CLIappHabits/internal/entities"
	"CLIappHabits/internal/usecases"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type HabitHandler struct {
	creator usecases.HabitCreator
	getter  usecases.HabitGetter
	lister  usecases.HabitLister
	marker  usecases.HabitMarker
	deleter usecases.HabitDeleter
}

func NewHabitHandler(
	creator usecases.HabitCreator,
	getter usecases.HabitGetter,
	lister usecases.HabitLister,
	marker usecases.HabitMarker,
	deleter usecases.HabitDeleter) *HabitHandler {
	return &HabitHandler{
		creator: creator,
		getter:  getter,
		lister:  lister,
		marker:  marker,
		deleter: deleter,
	}
}

func (h *HabitHandler) InitRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/habit/:id", h.GetHabit)
		api.POST("/habit", h.CreateHabit)
		api.GET("/habits", h.ListHabits)
		api.PATCH("/habit/:id", h.Completed)
		api.DELETE("habit/:id", h.DeleteHabit)
	}

	r.GET("/")
}

type CreateHabitRequest struct {
	Name string `json:"name"`
}

type HabitResponse struct {
	HabitID        int64     `json:"habit_id"`
	Name           string    `json:"name"`
	Repetitions    int64     `json:"repetitions"`
	LastRepetition time.Time `json:"last_repetition,omitempty"`
}

func toHabitResponse(dto usecases.GetHabitOutputDTO) HabitResponse {
	return HabitResponse{
		HabitID:        dto.HabitID,
		Name:           dto.Name,
		Repetitions:    dto.Repetitions,
		LastRepetition: dto.LastRepetition,
	}
}

func (h *HabitHandler) CreateHabit(c *gin.Context) {

	var req CreateHabitRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": entities.ErrEmptyName})
		return
	}

	habitID, err := h.creator.CreateHabit(usecases.CreateHabitInputDTO{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(habitID.HabitID)})
	if err != nil {
		if errors.Is(err, entities.ErrHabitNotExists) {
			c.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	response := toHabitResponse(output)

	c.JSON(http.StatusOK, response)
}

func (h *HabitHandler) GetHabit(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	output, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(idInt)})
	if err != nil {
		if errors.Is(err, entities.ErrHabitNotExists) {
			c.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	response := toHabitResponse(output)

	c.JSON(http.StatusOK, response)

}

type ListHabitsResponse struct {
	Habits []HabitResponse
}

func toHabitsResponse(dto usecases.ListHabitsOutputDTO) *ListHabitsResponse {

	var habits ListHabitsResponse

	for _, v := range dto.Habits {
		habits.Habits = append(habits.Habits, toHabitResponse(v))
	}

	return &habits
}

func (h *HabitHandler) ListHabits(c *gin.Context) {
	list, err := h.lister.ListHabits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(list.Habits) == 0 {
		c.String(http.StatusOK, "У вас ещё нету ни одной привычки!")
		return
	}

	habits := toHabitsResponse(list)

	c.JSON(http.StatusOK, habits)
}

func (h *HabitHandler) Completed(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	err = h.marker.MarkHabit(usecases.MarkHabitInputDTO{HabitID: int64(idInt)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	output, err := h.getter.GetHabit(usecases.GetHabitInputDTO{HabitID: int64(idInt)})
	if err != nil {
		if errors.Is(err, entities.ErrHabitNotExists) {
			c.JSON(http.StatusNotFound, gin.H{"error:": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	response := toHabitResponse(output)

	c.JSON(http.StatusOK, response)
}

func (h *HabitHandler) DeleteHabit(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	err = h.deleter.DeleteHabit(usecases.DeleteHabitInputDTO{HabitID: int64(idInt)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

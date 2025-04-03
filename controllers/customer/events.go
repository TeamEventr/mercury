package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/IAmRiteshKoushik/mercury/cmd"
	"github.com/IAmRiteshKoushik/mercury/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEventCatalog(c *gin.Context) {
	// Handle Pagination query
	pageParam, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil || pageParam < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query params. Pagination should be integers",
		})
		return
	}
	offset := int32(pageParam)
	if int(offset) != pageParam {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query params. Pagination too high",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := cmd.DBPool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops! Something happened. Please try again later",
		})
		return
	}
	defer conn.Release()

	q := db.New()
	eventCatalog, err := q.FetchEventsPaginatedQuery(ctx, conn, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops! Something happened. Please try again later",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Events retrived successfully",
		"data":    eventCatalog,
	})
	return

}

func GetEventByEventId(c *gin.Context) {
	eventIdParam := c.Query("eventId")
	if eventIdParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is missing.",
		})
		return
	}
	eventId, err := uuid.Parse(eventIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is invalid.",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := cmd.DBPool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops! Something happened. Please try again later",
		})
		return
	}
	defer conn.Release()

	q := db.New()
	result, err := q.FetchEventByIdQuery(ctx, conn, eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops! Something happened. Please try again later",
		})
		return
	}
	if result.EventName == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Requested event was not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Event retrived successfully",
		"event_name":       result.EventName,
		"event_type":       result.EventType,
		"description":      result.Description,
		"tags":             result.Tags,
		"event_poster_url": result.EventPosterUrl,
		"venue":            result.Venue,
		"start_time":       result.StartTime,
		"end_time":         result.EndTime,
		"age_limit":        result.AgeLimit,
		"artists":          result.Artists,
		"price_tiers":      result.PriceTiers,
	})
	return
}

func BookEventTicketsCsrf(c *gin.Context) {
	eventIdParam := c.Query("eventId")
	if eventIdParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is missing.",
		})
		return
	}
	_, err := uuid.Parse(eventIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is invalid.",
		})
		return
	}

	// Check whether the Event exists or not from caching servers
	// If cache is empty only then check in database and setup the
	// CSRF token in the cookies

	c.JSON(http.StatusOK, gin.H{
		"message": "Access token retrived successfully",
		"token":   "CSRF-TOKEN-HERE",
	})
	return
}

func BookEventTickets(c *gin.Context) {
	eventIdParam := c.Query("eventId")
	if eventIdParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is missing.",
		})
		return
	}
	_, err := uuid.Parse(eventIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param. EventID is invalid.",
		})
		return
	}

	// Check whether the Event exists or not from caching servers
	// If cache is empty only then check in database and setup the
	// CSRF token in the cookies

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := cmd.DBPool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer conn.Release()

	c.JSON(http.StatusOK, gin.H{})
	return
}

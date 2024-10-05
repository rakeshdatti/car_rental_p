package controllers

import (
    "net/http"
    "time"
"log"
    "github.com/gin-gonic/gin"
    "runners-mysql/models"
    "runners-mysql/services"
)

type BookingsController struct {
    bookingService *services.BookingService
}

func NewBookingsController(bookingService *services.BookingService) *BookingsController {
    return &BookingsController{bookingService: bookingService}
}


func (ctrl *BookingsController) CreateBooking(c *gin.Context) {
    var booking models.Booking
    if err := c.ShouldBindJSON(&booking); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    log.Println(c.Get("userID"))
    userID, exists := c.Get("userID")
    role,_:= c.Get("role")
    if exists==false {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
        return
    }

    if role=="admin"{
        c.JSON(http.StatusForbidden,gin.H{"error":"Admins are not allowed to make bookings"})
        return
    }

    
    booking.UserID = userID.(int) 

   
    now := time.Now()                     
    booking.Start = now                   
    booking.End = now.Add(24 * time.Hour) 

    createBooking,err := ctrl.bookingService.CreateBooking(&booking);
    if  err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated,gin.H{"message":"Booking Sucessfully","booking": createBooking})
}
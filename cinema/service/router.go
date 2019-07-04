package service

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"lucas-lima/cinema/configs"
	"lucas-lima/cinema/control"
	"net/http"
)

func createHandler()(handler *mux.Router){

	handler = mux.NewRouter()

	//handleFunc method for Login
	handler.HandleFunc(configs.LOGIN_PATH,control.LoginUser).Methods(http.MethodPost).Headers("content-type","application/json")

	// handleFunc method for User
	handler.HandleFunc(configs.USER_PATH, control.CreateUser).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.USER_PATH, control.UpdateUser).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.USER_PATH+"specific/", control.UpdateUserSpecific).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.USER_PATH, control.GetUserId).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.USER_PATH, control.GetAllUser).Methods(http.MethodGet)
	handler.HandleFunc(configs.USER_PATH, control.DeleteUser).Methods(http.MethodDelete).Queries("id","{value}")

	// handleFunc methods for Movie
	handler.HandleFunc(configs.MOVIE_PATH, control.CreateMovie).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.MOVIE_PATH, control.UpdateMovie).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.MOVIE_PATH, control.GetMovieID).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.MOVIE_PATH, control.GetAllMovie).Methods(http.MethodGet)
	handler.HandleFunc(configs.MOVIE_PATH, control.DeleteMovie).Methods(http.MethodDelete).Queries("id","{value}")

	//handleFunc methods for Room
	handler.HandleFunc(configs.ROOM_PATH, control.CreateRoom).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.ROOM_PATH, control.UpdateRoom).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.ROOM_PATH, control.GetRoomID).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.ROOM_PATH, control.GetAllRoom).Methods(http.MethodGet)
	handler.HandleFunc(configs.ROOM_PATH, control.DeleteRoom).Methods(http.MethodDelete).Queries("id","{value}")

	//handleFunc methods for Session
	handler.HandleFunc(configs.SESSION_PATH, control.CreateSession).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.SESSION_PATH, control.UpdateSession).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.SESSION_PATH, control.GetSessionID).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.SESSION_PATH, control.GetAllSession).Methods(http.MethodGet)
	handler.HandleFunc(configs.SESSION_PATH, control.DeleteSession).Methods(http.MethodDelete).Queries("id","{value}")

	//handleFunc methods for Cine
	handler.HandleFunc(configs.CINE_PATH, control.CreateCine).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.CINE_PATH, control.UpdateCine).Methods(http.MethodPut).Headers("content-type","application/json").Queries("id","{value}")
	handler.HandleFunc(configs.CINE_PATH, control.GetCineID).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.CINE_PATH, control.GetAllCine).Methods(http.MethodGet)
	handler.HandleFunc(configs.CINE_PATH, control.DeleteCine).Methods(http.MethodDelete).Queries("id","{value}")

	//handleFunc methods for Reservation
	handler.HandleFunc(configs.RESERVATION_PATH, control.CreateReservation).Methods(http.MethodPost).Headers("content-type","application/json")
	handler.HandleFunc(configs.RESERVATION_PATH, control.GetReservationID).Methods(http.MethodGet).Queries("id","{value}")
	handler.HandleFunc(configs.RESERVATION_PATH, control.GetAllReservation).Methods(http.MethodGet)
	handler.HandleFunc(configs.RESERVATION_PATH, control.DeleteReservation).Methods(http.MethodDelete).Queries("id","{value}")

	//return handle
	return
}

func createCORS() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})
}

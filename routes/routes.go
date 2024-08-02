package Routes

import (
	"github.com/gorilla/mux"
	user_routes "github.com/rajpurkait9/hotel_backend/routes/users"
)

func Routes(r *mux.Router) {
	user_routes.User_routes(r)
}

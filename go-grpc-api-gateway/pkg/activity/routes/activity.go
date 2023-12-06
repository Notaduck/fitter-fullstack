package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/notaduck/go-grpc-api-gateway/pkg/activity/pb"
)

type GetActivityRequestBody struct {
	ActivityID int64 `json:"activityId"`
}

func Register(w http.ResponseWriter, r *http.Request, c pb.ActivityServiceClient) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := GetActivityRequestBody{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	res, err := c.GetActivity(context.Background(), &pb.GetActivityRequest{
		UserId:     body.ActivityID,
		ActivityId: 1,
	})

	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.Status))

	responseJSON, _ := json.Marshal(res)
	w.Write(responseJSON)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var postMessages []PostMessage

// ParentPost defines the structure for an API parentPost
type ParentPost struct {
	PostID              int          `json:"postId"`              // Group postID
	TotalNumberComments int          `json:"totalNumberComments"` // Total number of comments
	Post                []*ChildPost `json:"post"`                // Each post
}

// ChildPost defines the structure for an API childPost
type ChildPost struct {
	ID    int    `json:"id"` // Unique identifier for the post-message
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}

// PostMessage defines the structure for an API post
type PostMessage struct {
	PostID int    `json:"postId"` // Group postID
	ID     int    `json:"id"`     // Unique identifier for the post-message
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Get all post
func getPostMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(postMessages)
}

// Get a single PostMessage
func getPostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["post_id"])
	if err != nil {
		panic(err)
	}

	for _, item := range postMessages {
		if item.ID == postID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&PostMessage{})

}

// Get group of PostMessage
func getPostMessageSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["group_id"])
	if err != nil {
		panic(err)
	}

	pp := new(ParentPost)
	count := 0

	for _, item := range postMessages {
		if item.PostID == groupID {
			cp := &ChildPost{
				ID:    item.ID,
				Email: item.Email,
				Name:  item.Name,
				Body:  item.Body,
			}
			count = count + 1

			pp.Post = append(pp.Post, cp)
		}
	}

	pp.PostID = groupID
	pp.TotalNumberComments = count

	json.NewEncoder(w).Encode(pp)
}

// Search API
func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	query := r.URL.Query()

	postID := 0
	valPostID := query.Get("post_id")
	if valPostID != "" {
		postID, _ = strconv.Atoi(valPostID)
	}

	valEmail := query.Get("email")
	valName := query.Get("name")
	valBody := query.Get("body")

	for _, item := range postMessages {

		if item.ID == postID && strings.Contains(item.Email, valEmail) == true && strings.Contains(item.Name, valName) == true && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}

		if item.ID == postID && strings.Contains(item.Email, valEmail) == true && strings.Contains(item.Name, valName) == true && valBody == "" {
			json.NewEncoder(w).Encode(item)
			return
		}
		if item.ID == postID && strings.Contains(item.Email, valEmail) == true && valName == "" && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && strings.Contains(item.Email, valEmail) == true && strings.Contains(item.Name, valName) == true && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
		if item.ID == postID && valEmail == "" && strings.Contains(item.Name, valName) == true && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}

		if item.ID == postID && strings.Contains(item.Email, valEmail) == true && valName == "" && valBody == "" {
			json.NewEncoder(w).Encode(item)
			return
		}
		if item.ID == postID && strings.Contains(item.Name, valName) == true && valEmail == "" && valBody == "" {
			json.NewEncoder(w).Encode(item)
			return
		}
		if item.ID == postID && strings.Contains(item.Body, valBody) == true && valEmail == "" && valName == "" {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && strings.Contains(item.Email, valEmail) == true && strings.Contains(item.Name, valName) == true && valBody == "" {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && strings.Contains(item.Email, valEmail) == true && valName == "" && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && valEmail == "" && strings.Contains(item.Name, valName) == true && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}

		if valEmail == "" && valName == "" && valBody == "" && item.ID == postID {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && valName == "" && valBody == "" && strings.Contains(item.Email, valEmail) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && valEmail == "" && valBody == "" && strings.Contains(item.Name, valName) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
		if valPostID == "" && valEmail == "" && valName == "" && strings.Contains(item.Body, valBody) == true {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// As default
	json.NewEncoder(w).Encode(&PostMessage{})

}

func main() {

	fmt.Println("Started on port :8000")

	// Init router
	r := mux.NewRouter()

	// Init data source
	// Refer https://jsonplaceholder.typicode.com/comments
	postMessages = append(postMessages, PostMessage{PostID: 1, ID: 1, Name: "id labore ex et quam laborum", Email: "Eliseo@gardner.biz", Body: "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"})
	postMessages = append(postMessages, PostMessage{PostID: 1, ID: 2, Name: "quo vero reiciendis velit similique earum", Email: "Jayne_Kuhic@sydney.com", Body: "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et"})
	postMessages = append(postMessages, PostMessage{PostID: 1, ID: 3, Name: "odio adipisci rerum aut animi", Email: "Nikita@garfield.biz", Body: "quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"})
	postMessages = append(postMessages, PostMessage{PostID: 1, ID: 4, Name: "alias odio sit", Email: "Lew@alysha.tv", Body: "non et atque\noccaecati deserunt quas accusantium unde odit nobis qui voluptatem\nquia voluptas consequuntur itaque dolor\net qui rerum deleniti ut occaecati"})
	postMessages = append(postMessages, PostMessage{PostID: 1, ID: 5, Name: "vero eaque aliquid doloribus et culpa", Email: "Hayden@althea.biz", Body: "harum non quasi et ratione\ntempore iure ex voluptates in ratione\nharum architecto fugit inventore cupiditate\nvoluptates magni quo et"})
	postMessages = append(postMessages, PostMessage{PostID: 2, ID: 6, Name: "et fugit eligendi deleniti quidem qui sint nihil autem", Email: "Presley.Mueller@myrl.com", Body: "doloribus at sed quis culpa deserunt consectetur qui praesentium\naccusamus fugiat dicta\nvoluptatem rerum ut voluptate autem\nvoluptatem repellendus aspernatur dolorem in"})
	postMessages = append(postMessages, PostMessage{PostID: 2, ID: 10, Name: "eaque et deleniti atque tenetur ut quo ut", Email: "Carmen_Keeling@caroline.name", Body: "voluptate iusto quis nobis reprehenderit ipsum amet nulla\nquia quas dolores velit et non\naut quia necessitatibus\nnostrum quaerat nulla et accusamus nisi facilis"})
	postMessages = append(postMessages, PostMessage{PostID: 3, ID: 14, Name: "et officiis id praesentium hic aut ipsa dolorem repudiandae", Email: "Nathan@solon.io", Body: "vel quae voluptas qui exercitationem\nvoluptatibus unde sed\nminima et qui ipsam aspernatur\nexpedita magnam laudantium et et quaerat ut qui dolorum"})

	// Route handlers
	r.HandleFunc("/api/posts", getPostMessages).Methods("GET")
	r.HandleFunc("/api/posts/{post_id:[0-9]+}", getPostMessage).Methods("GET")
	r.HandleFunc("/api/postgroup/{group_id:[0-9]+}", getPostMessageSet).Methods("GET")
	r.HandleFunc("/api/comments", getComments).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}

package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anfelo/comments-api/internal/comment"
	"github.com/gorilla/mux"
)

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Unable to parse UINT from ID", Error: err.Error()})
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Error Retrieving Comment By", Error: err.Error()})
		return
	}

	RespondJson(w, http.StatusOK, comment)
}

// GetAllComments - retrieve all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to retrieve all comments", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, comments)
}

// PostComment - creates a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to decode JSON Body", Error: err.Error()})
		return
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to post new comment", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusCreated, comment)
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to decode JSON Body", Error: err.Error()})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to parse uint from ID", Error: err.Error()})
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to update comment", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, comment)
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to parse uint from ID", Error: err.Error()})
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to delete comment by comment ID", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, Response{Message: "Successfully deleted comment"})
}

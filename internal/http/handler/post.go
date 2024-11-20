package handler

import (
	"boobook/internal/repository/model"
	"boobook/internal/service"
	"boobook/internal/slogger"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
	"time"
)

type postHandler struct {
	logger      *slog.Logger
	postService service.PostService
	clientRedis *redis.Client
}

func NewPostHandler(logger *slog.Logger, postService service.PostService, clientRedis *redis.Client) PostHandler {
	return &postHandler{
		logger:      logger,
		postService: postService,
		clientRedis: clientRedis,
	}
}

func (h *postHandler) GetList(ctx *gin.Context) {
	posts, err := h.getPostsFromCache(ctx)
	if err != nil {
		h.logger.WarnContext(ctx, "failed to get posts from cache, fetching from service", slogger.Err(err))
	}

	if posts == nil {
		posts, err = h.postService.GetList()
		if err != nil {
			h.respondWithError(ctx, http.StatusInternalServerError, "failed to get post list", err)
			return
		}

		err = h.cachePosts(ctx, posts)
		if err != nil {
			h.logger.WarnContext(ctx, "failed to cache posts", slogger.Err(err))
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (h *postHandler) getPostsFromCache(ctx *gin.Context) ([]*model.Post, error) {
	postsJson, err := h.clientRedis.Get(ctx, "post:feed").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // Ключ отсутствует в Redis
		}
		return nil, err
	}

	var posts []*model.Post
	if err := json.Unmarshal([]byte(postsJson), &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (h *postHandler) cachePosts(ctx *gin.Context, posts []*model.Post) error {
	postsJson, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	return h.clientRedis.Set(ctx, "post:feed", postsJson, 2*time.Hour).Err()
}

func (h *postHandler) respondWithError(ctx *gin.Context, statusCode int, message string, err error) {
	h.logger.ErrorContext(ctx, message, slogger.Err(err))
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}

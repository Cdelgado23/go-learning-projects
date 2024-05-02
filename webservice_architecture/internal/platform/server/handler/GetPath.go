package handler

import (
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/path"
	"github.com/gin-gonic/gin"
)

type NodeRequest struct {
	X float64
	Y float64
	Z float64
}

func NewNodeRequest(x, y, z float64) NodeRequest {
	return NodeRequest{x, y, z}
}

type PathRequest struct {
	StartNode struct {
		SX float64 `form:"sx"`
		SY float64 `form:"sy"`
		SZ float64 `form:"sz"`
	}
	EndNode struct {
		EX float64 `form:"ex"`
		EY float64 `form:"ey"`
		EZ float64 `form:"ez"`
	}
}

func GetPathsHandler(service path.PathFinderService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request PathRequest

		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		start := NewNodeRequest(request.StartNode.SX, request.StartNode.SY, request.StartNode.SZ)
		end := NewNodeRequest(request.EndNode.EX, request.EndNode.EY, request.EndNode.EZ)

		foundPath := service.PathFromTo(NodeRequestToLocation(start), NodeRequestToLocation(end))

		ctx.JSON(200, gin.H{"foundPath": foundPath})
	}
}

func NodeRequestToLocation(node NodeRequest) path.Location {
	return path.NewLocation(node.X, node.Y, node.Z)
}

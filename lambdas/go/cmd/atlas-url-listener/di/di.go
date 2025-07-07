package di

import (
	"github.com/alphatechnolog/atlas-infra/internal/handler"
	usecase "github.com/alphatechnolog/atlas-infra/usecase/atlas-url-listener"
)

func InitialiseAtlasURLListenerDependencies() *handler.URLListenerHandler {
	useCase := usecase.NewURLListenerUseCase()
	return handler.NewURLListenerHandler(useCase)
}

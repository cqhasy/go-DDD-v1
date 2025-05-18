package match

type RouteMatch interface {
	GetUrl() string
	GetTarget() string
}

package config

const (
	CAT_FACTS = "https://cat-fact.herokuapp.com/facts"
	CUTE_CATS = "https://cataas.com/api/cats?tags=cute"
)

type config struct {
	EndPoints  []string
	Concurency int
}

var clientConfig config

func init() {
	clientConfig = config{
		EndPoints:  []string{CUTE_CATS, CUTE_CATS, CAT_FACTS},
		Concurency: 3,
	}
}

func GetClientConfig() config {
	return clientConfig
}

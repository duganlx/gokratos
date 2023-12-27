package main

type Cfg struct {
	Server struct {
		Http struct {
			Addr    string `json:"addr"`
			Timeout string `json:"Timeout"`
		} `json:"http"`

		Grpc struct {
			Addr    string `json:"addr"`
			Timeout string `json:"Timeout"`
		} `json:"grpc"`
	} `json:"Server"`
}

func main() {

	// c := config.New()

}

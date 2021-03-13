package main

func main() {
	conf := NewConf()
	StartJobs(conf)
	StartUpDataToEc(conf)
	StartHttpServer(conf)
}

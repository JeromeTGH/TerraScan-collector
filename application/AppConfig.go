package application

type Config struct {
	Lcd struct {
		url string
	}
	Bdd struct {
		host_name string
		bdd_name  string
		user_name string
		password  string
		port      int16
	}
	Email struct {
		host_name string
		smtp_port int16
		from      string
		pwd       string
		to        string
	}
}
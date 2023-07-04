package config

type Schedule struct {
	Time string `hcl:"time"`
}

type EC2 struct {
	Name string `hcl:"name"`
}

type EC2WithSchedule struct {
	//EC2
	Schedule `hcl:"schedule"`
}

type AWS struct {
	Label string          `hcl:"name,label"`
	EC2   EC2WithSchedule `hcl:"ec2,block"`
}

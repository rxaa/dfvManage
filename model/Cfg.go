package model

type Cfg struct {
	List []*Proc
}

type Proc struct {
	Name    string
	Cmd     string
	Outmenu string
	Pid     int
}

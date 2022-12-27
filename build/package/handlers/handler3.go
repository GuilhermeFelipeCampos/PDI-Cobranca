package handlers

import "PDI-COBRANCA/build/package/model"

type RepositoryInterface interface {
	InsertUsers(u model.Users) (msg, err error)
	GetUsers() (us []model.Users, err error)
}

type pudim struct {
}

func (p *pudim) InsertUsers(u model.Users) (msg, err error) {

	return nil, nil
}
func (p *pudim) GetUsers() (us []model.Users, err error) {

	return []model.Users{}, nil
}

func main() {
	//	p := pudim{}
	//hd := NewHandler(&p)

}

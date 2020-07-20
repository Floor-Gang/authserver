package internal

func (a *AuthServer) Register(args *Feature, res *RegisterResponse) error {
	name := args.Name
	a.Features[name] = args

	*res = RegisterResponse{
		Token:   a.client.Token,
		Serving: a.config.Guild,
	}

	return nil
}

package OgameUtil

import "bitbucket.org/jc01rho/ogame"

func ResourcePricesSum (res ogame.Resources)  int64 {

	return res.Metal + res.Crystal + res.Deuterium

}
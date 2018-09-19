module Portfolio.Msg exposing (..)

import Http
import Model.Portfolio exposing (Portfolio)
import Material


type Msg
    = Mdl (Material.Msg Msg)
    | OnLoadPortfolioResponse (Result Http.Error Portfolio)

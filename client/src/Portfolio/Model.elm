module Portfolio.Model exposing (..)

import Material
import Model.Portfolio exposing (Portfolio)


type alias Model =
    { mdl : Material.Model
    , portfolio : Maybe Portfolio
    , error : Maybe String
    }


init : Model
init =
    { mdl = Material.model
    , portfolio = Nothing
    , error = Nothing
    }

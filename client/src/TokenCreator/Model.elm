module TokenCreator.Model exposing (Model, init)

import Material


type alias Model =
    { mdl : Material.Model
    , error : Maybe String
    , name : String
    , totalSupply : String
    }


init : Model
init =
    { mdl = Material.model
    , error = Nothing
    , name = ""
    , totalSupply = ""
    }

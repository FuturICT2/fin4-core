module Tokens.Model exposing (Model, init)

import Material
import Model.Tokens exposing (Tokens)


type alias Model =
    { mdl : Material.Model
    , tokens : Maybe Tokens
    , error : Maybe String
    }


init : Model
init =
    { mdl = Material.model
    , tokens = Nothing
    , error = Nothing
    }

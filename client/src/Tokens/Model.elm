module Tokens.Model exposing (Model, init)

import Main.User exposing (User, Users)
import Material
import Model.Tokens exposing (Tokens)


type alias Model =
    { mdl : Material.Model
    , tokens : Maybe Tokens
    , error : Maybe String
    , selectedTab : Int
    }


init : Model
init =
    { mdl = Material.model
    , tokens = Nothing
    , error = Nothing
    , selectedTab = 0
    }

module Actions.Model exposing (Model, init)

import Dict exposing (Dict)
import Material
import Model.Actions exposing (Actions)


type alias Model =
    { mdl : Material.Model
    , actions : Maybe Actions
    , error : Maybe String
    , proposals : Dict Int String
    }


init : Model
init =
    { mdl = Material.model
    , actions = Nothing
    , error = Nothing
    , proposals = Dict.empty
    }
